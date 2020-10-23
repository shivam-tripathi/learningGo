package limiter

import (
	"sync/atomic"
)

const (
	DefaultLimit = 100
)

type ConcurrencyLimiter struct {
	limit         int
	routines      chan int
	numInProgress int64
}

func NewConcurrencyLimiter(limit int) *ConcurrencyLimiter {
	if limit <= 0 {
		limit = DefaultLimit
	}

	c := ConcurrencyLimiter{limit: limit, routines: make(chan int, limit)}

	for i := 0; i < limit; i++ {
		c.routines <- i
	}

	return &c
}

func (c *ConcurrencyLimiter) Execute(job func()) int {
	routine := <-c.routines
	atomic.AddInt64(&c.numInProgress, 1)

	go func() {
		defer func() {
			c.routines <- routine
			atomic.AddInt64(&c.numInProgress, -1)
		}()
		job()
	}()
	return routine
}

func (c *ConcurrencyLimiter) ExecuteWithRoutineId(job func(routineId int)) int {
	routine := <-c.routines
	atomic.AddInt64(&c.numInProgress, 1)
	go func() {
		defer func() {
			c.routines <- routine
			atomic.AddInt64(&c.numInProgress, -1)
		}()
		job(routine)
	}()
	return routine
}

func (c *ConcurrencyLimiter) Wait() {
	for i := 0; i < c.limit; i++ {
		<-c.routines
	}
}

func (c *ConcurrencyLimiter) GetNumInProgress() int64 {
	return atomic.LoadInt64(&c.numInProgress)
}
