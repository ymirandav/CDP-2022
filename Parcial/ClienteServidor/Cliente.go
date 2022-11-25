package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Por favor ingrese un host y puerto de la siguiente manera:  host:port ")
		return
	}
	CONNECT := arguments[1]

	s, err := net.ResolveUDPAddr("udp4", CONNECT)
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("El servidor UDP es %s\n", c.RemoteAddr().String())

	defer c.Close()

	for {

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Mensaje a enviar:  ")
		text, _ := reader.ReadString('\n')
		data := []byte(text)
		_, err = c.Write(data)

		if strings.TrimSpace(string(data)) == "STOP" {
			fmt.Println("Saliendo del cliente UDP")
			return
		}

		if err != nil {
			fmt.Println(err)
			return
		}

	}
}
