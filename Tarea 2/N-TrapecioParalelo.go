package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Constructor struct {
	a    float64
	b    float64
	n    int
	area func(float64) float64
}

func (c *Constructor) Trapecio() float64 {

	altura := (c.b - c.a)
	delta := altura / float64(c.n)

	var areaTotal float64
	areaTotal = 0

	a2 := c.a

	var tiempoTotal int64

	for i := 1; i <= c.n; i++ {
		start := time.Now()

		a2 = c.a
		c.a = c.a + delta
		c.b = c.a
		altura2 := c.b - a2

		wg.Add(1)

		baseMenor := funcion(a2)
		baseMayor := funcion(c.b)

		acumulado := ((baseMenor + baseMayor) * altura2) / 2

		areaTotal += acumulado

		end := time.Since(start).Nanoseconds()
		fmt.Println("---------------------------")
		fmt.Println("Area Trapecio (", i, "): ", acumulado)
		// fmt.Println("Tiempo: ", end)

		tiempoTotal += end

		wg.Done()

	}
	return areaTotal
}

func funcion(x float64) float64 {
	return (x*x + 1) / 2
}

func main() {

	NumeroTrapecios := 100000

	wg := new(sync.WaitGroup)

	f := func(x float64) float64 {
		return (x*x + 1) / 2
	}

	c := Constructor{
		area: f, a: 5, b: 20, n: NumeroTrapecios,
	}

	start := time.Now()
	area := c.Trapecio()
	end := time.Since(start)

	wg.Wait()

	fmt.Println("---------------------------")
	fmt.Println("Area Total: ", area)
	fmt.Println("Tiempo Total: ", end.Nanoseconds())
	fmt.Println("---------------------------")
}
