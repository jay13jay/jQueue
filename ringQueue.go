package jqueue

import "fmt"

type ringQueue struct {
	data []interface{}
	start int
	end int
	size int
	isFull bool
	resize bool
}

func (q *ringQueue) Enqueue(v interface{}) {
    if q.isFull {
        if q.resize {
            q.ResizeQueue(1.25)
        } else {
            // Overwrite the oldest element
            q.start = (q.start + 1) % q.size
        }
    }

    q.data[q.end] = v
    q.end = (q.end + 1) % q.size

    if q.end == q.start {
        q.isFull = true
    } else {
        q.isFull = false
    }
}


// Pops 1 item from the queue, returns the value
// and a bool indicating if the queue is empty.
// true indicates an empty queue
func (q *ringQueue) Dequeue() (interface{}, bool) {
    if q.start == q.end && !q.isFull {
        return nil, true // Queue is empty
    }
    value := q.data[q.start]
    q.start = (q.start + 1) % q.size
    q.isFull = false
    return value, false
}

// Accepts an int argument, pops that amount from the queue.
// Returns a list with popped elements and a bool indicating
// if the queue is empty. True indicates empty queue.
// If requested amount is > items in queue, returns all items
func (q *ringQueue) DequeueAmount(amount int) ([]interface{}, bool) {
    var dequeued []interface{}

    // Determine the actual number of elements to dequeue
    numElements := amount
    if q.isFull {
        numElements = min(amount, q.size)
    } else if q.start <= q.end {
        numElements = min(amount, q.end-q.start)
    } else {
        numElements = min(amount, q.size-(q.start-q.end))
    }

    for i := 0; i < numElements; i++ {
        dequeued = append(dequeued, q.data[q.start])
        q.start = (q.start + 1) % q.size
        q.isFull = false
    }
		if q.start == q.end && !q.isFull {
        return dequeued, true // Queue is empty
    }

    return dequeued, false
}

// Print all elements in the queue.
// Does not pop any elements
func (q *ringQueue) Print() {
    fmt.Print("Queue elements: ")
    if q.isFull {
        fmt.Print(q.data[q.start])
        for i := (q.start + 1) % q.size; i != q.end; i = (i + 1) % q.size {
            fmt.Printf(" %v", q.data[i])
        }
    } else {
        for i := q.start; i != q.end; i = (i + 1) % q.size {
            fmt.Printf(" %v", q.data[i])
        }
    }
    fmt.Println()
}