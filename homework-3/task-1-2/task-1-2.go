package main

import (
	"fmt"
)

// Vehicle struct
type Vehicle struct {
	Brand     string
	IssueYear int

	// In litres
	TrunkVolume int

	EngineIsStarted  bool
	WindowsAreOpened bool

	// 0.0 is empy, 1.0 is completly full
	TrunkFullness float32
}

// I'm going to cheet a bit
func (v Vehicle) String() string {
	isStarted := "started"
	if v.EngineIsStarted == false {
		isStarted = "shutdown"
	}

	isOpened := "opened"
	if v.WindowsAreOpened == false {
		isOpened = "closed"
	}

	fullness := fmt.Sprintf("%.2f", v.TrunkFullness)
	if v.TrunkFullness < 0.0001 {
		fullness = "empty"
	} else if v.TrunkFullness > 0.9999 {
		fullness = "full"
	}

	return fmt.Sprintf("Brand: %s\nIssueYear: %d\nTrunkVolume: %d litres\nEngine: %s\nWindows: %s\nTrunkFullness: %s",
		v.Brand,
		v.IssueYear,
		v.TrunkVolume,
		isStarted,
		isOpened,
		fullness)
}

func main() {
	lambo := Vehicle{
		Brand:            "Lamborghini Huracan",
		IssueYear:        2019,
		TrunkVolume:      100,
		EngineIsStarted:  false,
		WindowsAreOpened: true,
		TrunkFullness:    0.0,
	}

	kamaz := Vehicle{
		Brand:            "KAMAZ",
		IssueYear:        2015,
		TrunkVolume:      10_000,
		EngineIsStarted:  false,
		WindowsAreOpened: false,
		TrunkFullness:    1.0,
	}

	fmt.Println(lambo, "\n")
	fmt.Println(kamaz, "\n")

	lambo.EngineIsStarted = true
	lambo.TrunkFullness = 0.5

	kamaz.TrunkFullness = 0.0
	kamaz.WindowsAreOpened = true

	fmt.Println("--------------------------\n")
	fmt.Println(lambo, "\n")
	fmt.Println(kamaz, "\n")
}
