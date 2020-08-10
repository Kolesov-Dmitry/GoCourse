package main

import (
	"fmt"
	"bufio"
	"flag"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// start TCP client and connect to server
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// start goroutine to receive messages from the server
	go receiver(conn)

	// Our first message is the nickname
	nickName := getNickName()
	fmt.Fprintf(conn, nickName + "\n")

	// read input and send message to the server
	input := bufio.NewScanner(os.Stdin)
	for {	
		msg := readUserInput(input)
		if len(msg) != 0 {
			fmt.Fprintf(conn, msg + "\n")
		} 
	}
}

// receiver reads TCP socket and prints messages to the terminal
// Input:
//  - conn in the TCP connection
func receiver(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		
		printMessage(msg)
	}
}

// readUserInput reads Stdin for user messages
// Input:
//  - input is a Scanner to handle user input
// Output:
//  - user message
func readUserInput(input *bufio.Scanner) string {
	fmt.Print("> ")
	input.Scan()
	fmt.Print("\033[A\033[K")	// erase line in the terminal where user typed his message

	return strings.TrimSpace(input.Text())
}

// printMessage outputs message to the terminal
// Input:
//  - msg is the message to output
func printMessage(msg string) {
	trimmedMsg := strings.TrimSpace(msg)
	if len(trimmedMsg) != 0 {
		fmt.Println("\r" + trimmedMsg)
		fmt.Print("> ")
	}
}


// getNickName reads user's nickname from the console flags
// Output:
//   - user's nickname
func getNickName() string {
	var nickName string
	flag.StringVar(&nickName, "n", "Anonymous", "user nickname")
	flag.Parse()

	return nickName
}