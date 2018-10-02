package main 

import "testing"
import "tutorials/concurrent-limiter/pool"
import "tutorials/concurrent-limiter/job"

func Benchmark1(b *testing.B) {
	collector := pool.StartDispatcher(WORKER_COUNT) // start up worker pool

	for n := 0; n < b.N; n++ {
		for i, job := range CreateJobs(20) {
			collector.Work <-pool.Work{Job: job, ID: i}
		}
	}
}

func Benchmark2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, work := range CreateJobs(20) {
			job.DoWork(work, 1)
		}
	}
}