package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- [2]int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		r := [2]int{id, j * 2}
		results <- r
	}
}

func main() {
	fmt.Println("start")
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan [2]int, numJobs)

	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	// Assign jobs and close channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for r := 1; r <= numJobs; r++ {
		res := <-results
		fmt.Println("worker", res[0], "result", res[1])
	}

	fmt.Println("stop")
}
