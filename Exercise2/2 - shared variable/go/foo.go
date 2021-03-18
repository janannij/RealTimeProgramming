// Use `go run foo.go` to run your program

package main

import (
	"fmt"
	"runtime"
)

func number_server(add <-chan int, sub <-chan int, read chan<- int) {
	var number = 0

	for {
		select {
		case <-add:
			number++
		case <-sub:
			number--
		case read <- number:
		}
	}
}

func incrementer(add chan<- int, finished chan<- bool) {
	for j := 0; j < 1000000; j++ {
		add <- 1
	}
	//TODO: signal that the goroutine is finished
	finished <- true
}

func decrementer(sub chan<- int, finished chan<- bool) {
	for j := 0; j < 1000000+1; j++ {
		sub <- 1
	}
	//TODO: signal that the goroutine is finished
	finished <- true
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// TODO: Construct the remaining channels
	read := make(chan int)
	add := make(chan int)
	sub := make(chan int)
	add_finished := make(chan bool)
	sub_finished := make(chan bool)

	// TODO: Spawn the required goroutines.
	go number_server(add, sub, read)
	go incrementer(add, add_finished)
	go decrementer(sub, sub_finished)

	// TODO: block on finished from both "worker" goroutines
	<-add_finished
	<-sub_finished
	fmt.Println("The magic number is:", <-read)
}
