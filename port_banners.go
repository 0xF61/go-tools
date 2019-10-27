package main

import (
	"log"
	"net"
	"strconv"
	"time"
)

var ipToScan = "192.168.1.3"
var minPort = 1
var maxPort = 65535

func main() {
	activeThreads := 0
	doneChannel := make(chan bool)

	for port := minPort; port <= maxPort; port++ {
		go grabBanner(ipToScan, port, doneChannel)
		activeThreads++
	}

	for activeThreads > 0 {
		<-doneChannel
		activeThreads--
	}
}

func grabBanner(ip string, port int, doneChannel chan bool) {
	connection, err := net.DialTimeout("tcp", ip+":"+strconv.Itoa(port), time.Second*10)
	if err != nil {
		doneChannel <- true
		return
	}

	// See if servers offers anything to read
	buffer := make([]byte, 4096)
	connection.SetReadDeadline(time.Now().Add(time.Second * 5))
	// Set timeout
	numBytesRead, err := connection.Read(buffer)
	if err != nil {
		doneChannel <- true
		return
	}
	log.Printf("Banner from port: %d\n%s\n", port, buffer[0:numBytesRead])
	doneChannel <- true
}
