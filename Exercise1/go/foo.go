// Use `go run foo.go` to run your program

package main

import (
    . "fmt"
    "runtime"
    "time"
)

var i = 0

func incrementing() {
    //increment i 1000000 times
    for a := 0; a <= 1000000; a++ {
      i++;
    }
}

func decrementing() {
  //increment i 1000000 times
  for a := 0; a <= 1000000; a++ {
    i--;
  }
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())

    //spawn both functions as goroutines
    go incrementing()
    go decrementing()

    // We have no way to wait for the completion of a goroutine (without additional syncronization of some sort)
    // We'll come back to using channels in Exercise 2. For now: Sleep.
    time.Sleep(500*time.Millisecond)
    Println("The magic number is:", i)
}
