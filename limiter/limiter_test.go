package limiter

import (
	"sync"
	"testing"
	"time"
)

func TestExample(t *testing.T) {
	N := 1000
	limit := NewConcurrencyLimiter(10)
	doneCount := 0
	doneMap := make(map[int]bool)
	mutex := &sync.Mutex{}
	for i := 0; i < N; i++ {
		doneMap[i] = false
	}
	for i := 0; i < N; i++ {
		iter := i
		limit.Execute(func() {
			time.Sleep(time.Millisecond * 5)
			mutex.Lock()
			doneMap[iter] = true
			doneCount++
			mutex.Unlock()
		})
	}
	limit.Wait()

	t.Log("doneCount", doneCount, N)
	if N != doneCount {
		t.Error("Done count not equals number of jobs")
	}
	for id := 0; id < N; id++ {
		status := doneMap[id]
		if !status {
			t.Error("Job with id", id, " not finished", status)
		}
	}
}

func TestExecute(t *testing.T) {
	LIMIT := 15
	N := 100

	c := NewConcurrencyLimiter(LIMIT)
	m := make(map[int]bool)
	lock := &sync.Mutex{}

	max := int64(0)
	for i := 0; i < N; i++ {
		x := i
		c.Execute(func() {
			lock.Lock()
			m[x] = true
			currentMax := c.GetNumInProgress()
			if currentMax > max {
				max = currentMax
			}
			lock.Unlock()
		})
	}

	c.Wait()
	t.Log("results:", len(m))
	t.Log("max:", max)
	if len(m) != N {
		t.Error("Invalid number of results", len(m), N)
	}

	if max > int64(LIMIT) || max == 0 {
		t.Error("Invalid number of max", max, LIMIT)
	}
}

func TestExecuteWithRoutineId(t *testing.T) {
	LIMIT := 15
	N := 100

	c := NewConcurrencyLimiter(LIMIT)
	m := make(map[int]int)
	lock := &sync.Mutex{}

	for i := 0; i < N; i++ {
		c.ExecuteWithRoutineId(func(routineId int) {
			lock.Lock()
			m[routineId] += 1
			if routineId >= LIMIT {
				t.Errorf("expected routineId %d to be less than %d", routineId, LIMIT)
			}
			lock.Unlock()
		})
	}

	c.Wait()

	sum := 0
	for _, count := range m {
		sum += count
	}

	if sum != N {
		t.Errorf("invalid number of results: %d expected %d", sum, N)
	}
}

func TestLimit(t *testing.T) {
	c := NewConcurrencyLimiter(0)
	if c.limit != DefaultLimit {
		t.Errorf("invalid default limit: %d expected %d", c.limit, DefaultLimit)
	}

	LIMIT := DefaultLimit + DefaultLimit/2
	c = NewConcurrencyLimiter(LIMIT)
	if c.limit != LIMIT {
		t.Errorf("invalid limit: %d expected %d", c.limit, LIMIT)
	}
}
