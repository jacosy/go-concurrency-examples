package main

import (
	"fmt"
	"sync"
	"time"
)

// Job represents a unit of work
type JobFn func() error

// WorkerPool manages a pool of workers
type WorkerPool struct {
	workerCount int
	jobQueue    chan JobFn
	quit        chan bool
	wg          sync.WaitGroup
}

// NewWorkerPool creates a new worker pool
func NewWorkerPool(workerCount, queueSize int) *WorkerPool {
	return &WorkerPool{
		workerCount: workerCount,
		jobQueue:    make(chan JobFn, queueSize),
		quit:        make(chan bool),
	}
}

// Start begins the worker pool
func (wp *WorkerPool) Start() {
	for i := range wp.workerCount {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

// worker is the main worker loop
func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()

	for {
		select {
		case job := <-wp.jobQueue:
			fmt.Printf("Worker %d processing job\n", id)
			if err := job(); err != nil {
				fmt.Printf("Worker %d error: %v\n", id, err)
			}
		case <-wp.quit:
			fmt.Printf("Worker %d stopping\n", id)
			return
		}
	}
}

// Submit adds a job to the queue
func (wp *WorkerPool) Submit(job JobFn) {
	wp.jobQueue <- job
}

// Stop gracefully shuts down the worker pool
func (wp *WorkerPool) Stop() {
	close(wp.quit)
	wp.wg.Wait()
	close(wp.jobQueue)
}

// Example usage
func basicWorkerPoolUsageExample() {
	// Create a pool with 3 workers and queue size of 10
	pool := NewWorkerPool(3, 10)
	pool.Start()

	// Submit some jobs
	for i := range 10 {
		jobID := i
		pool.Submit(func() error {
			fmt.Printf("Executing job %d\n", jobID)
			time.Sleep(time.Millisecond * 100) // Simulate work
			return nil
		})
	}

	// Let jobs process
	time.Sleep(time.Second * 2)

	// Shutdown
	pool.Stop()
	fmt.Println("All workers stopped")
}
