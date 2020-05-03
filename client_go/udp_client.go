/* MY UDP Client */

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
		fmt.Println("Please provide a host:port string")
		return
	}

	CONNECT := arguments[1]

	s, err := net.ResolveUDPAddr("udp4", CONNECT)
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("The UDP server is %s\n", c.RemoteAddr().String())
	defer c.Close()

	for {
		fmt.Print(">> ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		data := []byte(text)
		_, err = c.Write(data)

		if strings.TrimSpace(string(data)) == "GOODBYE" {
			fmt.Println("Existing UDP client!")
			return
		}

		if strings.TrimSpace(string(data)) == "SEND" {
			fmt.Println("Sending larger packet")
			a := make([]byte, 5000)
			_, err = c.Write(a)
		}

		if err != nil {
			fmt.Println(err)
			return
		}

		buffer := make([]byte, 1024)
		n, _, err := c.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("$$ %s\n", string(buffer[0:n]))
	}
}
