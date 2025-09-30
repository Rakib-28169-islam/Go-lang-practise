package worker

import (
	"fmt"
	"time"
)

// Job represents a task to process
type Job struct {
	ID   int
	Name string
}

// Result represents the result of a processed job
type Result struct {
	Job       Job
	Processed string
}

// Worker processes jobs from the jobs channel and sends results to results channel
func Worker(id int,jobs <-chan Job,results chan<- Result){

	for job:=range jobs{
		time.Sleep(1*time.Second) // Simulate work
		results<-Result{
			Job:job,
			Processed: fmt.Sprintf("Worker id: (%d) Job Processed: (%s)",id,job.Name),
		}
		fmt.Println()
   }
}
// Example Output:


// Worker id: (3) Job Processed: (Task-3)
// Worker id: (1) Job Processed: (Task-1)
// Worker id: (2) Job Processed: (Task-2)

// Worker id: (2) Job Processed: (Task-6)


// Worker id: (3) Job Processed: (Task-4)
// Worker id: (1) Job Processed: (Task-5)

// Worker id: (3) Job Processed: (Task-8)


// Worker id: (1) Job Processed: (Task-9)
// Worker id: (2) Job Processed: (Task-7)

// Worker id: (3) Job Processed: (Task-10)
// 🎉 All tasks processed