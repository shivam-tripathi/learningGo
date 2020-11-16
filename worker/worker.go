package worker

import (
	"errors"
	"sync"
)

type WorkerJob struct {
	Job         func()
	stop        bool
	syncChannel chan bool
	waitGroup   *sync.WaitGroup
}

type Worker struct {
	Name  string
	queue chan WorkerJob
}

func (worker *Worker) Stop(f func()) {
	workerJob := WorkerJob{Job: f, stop: true, syncChannel: nil}
	worker.queue <- workerJob
}

func (worker *Worker) Close() {
	worker.Stop(func() {})
}

func (worker *Worker) Enqueue(job func(), waitGroup *sync.WaitGroup) {
	workerJob := WorkerJob{Job: job, stop: false, syncChannel: nil, waitGroup: waitGroup}
	worker.queue <- workerJob
}

func (worker *Worker) EnqueueSync(job func()) error {
	syncChannel := make(chan bool)
	workerJob := WorkerJob{Job: job, stop: false, syncChannel: syncChannel, waitGroup: nil}
	worker.queue <- workerJob
	status := <-syncChannel
	if !status {
		return errors.New("Sync failed")
	}
	return nil
}

func (worker *Worker) getQueueSize() int {
	return len(worker.queue)
}

func NewWorker(name string, queue chan WorkerJob) *Worker {
	if queue == nil {
		queue = make(chan WorkerJob, 1000)
	}

	workerLoop := func(w *Worker) {
		for {
			workerJob := <-w.queue
			workerJob.Job()
			if workerJob.stop {
				break
			}
			if workerJob.syncChannel != nil {
				workerJob.syncChannel <- true
			}
			if workerJob.waitGroup != nil {
				workerJob.waitGroup.Done()
			}
		}
	}

	w := Worker{Name: name, queue: queue}
	go workerLoop(&w)
	return &w
}

func NewWorkerDefault(name string) *Worker {
	return NewWorker(name, nil)
}
