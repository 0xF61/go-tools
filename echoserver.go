package main

import (
	"bufio"
	//"io"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()

	/*
	 *  for {
	 *    b := make([]byte, 512)
	 *    size, err := conn.Read(b[0:])
	 *    if err == io.EOF {
	 *      log.Println("Client Disconnected")
	 *      break
	 *    }
	 *
	 *    if err != nil {
	 *      log.Println("Unexpected Error")
	 *    }
	 *
	 *    log.Printf("Received %d bytes: %s", size, string(b))
	 *
	 *    log.Printf("Writing Data")
	 *    if _, err := conn.Write(b[0:size]); err != nil {
	 *      log.Fatalln("Unable to write data")
	 *    }
	 *  }
	 */

	for {

		reader := bufio.NewReader(conn)
		s, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln("Unable to read data")
		}
		log.Println("Read %d bytes: %s", len(s), s)

		log.Println("Writing Data")
		writer := bufio.NewWriter(conn)
		if _, err := writer.WriteString(s); err != nil {
			log.Fatalln("Unable to write data")
		}
		log.Println("Read %d bytes: %s", len(s), s)

		writer.Flush()
	}

}

func main() {
	listener, err := net.Listen("tcp", ":1337")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:1337")

	for {
		conn, err := listener.Accept()
		log.Println("Received connection")

		if err != nil {
			log.Fatalln("Unable to Accept Connection")
		}

		go echo(conn)
	}
}
