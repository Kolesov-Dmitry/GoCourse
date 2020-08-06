package main

import (
	"fmt"
	"log"
	"net/http"
)

// mirroredQuery makes GET request to three different hosts
// Output:
//  - name of the host which answered first
func mirroredQuery() string {
	responses := make(chan string, 3)

	// make three GET requests
	go func() {
		responses <- request("https://www.google.com")
	}()

	go func() {
		responses <- request("https://www.hyperskill.org")
	}()

	go func() {
		responses <- request("https://www.codewars.com")
	}()

	return <-responses
}

// request makes GET request to the host
// Input:
//  - host name to make a request
// Output:
//  - host name it self
func request(hostname string) string {
	resp, err := http.Get(hostname)
	if err != nil {
		log.Fatal(err)
	}

	resp.Body.Close()

	return hostname
}

func main() {
	host := mirroredQuery()
	fmt.Println(host, "is the quickest one")
}
