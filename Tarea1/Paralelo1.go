package main

import (
	"fmt"
	"runtime"
	"time"
)

var quit chan int = make(chan int)

func loop(a int) {
	for i := 0; i < 100; i++ {
		runtime.Gosched()
		fmt.Printf("%d", a)
	}
	quit <- 0
}

func main() {

	start := time.Now()
	go loop(1)
	go loop(2)
	go loop(3)

	for i := 0; i < 3; i++ {
		<-quit
	}
	end := time.Since(start)
	fmt.Println()
	fmt.Printf("Time: %d", end.Nanoseconds())
}
