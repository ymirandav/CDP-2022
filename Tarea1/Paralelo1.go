package main

import (
	"fmt"
	"sync"
	"time"
)

func loop(a int) {
	for i := 0; i < 50; i++ {
		fmt.Printf("%d", a)
	}
}
func loop2(a int) {
	for i := 0; i < 50; i++ {
		fmt.Printf("%d", a)
	}
}

func main() {

	var wg sync.WaitGroup

	start := time.Now()

	for i := 0; i < 3; i++ {
		wg.Add(2)

		i := i

		go func() {

			loop(i)
			defer wg.Done()
		}()
		go func() {

			loop2(i)
			defer wg.Done()
		}()
	}

	wg.Wait()
	end := time.Since(start)
	fmt.Println()
	fmt.Printf("Time: %d", end.Nanoseconds())
}
