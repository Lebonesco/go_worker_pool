package pool

import (
	"log"
)

var WorkerChannel = make(chan chan Work)

func StartDispatcher(workerCount int) chan<- Work {
	var i int
	input := make(chan Work)
	for i < workerCount {
		i++
		log.Println("starting worker: ", i)
		worker := Worker{
							ID: i,
							Channel: make(chan Work),
							WorkerChannel: WorkerChannel,
							End: make(chan bool)}
		worker.Start()
	}

	// start collector
	go func() {
		for work := range input {
			worker := <-WorkerChannel // wait for available channel
			worker <-work // dispatch work to worker
		}
	}()
	return input
}