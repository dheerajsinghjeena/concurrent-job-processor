package main

import (
	"fmt"
	"sync"

	"github.com/dheerajsinghjeena/concurrent-job-processor/internals/worker"
)

func main() {
	jobs := make(chan worker.Job)
	results := make(chan worker.Result)

	var wg sync.WaitGroup

	// start workers
	worker.StartWorkerPool(3, jobs, results, &wg)

	// producer
	go func() {
		for i := 1; i <= 5; i++ {
			jobs <- worker.Job{ID: i}
		}
		close(jobs)
	}()

	// close results after workers finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// consume results
	for res := range results {
		fmt.Println(res)
	}
}
