package worker_pool

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func Start(jobsq int, workersq int) {
	jobs := make(chan int, 100)
	results := make(chan string, 100)
	completed := make(chan bool)

	for w := 0; w < workersq; w++ {
		go worker(w, jobs, results)
	}

	for j := 0; j < jobsq; j++ {
		jobs <- j
	}
	close(jobs)

	go foreman(jobsq, results, completed)

	<-completed

	fmt.Println("all jobs are completed")
}

func foreman(jobsq int, results <-chan string, completed chan<- bool) {
	for j := 0; j < jobsq; j++ {
		fmt.Println(<-results)
	}
	completed <- true
}
func worker(id int, jobs <-chan int, results chan<- string) {
	for j := range jobs {
		fmt.Println("worker id:", id+1, "\tjob:", j+1, "processing")
		time.Sleep(time.Second * (time.Duration)(rand.Intn(3)+1))
		results <- "worker id: " + strconv.Itoa(id+1) + "\tjob: " + strconv.Itoa(j+1) + " COMPLETED"
	}
}
