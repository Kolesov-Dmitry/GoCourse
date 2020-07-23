package queue

import "errors"

// Queue implemets First In First Out data structure
type Queue struct {
	
	// Let's use slice instead of pointers to the Head and Current element of the queue
	// and let Go do all the "memory things" on its own
	
	// data contains the elements of the queue
	data []interface{}
}

// NewQueue creates empty Queue
// Qutput:
//  - pointer to the created instance of Queue
func NewQueue() *Queue {
	// Preallocate queue of 10 elements
	return &Queue {
		data: make([]interface{}, 0, 10),
	}
}

// Size returns the total amount of items in the queue
// Output:
//  - size of the queue
func (q *Queue) Size() int {
	return len(q.data)
}

// Push appends item on the front of the queue
// Input:
//  - item is an item to be appended to the queue
func (q *Queue) Push(item interface{}) {
	q.data = append(q.data, item)
}

// Pop exstracts the item from the back of the queue
// Output:
//  - the item, in case of success
//  - error if the queue is empty
func (q *Queue) Pop() (interface{}, error) {	
	if len(q.data) == 0 {
		return nil, errors.New("The queue is empty")
	}

	// exstract item and remove it from the queue
	item := q.data[0]
	q.data = q.data[1:]

	return item, nil
}