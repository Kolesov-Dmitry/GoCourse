package main

// The first idea was to implement Filter interface similar to Sort interface
// which is present in Go.
// But it takes me the whole eternity to realize that Go doesn't work like that and I can't
// implement it in this way... :'(
// And now I know why there is no Filter interface in Go.

// So ...

import (
	"fmt"
)

// Counter interface
type Counter interface {
	Len() int
	Predicate(i int) bool
}

// Count the amount of items satisfying the condition in the Predicate
// Input:
//  - object which implements Counter interface
// Output:
//  - amount of items
func Count(c Counter) int {
	result := 0
	for idx := 0; idx < c.Len(); idx++ {
		if c.Predicate(idx) {
			result++
		}
	}

	return result
}

// Odds wrapper to count the amount of odd numbers
type Odds []int

// Len implementation for Odds
func (a Odds) Len() int {
	return len(a)
}

// Predicate implementation for Odds
func (a Odds) Predicate(i int) bool {
	if i < 0 || i >= a.Len() {
		return false
	}

	return a[i]%2 != 0
}

// Evens wrapper to count the amount of even numbers
type Evens []int

// Len implementation for Evens
func (a Evens) Len() int {
	return len(a)
}

// Predicate implementation for Evens
func (a Evens) Predicate(i int) bool {
	if i < 0 || i >= a.Len() {
		return false
	}

	return a[i]%2 == 0
}

func main() {

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	odds := Count(Odds(a))
	evens := Count(Evens(a))

	fmt.Println("Odds: ", odds)
	fmt.Println("Evens:", evens)
}
