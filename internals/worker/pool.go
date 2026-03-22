package worker

import "sync"

func StartWorkerPool(numWorkers int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			Worker(id, jobs, results)
		}(i)
	}
}
