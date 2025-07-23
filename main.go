package main

import (
	"fmt"
	"flag"
	"distributed-log-querier/app/client"
	"distributed-log-querier/app/server"
)

func usage() {
	fmt.Println("Usage: go run main.go [options]")
	fmt.Println("Options:")
	fmt.Println("  -h, --help\tShow this help message")
	fmt.Println("  -d, --device\tSpecify the device to connect to (server/client)")
}


func main() {
	// put this in another function to avoid infinite loop
	var help bool
	var device string

	flag.BoolVar(&help, "h", false, "Show help message")
    flag.StringVar(&device, "d", "client", "Specify the device to connect to (server/client)")
	flag.Parse()

	switch device {
	case "server":
		server.Run()
		return
	case "client":	
		client.Run()
		return
	default:
		fmt.Println("Invalid device specified. Use 'server' or 'client'.")
		usage()
		return
	}
}

// wg := sync.WaitGroup{}
// defer wg.Wait()