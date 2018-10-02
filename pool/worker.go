package pool

import (
	"log"
	"tutorials/concurrent-limiter/job"
)

// worker struct
type Worker struct {
	ID int
	WorkerChannel chan chan Work
	Channel chan Work
	End chan bool
}

// start worker
func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerChannel <-w.Channel
			select {
			case work := <-w.Channel:
				// do work
				job.DoWork(work.Job, w.ID)
			case <-w.End:
				return 
			}
		}
	}()
}

// end worker
func (w *Worker) Stop() {
	log.Printf("worker [%d] is stopping", w.ID)
	w.End <- true
}