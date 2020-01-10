package worker_pool

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Starter provides functions for starting workers pool
type Starter interface {
	Start()
}

type foreman struct {
	jobs    int
	workers int
}

// Start launches main algorithm
func (f *foreman) Start() {
	jobsChan := make(chan int, 100)
	resChan := make(chan string, 100)
	isCompletedChan := make(chan bool)

	for w := 0; w < f.workers; w++ {
		go worker(w, jobsChan, resChan)
	}

	for j := 0; j < f.jobs; j++ {
		jobsChan <- j
	}
	close(jobsChan)

	go checkWork(f.jobs, resChan, isCompletedChan)

	<-isCompletedChan
	close(resChan)

	fmt.Println("all jobs are completed")
}

func checkWork(jobs int, results <-chan string, completed chan<- bool) {
	for j := 0; j < jobs; j++ {
		fmt.Println(<-results)
	}
	completed <- true
}

func worker(id int, jobs <-chan int, results chan<- string) {
	for j := range jobs {
		fmt.Println("worker id:", id+1, "\tjob:", j+1, "PROCESSING")
		time.Sleep(time.Second * (time.Duration)(rand.Intn(3)+1))
		results <- "worker id: " + strconv.Itoa(id+1) + "\tjob: " + strconv.Itoa(j+1) + " COMPLETED"
	}
}

// NewForeman returns foreman struct
func NewForeman(jobs, workers int) Starter {
	return &foreman{
		jobs:    jobs,
		workers: workers,
	}
}
