package main

import (
	"fmt"
	"worker-pool/worker"
	"sync"
)

func main() {
	const numWorkers = 3
	const numJobs = 10

	 jobs := make(chan worker.Job,numJobs)
	 results := make(chan worker.Result,numJobs)

	 var wg sync.WaitGroup

	// Start workers
	for workerId:=1;workerId<=numWorkers;workerId++{
        wg.Add(1)
		go func(id int){
			defer wg.Done()
			worker.Worker(id,jobs,results)
		}(workerId)
	}
	//creatig and sending jobs
	for j:=1;j<=numJobs;j++{
	    jobs<-worker.Job{
			ID:j,
			Name:fmt.Sprintf("Task-%d",j),
		}
	}
	close(jobs)
	go func(){
		wg.Wait()
		close(results)
	}()

	for result:=range results{
		fmt.Println(result.Processed)
	}

	// Wait for all workers to finish processing


	fmt.Println("🎉 All tasks processed")
}
