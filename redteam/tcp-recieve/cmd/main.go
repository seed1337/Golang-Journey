package main

import (
	"fmt"
	"net"
)

func main() {

	ln, err := net.Listen("tcp", ":1337")

	if err != nil {
		println("Couldnt listen on port 1337, please run with sudo permissions")
	}

	fmt.Println("Started server on port 1337")

	for {
		conn, err := ln.Accept()

		if err != nil {
			// handle error
			fmt.Printf("An error occured on connection: %v", err)
		}

		fmt.Printf("New connection from %v\n", conn.RemoteAddr())
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)

		if err != nil {
			fmt.Printf("An error occured on connection: %v", err)
			break
		}

		fmt.Printf("Received %d bytes from %v\n", n, conn.RemoteAddr())
		fmt.Printf("Message: %s\n", string(buf[:n]))
	}
}
