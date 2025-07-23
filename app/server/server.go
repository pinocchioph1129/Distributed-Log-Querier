package server

import (
	"fmt"
	"log"
	"net"
	"github.com/google/shlex"
	"os/exec"
)

func Run() {
	fmt.Println("Listening on port 8080, hello!")
	ln, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}


func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Handling connection from", conn.RemoteAddr())

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from connection:", err)
		return
	}

	fmt.Printf("Received data from %s: %s\n", conn.RemoteAddr(), string(buffer[:n]))
	commandLine := string(buffer[:n])
	var commandParts []string
	commandParts, err = shlex.Split(commandLine)
	if err != nil {
		fmt.Println("Error parsing command:", err)
		return
	}
	for i, part := range commandParts {
		fmt.Printf("Command part %d: %s\n", i, part)
	}
	filepath := commandParts[len(commandParts)-1]
	filepath = fmt.Sprintf("./log_files/%s", filepath)
	if len(commandParts) > 0 {
		commandParts = commandParts[:len(commandParts)-1]
	}
	commandParts = append(commandParts, filepath)
	stdout, err := exec.Command("/bin/grep", commandParts[1:]...).Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok && exitError.ExitCode() == 1 {
			fmt.Println("No matches found")
			return
		}
		fmt.Println("Error executing command:", err)
		return
	}
	fmt.Printf("Command output: %s\n", stdout)
	_, err = conn.Write(stdout)
	if err != nil {
		fmt.Println("Error writing to connection:", err)
		return
	}
	fmt.Println("Data sent to client successfully")
}
