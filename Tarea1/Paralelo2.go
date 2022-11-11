package main

import (
	"fmt"
	"sync"
	"time"
)

func runner1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print("\nPrimer corredor")

}

func runner2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print("\nSegundo corredor")
}

func runner3(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print("\nTercer corredor")
}

func execute() {
	wg := new(sync.WaitGroup)
	wg.Add(3)

	go runner1(wg)
	go runner2(wg)
	go runner3(wg)

	wg.Wait()
}

func main() {
	start := time.Now()

	execute()

	end := time.Since(start)
	fmt.Println()
	fmt.Printf("Time: %d", end.Nanoseconds())
}
