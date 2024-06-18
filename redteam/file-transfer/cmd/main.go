package main

import (
	"file-transfer/internals"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) > 3 {
		fmt.Println("Usage: main <server/client>")
	}

	if len(os.Args) == 1 {
		fmt.Println("Usage: main <server/client>")
		return
	}

	switch os.Args[1] {
	case "server":
		internals.StartServer()
	case "client":
		if len(os.Args) < 3 {
			fmt.Println("Usage: main client <file>")
			return
		}
		internals.StartClient(os.Args[2])
	case "-h", "--help":
		fmt.Println("Usage: main <server>")
		fmt.Println("Usage: main <client> file.txt")
	default:
		fmt.Println("Invalid option, use -h or --help for help")
	}

}
