package main

import (
	"io"
	"log"
	"net"
	"time"
)

const (
	exitCmd = "exit"
	exitCmdLen = len(exitCmd)
)

// sendTime sends current time throught TCP connection
// Input:
//  - conn is the TCP connection
func sendTime(conn net.Conn) {
	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n\r"))
		if err != nil {
			return
		}

		time.Sleep(1 * time.Second)
	}
}

// handleConn handles the incoming TCP connection
// Input:
//  - conn is the incoming TCP connection
// Output:
//  - true if the command to stop the listener has been received
//  - false if connection was closed
func handleConn(conn net.Conn) bool {
	defer conn.Close()

	// start gorutine which will infinitly send current time throught the connection
	go sendTime(conn)

	// read socket for incomming command
	buffer := make([]byte, exitCmdLen)
	for {
		if n, err := conn.Read(buffer); err == nil && n == exitCmdLen {
			if cmd := string(buffer[:n]); cmd == exitCmd {
				return true
			}
		} else if err != nil {
			return false
		}
	}

	return true
}

func main() {
	// start TCP listener
	listener, err := net.Listen("tcp", "10.4.8.223:8080")
	if err != nil {
		log.Fatal(err)
	}

	// waiting for incomming connection
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		
		if exit := handleConn(conn); exit {
			return
		}
	}
}
