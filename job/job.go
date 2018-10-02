package job

import (
	"log"
	"hash/fnv"
	"time"
)

// mimics any type of job that can be run concurrently
func DoWork(word string, id int) {
	h := fnv.New32a()
	h.Write([]byte(word))
	time.Sleep(time.Second)
	log.Printf("worker [%d] - created hash [%d] from word [%s]\n", id, h.Sum32(), word)
}