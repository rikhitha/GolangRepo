package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

func write(ch chan int) {

	for u := 0; u < 5; u++ {
		ch <- u

		fmt.Println("wrote:", u, "to ch channel")
	}

	close(ch)

}

func callfunc(i int, waitgrp *sync.WaitGroup) {
	fmt.Println("go routine started", i)

	waitgrp.Done()

}

func getnumber() int {
	var num int
	fmt.Println("enter number")
	fmt.Scanln(&num)

	return num
}

type Job struct {  
    id       int
    randomno int
}
type Result struct {  
    job         Job
    sumofdigits int
}

var jobs = make(chan Job, 10)  
var results = make(chan Result, 10)

func digits(number int) int {  
    sum := 0
    no := number
    for no != 0 {
        digit := no % 10
        sum += digit
        no /= 10
    }
    time.Sleep(2 * time.Second)
    return sum
}
func worker(wg *sync.WaitGroup) {  
    for job := range jobs {
        output := Result{job, digits(job.randomno)}
        results <- output
    }
    wg.Done()
}
func createWorkerPool(noOfWorkers int) {  
    var wg sync.WaitGroup
    for i := 0; i < noOfWorkers; i++ {
        wg.Add(1)
        go worker(&wg)
    }
    wg.Wait()
    close(results)
}
func allocate(noOfJobs int) {  
    for i := 0; i < noOfJobs; i++ {
        randomno := rand.Intn(999)
        job := Job{i, randomno}
        jobs <- job
    }
    close(jobs)
}
func result(done chan bool) {  
    for result := range results {
        fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
    }
    done <- true
}

func main() {

	fmt.Println("choose \n 1. Buffered channels  \n 2. Dead lock  \n 3. Lenth vs capacity \n 4. Wait group")
	switch num := getnumber(); { //num is not a constant
	case num == 1:

		fmt.Println("Buffered channels")

		chan1 := make(chan int, 3)

		chan1 <- 100

		chan1 <- 200

		chan1 <- 300

		fmt.Println("values", <-chan1, <-chan1, <-chan1)

		chan2 := make(chan int, 2)

		go write(chan2)

		time.Sleep(3 * time.Second)

		for b := range chan2 {
			fmt.Println("read value ", b, "from chan2")
			time.Sleep(3 * time.Second)
		}
	case num == 2:
		fmt.Println("Deadlock")

		chan3 := make(chan string, 1)

		chan3 <- "Rikhitha"

		chan3 <- "Manoj"

		chan3 <- "Kumar"

		fmt.Println("Firstname :", <-chan3)

		fmt.Println("Middlename :", <-chan3)

		fmt.Println("Lastname :", <-chan3)

	case num == 3:
		fmt.Println("Length vs capacity")

		channel := make(chan int, 3)

		channel <- 21

		channel <- 23

		fmt.Println("capacity of channel : ", cap(channel))

		fmt.Println("length of channel : ", len(channel))

		fmt.Println("read from channel : ", <-channel)

		fmt.Println("length of channel after reading :", len(channel))

	case num == 4:
		fmt.Println("Waitgroup")

		var waitgrp sync.WaitGroup

		for g := 0; g < 4; g++ {
			waitgrp.Add(1)
			go callfunc(g, &waitgrp)
		}

		waitgrp.Wait()

		fmt.Println("all go routines executed ,,, in main function ")

	case num == 5:
		fmt.Println("Worker pool")

		startTime := time.Now()
    noOfJobs := 100
    go allocate(noOfJobs)
    done := make(chan bool)
    go result(done)
    noOfWorkers := 10
    createWorkerPool(noOfWorkers)
    <-done
    endTime := time.Now()
    diff := endTime.Sub(startTime)
    fmt.Println("total time taken ", diff.Seconds(), "seconds")
	}

}
