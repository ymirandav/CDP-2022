package main

import (
	"fmt"
	"time"
)

func runner1() {
	fmt.Print("\nPrimer corredor")
}

func runner2() {
	fmt.Print("\nSegundo corredor")
}

func runner3() {
	fmt.Print("\nTercer corredor")
}

func execute() {
	runner1()
	runner2()
	runner3()
}

func main() {
	start := time.Now()

	execute()

	end := time.Since(start)
	fmt.Println()
	fmt.Printf("Time: %d", end.Nanoseconds())
}
