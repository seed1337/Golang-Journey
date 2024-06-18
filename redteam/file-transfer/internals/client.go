package internals

import (
	"encoding/base64"
	"fmt"
	"net"
	"os"
)

func StartClient(file string) {
	conn, err := net.Dial("tcp", ":1337")
	if err != nil {
		fmt.Printf("Couldn't connect to server, please start it. Error: %v\n", err)
		return
	}

	fmt.Printf("Connected to server: %v\n", conn.RemoteAddr().String())
	fmt.Println("File to send:", file)
	fmt.Println("Sending file...")

	sendFile(conn, file)
}

func sendFile(conn net.Conn, filename string) {
	defer conn.Close()

	f, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Couldn't read file, %v\n", err)
		return
	}

	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	b64 := base64.StdEncoding.EncodeToString(f)
	if err != nil {
		fmt.Printf("Error writing to encoder: %v\n", err)
		return
	}

	conn.Write([]byte(b64))
	fmt.Println("File sent.")

	encoder.Close()

}
