package main

import (
	"fmt"
	"sync"
	"time"
)

//https://talks.golang.org/2012/concurrency.slide#1

func main() {
	blockingExample()
	// closingChannelExample()
	// rangeOverChannelExample()
	// waitingGroupExample()
	// doneChannelExample()

	// bufferedChannelsExample()
	// inputOnlyAndOutputOnlyChanneslExample()
	// // rangeOverChannelExample()
	// selectOverChannelsExample()
	// sendFunctionIntoChannelExample()
}

//read input/output for channels from right to left
//output = <- aChannel <- input

func blockingExample() {
	intChan := make(chan int)

	intChan <- 5
	// go func() { intChan <- 5 }()
	// go func() {
	// 	fmt.Println("staring goroutine")
	// 	intChan <- 5
	// 	fmt.Println("finishing goroutine")
	// }()

	fmt.Println("after input")

	output := <-intChan

	fmt.Println("after output1")

	fmt.Println(output)

	// output = <-intChan

	// fmt.Println("after output2")

	// fmt.Println(output)
}

func closingChannelExample() {

	{
		intChan := make(chan int)
		defer close(intChan)

		// intChan <- 5
		go func() { intChan <- 5 }()

		value, isOpen := <-intChan

		fmt.Println(value, isOpen)
	}

	{
		intChan := make(chan int)
		close(intChan)

		// go func() { intChan <- 5 }()
		// intChan <- 5

		value, isOpen := <-intChan

		fmt.Println(value, isOpen)
	}

}

func waitingGroupExample() {
	complicatedCalculation := func() {
		time.Sleep(time.Second)
	}

	worker := func(id int, wg *sync.WaitGroup) {

		defer wg.Done()

		fmt.Printf("Worker %d starting\n", id)

		complicatedCalculation()

		fmt.Printf("Worker %d done\n", id)
	}

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
}

//`watch -n 0.3  go run main.go`
func doneChannelExample() {
	jobs := make(chan int)
	// jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}

func bufferedChannelsExample() {

}

func rangeOverChannelExample() {
	//with buffer
	{
		queue := make(chan string, 2)
		queue <- "one"
		queue <- "two"
		close(queue)

		for elem := range queue {
			fmt.Println(elem)
		}
	}

	//...
	// {
	// 	secChan := make(chan bool)

	// 	go func(c chan bool) {
	// 		count := 0
	// 		for {
	// 			time.Sleep(time.Second)
	// 			c <- true
	// 			count++
	// 			if count == 5 {
	// 				close(c)
	// 			}
	// 		}
	// 	}(secChan)

	// 	for elem := range secChan {
	// 		fmt.Println(elem)
	// 	}
	// }
}

func inputOnlyAndOutputOnlyChanneslExample() {}

func selectOverChannelsExample() {}

func sendFunctionIntoChannelExample() {}
