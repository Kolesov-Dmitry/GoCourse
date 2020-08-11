package main

import (
	"fmt"
	"time"
)

// spinner plays a spinner in a terminal
// Input:
//  - dalay is the speed of the spin
func spinner(delay time.Duration) {
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r)
			time.Sleep(delay)
		}
	}
}

func main() {
	// start spinner in a sepatate gorutine
	go spinner(100 * time.Millisecond)

	// waiting for 10 seconds
	time.Sleep(10 * time.Second)
}

