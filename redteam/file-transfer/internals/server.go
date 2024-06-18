package internals

import (
	"encoding/base64"
	"fmt"
	"net"
	"os"

	"github.com/google/uuid"
)

func StartServer() {
	ln, err := net.Listen("tcp", ":1337")
	if err != nil {
		fmt.Printf("Couldn't start server, %v\n", err)
		return
	}

	fmt.Println("Server started on port 1337")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("An error occurred on connection: %v\n", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 1024)

		n, err := conn.Read(buf)
		if err != nil {
			break
		}

		msgRecv := string(buf[:n])

		data, err := base64.StdEncoding.DecodeString(msgRecv)

		if err != nil {
			fmt.Println("error:", err)
			return
		}

		fileName := uuid.NewString()
		os.WriteFile(fileName, data, 0644)
		fmt.Printf("File written: %v", fileName)

	}
}
