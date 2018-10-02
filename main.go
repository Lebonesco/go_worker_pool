package main 

import (
	"log"
	//"./limiter"
	"tutorials/concurrent-limiter/pool"

	"math/rand"
)

const WORKER_COUNT = 5
const JOB_COUNT = 100

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}

func CreateJobs(amount int) []string {
	log.Println("creating jobs...")
	var jobs []string

	for i := 0; i < amount; i++ {
		jobs = append(jobs, RandStringRunes(8))
	}
	return jobs
}

func main() {
	log.Println("starting application...")
	collector := pool.StartDispatcher(WORKER_COUNT) // start up worker pool

	for i, job := range CreateJobs(JOB_COUNT) {
		collector.Work <-pool.Work{Job: job, ID: i}
	}
}	