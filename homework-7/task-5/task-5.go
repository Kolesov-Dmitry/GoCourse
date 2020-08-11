package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

const (
	carsCount = 5
	distance = 100
)

// startCar starts a cat for the race
// Input:
//  - wg is required to notify what the car reached the finish line
//  - finishLine is the channel to send the final time
//  - carNumber is the car number
func startCar(wg *sync.WaitGroup, finishLine chan<- string, carNumber int) {
	defer wg.Done()

	// do some calculation befor the race
	readyTime := rand.Intn(10)
	speed := rand.Intn(5) + 5
	totalTime := int(distance / speed) + readyTime

	// imitation of participation in a race
	time.Sleep(time.Duration(totalTime) * time.Second)

	// send the final time to the channel
	finishLine <- fmt.Sprintf("%d - 00:%d", carNumber, totalTime)
}

func main() {
	// prepare WaitGroup and finishLine channel
	var wg sync.WaitGroup
	finishLine := make(chan string, carsCount)
	rand.Seed(int64(distance))

	// start cars for the race
	wg.Add(carsCount)
	for x := 0; x < carsCount; x++ {
		go startCar(&wg, finishLine, x + 1)
	}

	// waih until the race is finished
	wg.Wait()
	close(finishLine)

	// print the race results
	for car := range finishLine {
		fmt.Println(car)
	}
}
