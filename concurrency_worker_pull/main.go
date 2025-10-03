package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Task struct {
	ID int
}

func main() {
	numWorker := 10

	numTasks := 100

	tasks := make(chan Task, numTasks)

	for i := 0; i < numWorker; i++ {
		wg.Add(1)
		go worker(i, tasks)

	}

	for i := 0; i < numTasks; i++ {
		tasks <- Task{ID: i}

	}

	close(tasks)
	wg.Wait()

}

func worker(id int, task <-chan Task) {
	defer wg.Done()
	for task := range task {
		fmt.Printf("worker %d is processing task %d \n", id, task.ID)
		time.Sleep(1 * time.Second)

	}
}
