/* My TCP server */

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	fmt.Println("TCP server is running!")

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	// buffer := make([]byte, 10000)

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		t := time.Now()
		myTime := t.Format(time.RFC3339)
		fmt.Printf("-> %s (len %d) at %s\n", strings.TrimRight(string(netData), "\r\n"), len(netData), myTime)
		if strings.TrimSpace(string(netData)) == "GOODBYE" {
			fmt.Println("Exiting TCP server!")
			return
		}

		data := strings.TrimRight(string(netData), "\r\n") + " (Received)"
		_, err = c.Write([]byte(data))
		if err != nil {
			fmt.Println(err)
			return
		}

	}
}
