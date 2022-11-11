package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randSleep(name string, limit int, sleep int) {
	for i := 1; i <= limit; i++ {
		fmt.Println(name, rand.Intn(i))
		time.Sleep(time.Duration(sleep * int(time.Second)))

	}

}
func main() {
	start := time.Now()

	randSleep("a:", 4, 2)
	randSleep("b:", 4, 2)

	end := time.Since(start)
	fmt.Println()
	fmt.Printf("Time: %d", end.Nanoseconds())
}
