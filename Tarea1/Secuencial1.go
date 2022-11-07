package main

import (
	"fmt"
	"time"
)

func loop(a int) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%d", a)
	}
}

func main() {
	start := time.Now()
	loop(1)
	loop(2)
	loop(3)
	end := time.Since(start)
	fmt.Println()
	fmt.Printf("Time: %d", end.Nanoseconds())
}
