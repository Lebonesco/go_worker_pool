package main 

import (
	"log"
	//"./limiter"
	"tutorials/concurrent-limiter/pool"
	"tutorials/concurrent-limiter/work"
)

const WORKER_COUNT = 5
const JOB_COUNT = 100

func main() {
	log.Println("starting application...")
	collector := pool.StartDispatcher(WORKER_COUNT) // start up worker pool

	for i, job := range work.CreateJobs(JOB_COUNT) {
		collector.Work <-pool.Work{Job: job, ID: i}
	}
}	