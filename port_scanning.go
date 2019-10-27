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
		go testTCPConnection(ipToScan, port, doneChannel) // <-- go threads
		activeThreads++
	}

	for activeThreads > 0 {
		<-doneChannel
		activeThreads--
	}
}

func testTCPConnection(ip string, port int, doneChannel chan bool) {
	_, err := net.DialTimeout("tcp", ip+":"+strconv.Itoa(port), time.Second*10)
	if err == nil {
		log.Printf("Host %s has open port: %d\n", ip, port)
	}
	doneChannel <- true
}
