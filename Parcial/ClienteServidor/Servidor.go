package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func funcionTrapecio(funcion func(float64) float64, min, max float64) float64 {

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

func main() {
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Por favor ingrese un puerto")
		return
	}
	PORT := ":" + arguments[1]

	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer connection.Close()

	buffer := make([]byte, 1024)

	for {

		n, addr, err := connection.ReadFromUDP(buffer)
		fmt.Print("Mensaje recibido: ", string(buffer[0:n]))

		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Saliendo del servidor UDP")
			return
		}
		fmt.Printf("Mensaje recibido de: %s\n", addr)

		if err != nil {
			fmt.Println(err)
			return
		}
	}

}
