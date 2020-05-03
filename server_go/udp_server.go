/* MY UDP Server */

package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}
	fmt.Println("UDP server is running!")

	PORT := ":" + arguments[1]

	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	buffer := make([]byte, 10000)

	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("->%s (len %d)\n", string(buffer[0:n-1]), n)
		if strings.TrimSpace(string(buffer[0:n-1])) == "GOODBYE" {
			fmt.Println("Exiting UDP server!")
			return
		}

		data := string(buffer[0:n-1]) + " (Received)"
		_, err = conn.WriteToUDP([]byte(data), addr)
		if err != nil {
			fmt.Println(err)
			return
		}

	}
}
