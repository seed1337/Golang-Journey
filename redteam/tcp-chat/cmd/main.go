package main

import (
	"fmt"
	"net"
	"sync"
)

var (
	clients      = make(map[int]net.Conn)
	clientsLock  sync.RWMutex
	nextClientID = 1
)

func main() {
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

		clientsLock.Lock()
		clients[nextClientID] = conn
		clientID := nextClientID
		nextClientID++
		clientsLock.Unlock()

		fmt.Printf("Client %d connected.\n", clientID)
		go handleConnection(conn, clientID)
	}
}

func handleConnection(conn net.Conn, clientID int) {
	defer func() {
		conn.Close()
		clientsLock.Lock()
		delete(clients, clientID)
		clientsLock.Unlock()
		fmt.Printf("Client %d disconnected.\n", clientID)
	}()

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("An error occurred on connection for client %d: %v\n", clientID, err)
			break
		}

		msgRecv := string(buf[:n])
		handleMessage(clientID, msgRecv)
	}
}

func handleMessage(clientID int, msg string) {
	fmt.Printf("Message received from client %d: %s\n", clientID, msg)
}
