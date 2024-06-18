package main

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
)

func main() {

	ipv4_regex := `^(((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4})`

	switch os.Args[1] {
	case "-h", "--help":
		println("Usage: main <ip> <port>")
		println("-e | run a program when connecting.")

	default:
		fmt.Println("Usage: main <ip> <port>")
	}

	switch {
	case len(os.Args) < 2:
		fmt.Println("Usage: main <ip> <port>")
	case len(os.Args) < 3:
		fmt.Println("Usage: main <ip> <port>")

	default:

		// checks before connection

		match, _ := regexp.MatchString(ipv4_regex, os.Args[1])
		if !match {
			fmt.Println("Invalid IP address")
			return
		}

		port, err := strconv.Atoi(os.Args[2])
		if err != nil || port < 0 || port > 65535 {
			fmt.Println("Invalid port number")
			return
		}

		// connect to ip and port.

		fmt.Println("Connecting...")
		connect(os.Args[1], port)
	}
}

func connect(ip string, port int) {

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		fmt.Println("Error connecting, is the server running? Error:", err)
		return
	}

	fmt.Printf("Connected: %v", conn)
	// status, err := bufio.NewReader(conn).ReadString('\n')
}

func runProgram(conn net.Conn, program string) error {

}
