package main

import (
	"fmt"
)

// produceNumbers sends natural numbers to the channel
// Input:
//  - naturals is the channel to send numbers
//  - digitsCount is the amount of number to send
func produceNumbers(naturals chan<- int, digitsCount int) {
	for x := 1; x <= digitsCount; x++ {
		naturals <- x
	}

	// close naturals channel and terminate the produceSquares gorutine
	close(naturals)
}

// produceSquares reads natural numbers from the naturals channel 
// and sends squares of them to the squares channel
// Input:
//  - naturals is the channel for reading natural numbers
//  - squares is the channel to send squares of the numbers
func produceSquares(naturals <-chan int, squares chan<- int) {
	for x := range naturals {
		squares <- x * x
	}

	// close squares channel
	close(squares)
}

func main() {
	// make channels to organize interaction between gorutines
	naturals := make(chan int)
	squares := make(chan int)

	// produce natural numbers
	go produceNumbers(naturals, 10)

	// consume natural numbers and produce squares
	go produceSquares(naturals, squares)

	// consume squares and print them to the terminal	
	for x := range squares {		
		fmt.Println(x)
	}
}