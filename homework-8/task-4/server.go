package main
import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn)
	}
}

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {	
	ch := make(chan string)
	go clientWriter(conn, ch)
	
	var nickName string
	input := bufio.NewReader(conn)
	for {
		msg, err := input.ReadString('\n')
		if err != nil {			
			break
		}

		// The fist user message his nickname
		if len(nickName) == 0 {
			nickName = strings.TrimSpace(msg)
			messages <- nickName + " has arrived"
			entering <- ch
		} else {
			messages <- nickName + ": " + msg
		}
	}

	leaving <- ch
	messages <- nickName + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}