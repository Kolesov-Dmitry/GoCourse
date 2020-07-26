package main

// As you mentioned on the previous lesson we didn't cover methods yet
// But who really cares, let's do some FIFO stuff 

import (
	"fmt"
	"homework.com/v3/task-3/internal/app/queue"
)

func main() {
	// create new queue
	queue := queue.NewQueue()

	// insert some elements
	queue.Push(10)
	queue.Push(11)
	queue.Push(12)

	// print them out
	for queue.Size() > 0 {
		item, _ := queue.Pop();
		fmt.Println(item)
	}
	
	// catch the empty queue error
	_, err := queue.Pop();	
	fmt.Println(err)
}