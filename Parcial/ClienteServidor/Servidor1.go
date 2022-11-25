package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func funcionTrapecio(f func(float64) float64, min, max float64) float64 {

	numTrapecios := 10

	h := (max - min)
	delta := h / float64(numTrapecios)

	var ATOTAL float64

	ATOTAL = 0

	min2 := min

	for i := 1; i <= numTrapecios; i++ {
		min2 = min
		min = min + delta
		max = min

		h2 := max - min2

		baseMenor := funcion(min2)

		baseMayor := funcion(max)

		acum := ((baseMenor + baseMayor) * h2) / 2

		ATOTAL += acum

	}

	return ATOTAL
}

func handler(conn net.Conn) {
	for {
		m, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("Connection closed")
				conn.Close()
				return
			}
			fmt.Println("Error reading from connection", err)
			return
		}
		_, err = conn.Write([]byte(m))
		if err != nil {
			fmt.Println("Error writing to connection")
			return
		}
		fmt.Printf("%v %q\n", conn.RemoteAddr(), m)
	}
}

func funcion(x float64) float64 {
	return (x*x + 1) / 2
}

func main() {
	fmt.Println("Listening on port 8080")

	ln, _ := net.Listen("tcp", ":8080")

	for {
		conn, _ := ln.Accept()
		fmt.Println("Connection accepted")
		go handler(conn)
		// float64 a; := go funcionTrapecio(funcion, min, max)
		// go funcionTrapecio(min, max)
		fmt.Println(funcionTrapecio(funcion, 5, 20))
	}
}
