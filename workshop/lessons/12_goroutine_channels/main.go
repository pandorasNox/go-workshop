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
	// bufferedChannelsExample()
	// rangeOverChannelExample()
	// waitingGroupExample()
	// doneChannelExample()
	// inputOnlyAndOutputOnlyChanneslExample()
	// selectOverChannelsExample()
	// sendFunctionIntoChannelExample()
	// actorModelExample()
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

	// intChan := make(chan int)
	// defer close(intChan)

	// // intChan <- 5
	// go func() { intChan <- 5 }()

	// value, isOpen := <-intChan

	// fmt.Println(value, isOpen)

	//
	//
	//

	// // close channel direct

	// intChan := make(chan int)
	// close(intChan)

	// // go func() { intChan <- 5 }()
	// // intChan <- 5

	// value, isOpen := <-intChan

	// fmt.Println(value, isOpen)
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
			j, isOpen := <-jobs
			if isOpen {
				//process job
				fmt.Println("received job", j)
				continue
			}

			fmt.Println("no more jobs to recive")
			done <- true
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
	// byteChan := make(chan byte, 2)
	// byteChan <- 107
	// byteChan <- 217
	// fmt.Println(<-byteChan)
	// fmt.Println(<-byteChan)
	// fmt.Println(<-byteChan)

	// {
	// 	//closed buffered channels
	// 	byteChan := make(chan byte)
	// 	var val byte
	// 	var isOpen bool
	// 	go func() {
	// 		fmt.Println("before")
	// 		byteChan <- 107
	// 		fmt.Println("after")
	// 	}()

	// 	time.Sleep(time.Millisecond * 100)
	// 	runtime.Gosched()
	// 	time.Sleep(time.Millisecond * 1000)
	// 	// runtime.Gosched()
	// 	// runtime.Gosched()
	// 	// runtime.Gosched()

	// 	close(byteChan)
	// 	val, isOpen = <-byteChan
	// 	fmt.Println(val, isOpen)

	// 	time.Sleep(time.Millisecond * 100)
	// }

	//
	//
	//

	// // close _buffered_ channel after a send

	// intChan := make(chan int, 1)

	// go func() { intChan <- 5 }()

	// runtime.Gosched()
	// time.Sleep(time.Millisecond * 100)

	// close(intChan)

	// value, isOpen := <-intChan
	// fmt.Println(value, isOpen)

	// value, isOpen = <-intChan
	// fmt.Println(value, isOpen)
}

func rangeOverChannelExample() {
	//with buffer
	// {
	// 	queue := make(chan string, 2)
	// 	queue <- "one"
	// 	queue <- "two"
	// 	close(queue)

	// 	// var val string
	// 	// var isOpen bool
	// 	// val, isOpen = <-queue
	// 	// val, isOpen = <-queue
	// 	// fmt.Println(val, isOpen)

	// 	//range as long as channel is open or not empty (buffered).
	// 	for elem := range queue {
	// 		fmt.Println(elem)
	// 	}
	// }

	//range with channel which is not buffered  is blocking
	// {
	// 	queue := make(chan string)
	// 	go func() { queue <- "one" }()
	// 	go func() { queue <- "two" }()
	// 	go func() {
	// 		for {
	// 			queue <- "endless"
	// 			time.Sleep(time.Millisecond * 100)
	// 		}
	// 	}()
	// 	defer close(queue)

	// 	// var val string
	// 	// var isOpen bool
	// 	// val, isOpen = <-queue
	// 	// val, isOpen = <-queue
	// 	// fmt.Println(val, isOpen)

	// 	//range as long as channel is open or not empty (buffered).
	// 	for elem := range queue {
	// 		fmt.Println(elem)
	// 	}
	// }

	//range, chan with buffer, but not closed... is crashing
	//range only stops trying to recive when channle is closed
	// {
	// 	queue := make(chan string, 2)
	// 	queue <- "one"
	// 	queue <- "two"
	// 	close(queue)

	// 	// var val string
	// 	// var isOpen bool
	// 	// val, isOpen = <-queue
	// 	// val, isOpen = <-queue
	// 	// fmt.Println(val, isOpen)

	// 	//range as long as channel is open or not empty (buffered).
	// 	for elem := range queue {
	// 		fmt.Println(elem)
	// 	}
	// }

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

func inputOnlyAndOutputOnlyChanneslExample() {
	// //input only
	// {
	// 	getIntInOnlyChan := func() chan<- int {
	// 		mychan := make(chan int)
	// 		go func() {
	// 			fmt.Println(<-mychan)
	// 		}()
	// 		return mychan
	// 	}
	// 	myIntChan := getIntInOnlyChan()

	// 	go func() {
	// 		myIntChan <- 0
	// 	}()

	// 	// fmt.Println(<-myIntChan)
	// 	time.Sleep(time.Microsecond * 100)
	// }

	// //output only
	// {
	// 	getIntOutOnlyChan := func() <-chan int {
	// 		mychan := make(chan int)
	// 		go func() {
	// 			mychan <- 0
	// 		}()
	// 		return mychan
	// 	}

	// 	myIntChan := getIntOutOnlyChan()

	// 	fmt.Println(<-myIntChan)

	// 	// go func() {
	// 	// 	myIntChan <- 5
	// 	// }()
	// }
}

type joined struct {
	intJob int
	strJob string
}

func selectOverChannelsExample() {
	//
	//
	//

	// // without select

	// intJobs := make(chan int)
	// strJobs := make(chan string)
	// joinedJobs := make(chan joined)

	// //send routines
	// go func() {
	// 	intJobs <- 5
	// }()

	// go func() {
	// 	strJobs <- "AA"
	// }()

	// // recive + merge goroutines
	// go func() {
	// 	joinedJobs <- joined{<-intJobs, ""}
	// }()
	// go func() {
	// 	joinedJobs <- joined{0, <-strJobs}
	// }()

	// go func() {
	// 	fmt.Println(<-joinedJobs)
	// 	fmt.Println(<-joinedJobs)
	// }()

	// time.Sleep(time.Microsecond * 500)

	//
	//
	//

	// // using select

	// {
	// 	intJobs := make(chan int)
	// 	strJobs := make(chan string)

	// 	//send routines
	// 	go func() {
	// 		intJobs <- 5
	// 	}()

	// 	go func() {
	// 		strJobs <- "AA"
	// 	}()

	// 	time.Sleep(time.Microsecond * 500)

	// 	for {
	// 		select {
	// 		case x := <-intJobs:
	// 			fmt.Println("select intJob:", x)
	// 		case y := <-strJobs:
	// 			fmt.Println("select strJobs:", y)
	// 		}

	// 		time.Sleep(time.Microsecond * 500)
	// 	}

	// 	time.Sleep(time.Microsecond * 500)
	// }

	//
	//
	//

	// // select over incoming channel
	// todo

}

func sendFunctionIntoChannelExample() {
	// afunc := func(fn func()) { fmt.Println("afunc"); fn() }
	// afunc(func() {
	// 	fmt.Println("fn")
	// })

	x := 6

	myFnChan := make(chan func())

	go func() {
		myFnChan <- func() { fmt.Println("I was send for something", x) }
	}()

	go func() {
		myFnChan <- func() { y := 7; fmt.Println("another send for something", x, y) }
	}()

	afunc := <-myFnChan
	afunc()

	afunc = <-myFnChan
	afunc()
}

//actor model
//https://youtu.be/yCbon_9yGVs?t=1303

type config string
type ConfigActor struct {
	channel chan func(currentState *config)
	state   config
}

func (cfgActor *ConfigActor) Init(initalState config) {
	channel := make(chan func(state *config))
	cfgActor.channel = channel
	cfgActor.state = initalState
}

func (cfgActor *ConfigActor) Read() (state config) {
	// returnCh := make(chan Config)
	// defer close(returnCh)

	// that example does only randomly works, bec sending into a channel is only sending a message
	// when the function is executed, there is no guarantee about that or that it gets even executed
	// reasons:
	// - sheduling
	// - conditions on the other side after picking up the message what happens with the message
	// - other....
	//
	// var cpState config
	// cfgActor.channel <- func(currentState *config) {
	// 	cpState = *currentState
	// }
	// return cpState

	//better example
	returnCh := make(chan config)
	defer close(returnCh)
	cfgActor.channel <- func(currentState *config) {
		returnCh <- *currentState
	}
	var cpState config
	cpState = <-returnCh
	return cpState
}

func (cfgActor *ConfigActor) Run() {
	go func() {
		//actor loop
		for {
			fn, isOpen := <-cfgActor.channel
			if !isOpen {
				break
			}
			time.Sleep(time.Millisecond * 100)

			fn(&cfgActor.state)
			// time.Sleep(time.Millisecond * 100)
		}
	}()
}

func (cfgActor *ConfigActor) Stop() {
	close(cfgActor.channel)

	//make here a clean shutdown handling
	time.Sleep(time.Microsecond * 100)
}

func (cfgActor *ConfigActor) InitAndRun(initState config) {
	cfgActor.Init(initState)
	cfgActor.Run()
}

func (cfgActor *ConfigActor) Write(cfg config) {
	cfgActor.channel <- func(currentState *config) {
		*currentState = cfg
	}
}

func actorModelExample() {
	cfgActor := ConfigActor{}
	cfgActor.InitAndRun("my init state")
	defer cfgActor.Stop()

	fmt.Println(cfgActor.Read())
	cfgActor.Write("changed")
	fmt.Println(cfgActor.Read())

	// cfgActor.send("message")
}
