package pool

import (
	"log"
	"hash/fnv"	
	"time"
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
				h := fnv.New32a()
				h.Write([]byte(work.Job))
				time.Sleep(time.Second)
				log.Printf("worker [%d] - created hash [%d] from word [%s]\n", w.ID, h.Sum32(), work.Job)
			case <-w.End:
				return 
			}
		}
	}()
}

// end worker
func (w *Worker) Stop() {
	w.End <- true
}