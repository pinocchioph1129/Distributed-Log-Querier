package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"github.com/google/shlex"
	"strings"
	"sync"
	"io"
)

var localaddresses = []string{
	"host.docker.internal:8081",
	"host.docker.internal:8082",
	"host.docker.internal:8083",
	"host.docker.internal:8084",
	"host.docker.internal:8085",
	"host.docker.internal:8086",
	"host.docker.internal:8087",
	"host.docker.internal:8088",
	"host.docker.internal:8089",
	"host.docker.internal:8090",
}

var line_countlock sync.Mutex
var total_line int


func runGrepOnAllLogFiles(grepCommand string) {
    wg := sync.WaitGroup{}
    for i := 0; i < len(localaddresses); i++ {
        conn, err := net.Dial("tcp", localaddresses[i])
        if err != nil {
            fmt.Printf("Error connecting to %s: %v\n", localaddresses[i], err)
            continue
        }
        wg.Add(1)
        go func(address string, index int, conn net.Conn) {
            defer wg.Done()
            defer conn.Close()

            grepFile := fmt.Sprintf("machine.%d.log", index+1)
            fullCommand := fmt.Sprintf("%s %s", grepCommand, grepFile)
            _, err := conn.Write([]byte(fullCommand))
            if err != nil {
                fmt.Printf("Error sending command to %s: %v\n", address, err)
                return
            }
            fmt.Printf("Data sent to %s successfully\n", address)

            buffer := make([]byte, 4096)
            var results string
            for {
                n, err := conn.Read(buffer)
                if err != nil {
                    if err == io.EOF {
                        // End of data stream
                        break
                    }
                    fmt.Printf("Error reading from %s: %v\n", address, err)
                    return
                }
                results += string(buffer[:n])
            }
			line_countlock.Lock()
			line_count := strings.Count(results, "\n")
			total_line += line_count
			fmt.Printf("Received data from %s: %d lines\n", address, line_count)
			line_countlock.Unlock()

            // Save results to file
            SaveResultsToFile(results, fullCommand, &wg)
        }(localaddresses[i], i, conn)
    }
    wg.Wait()
	fmt.Printf("Total lines processed: %d\n", total_line)
	total_line = 0 
    fmt.Println("Grep command executed on all log files.")
}

func runGrepOnSpecificLogFile(grepCommand string) {
	wg := sync.WaitGroup{}
	conn, err := net.Dial("tcp", "host.docker.internal:8081")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = conn.Write([]byte(grepCommand))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data sent to server successfully")
	
	buffer := make([]byte, 4096)
	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from connection:", err)
		return
	}
	wg.Add(1)
	SaveResultsToFile(string(buffer), grepCommand, &wg)

	defer conn.Close()

	fmt.Println("Data saved from server successfully")
}

func SaveResultsToFile(results string, grepCommand string, wg *sync.WaitGroup) {
	// wg.Done()
	commandParts, err := shlex.Split(grepCommand)
	if err != nil {
		fmt.Println("Error parsing command:", err)
		return
	}
	commandParts  = strings.Split(commandParts[len(commandParts)-1], ".")
	outputfile := "./result_files/" + commandParts[0] + "." + commandParts[1] + "." + "output.txt"
	file, err := os.Create(outputfile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	fmt.Println("Results saved to file:", outputfile)

	_, err = file.WriteString(results)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func Run() {
    fmt.Println("What would you like to do? (type 'exit' to quit)")
    fmt.Println("1. Run grep on all log files")
    fmt.Println("2. Run grep on a specific log file")
    fmt.Println("3. Exit")
	inputReader := bufio.NewReader(os.Stdin)
	userInput, err := inputReader.ReadString('\n')
	if err != nil {
		if err.Error() == "EOF" {
			fmt.Println("Input stream closed unexpectedly. Please try again.")
			return
		}
		fmt.Println("Error reading input:", err)
		return
	}
    userInput = userInput[:len(userInput)-1] // Remove the newline character
    if userInput == "" {
        fmt.Println("No input provided, please try again.")
    }

    switch userInput {
    case "1":
        fmt.Println("Running grep on all log files...")
		fmt.Println("Enter grep command (e.g., 'grep -i \"error\"'):")
        var grepCommand string
		inputReader := bufio.NewReader(os.Stdin)
		grepCommand, err = inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading grep command:", err)
			return
		}
		grepCommand = grepCommand[:len(grepCommand)-1] // Remove the newline character
        fmt.Printf("Running command: %s\n", grepCommand)
		runGrepOnAllLogFiles(grepCommand)
		Run()
        // Call the function to run grep on all log files
    case "2":
        fmt.Println("Enter grep command (e.g., 'grep -i \"error\" /path/to/file.log'):")
		inputReader := bufio.NewReader(os.Stdin)
		var grepCommand string
		grepCommand, err = inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading grep command:", err)
			Run()
		}
		if grepCommand == "" {
			fmt.Println("No grep command provided, please try again.")
			Run()
		}
		grepCommand = grepCommand[:len(grepCommand)-1] // Remove the newline character
        fmt.Printf("Running command: %s\n", grepCommand)
		runGrepOnSpecificLogFile(grepCommand)
	case "3":
        fmt.Println("Exiting...")
        return
    default:
        fmt.Println("Invalid option, please try again.")
    }
}
