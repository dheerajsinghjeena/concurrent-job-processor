package worker

import (
	"fmt"
	"time"
)

func Worker(id int, jobs <-chan Job, results chan<- Result) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job.ID)

		time.Sleep(time.Second)

		results <- Result{
			JobID:  job.ID,
			Output: "done",
		}
	}
}
