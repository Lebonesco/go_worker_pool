package main 

import (
	"log"
	//"./limiter"
	"./pool"
)

var jobs = []string{"apple", "pear", "orange", "banana", "mango", "grapes", "kiwi"}

const WORKER_COUNT = 3

func main() {
	log.Println("starting application...")
	collector := pool.StartDispatcher(WORKER_COUNT) // start up worker pool

	for i, job := range jobs {
		collector <-pool.Work{Job: job, ID: i}
	}
	
}	