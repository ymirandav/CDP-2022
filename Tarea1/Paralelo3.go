package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func randSleep(wg *sync.WaitGroup, name string, limit int, sleep int) {
	defer wg.Done()
	for i := 1; i <= limit; i++ {
		fmt.Println(name, rand.Intn(i))
		time.Sleep(time.Duration(sleep * int(time.Second)))

	}

}
func main() {
	start := time.Now()

	wg := new(sync.WaitGroup)
	wg.Add(2)
	go randSleep(wg, "a:", 4, 2)
	go randSleep(wg, "b:", 4, 2)
	wg.Wait()

	end := time.Since(start)
	fmt.Println()
	fmt.Printf("Time: %d", end.Nanoseconds())

}
