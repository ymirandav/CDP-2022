package main

import (
	"fmt"
	"time"
)

func worker(jobs chan int, results chan float64) {
	f := func(x float64) float64 {
		return ((x*x + 1) / 2)
	}

	for n := range jobs {
		results <- Trapecio(f, 5, 20, n)
	}
}

func Trapecio(f func(float64) float64, min, max float64, n int) float64 {
	altura := (max - min)
	delta := altura / float64(n)
	var areaTotal float64
	areaTotal = 0
	min2 := min

	for i := 1; i <= n; i++ {

		min2 = min
		min = min + delta
		max = min
		altura2 := max - min2

		baseMenor := f(min2)
		baseMayor := f(max)

		acumulado := ((baseMenor + baseMayor) * altura2) / 2

		areaTotal += acumulado

	}
	return areaTotal
}

func main() {
	numTrapecios := 100000
	jobs := make(chan int, numTrapecios)
	results := make(chan float64, numTrapecios)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	numTrapecios = numTrapecios + 1
	startMain := time.Now()

	for i := 1; i < numTrapecios; i++ {
		start := time.Now()

		jobs <- i

		end := time.Since(start)

		fmt.Println("-----------------------------------------")
		fmt.Println("Area Trapecio (", i, "): ", <-results)
		fmt.Println("Tiempo: ", end.Nanoseconds())
	}

	close(jobs)

	endMain := time.Since(startMain)
	fmt.Println("")
	fmt.Println("Tiempo total: ", endMain.Nanoseconds())
	fmt.Println("")
}
