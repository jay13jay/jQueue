package jqueue

import "fmt"

type ringQueue struct {
	Data []interface{}
	Start int
	End int
	Size int
	IsFull bool
	Resize bool
}

func (q *ringQueue) Enqueue(v interface{}) {
    if q.IsFull {
        if q.Resize {
            q.ResizeQueue(1.25)
        } else {
            // Overwrite the oldest element
            q.Start = (q.Start + 1) % q.Size
        }
    }

    q.Data[q.End] = v
    q.End = (q.End + 1) % q.Size

    if q.End == q.Start {
        q.IsFull = true
    } else {
        q.IsFull = false
    }
}


// Pops 1 item from the queue, returns the value
// and a bool indicating if the queue is empty.
// true indicates an empty queue
func (q *ringQueue) Dequeue() (interface{}, bool) {
    if q.Start == q.End && !q.IsFull {
        return nil, true // Queue is empty
    }
    value := q.Data[q.Start]
    q.Start = (q.Start + 1) % q.Size
    q.IsFull = false
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
    if q.IsFull {
        numElements = min(amount, q.Size)
    } else if q.Start <= q.End {
        numElements = min(amount, q.End-q.Start)
    } else {
        numElements = min(amount, q.Size-(q.Start-q.End))
    }

    for i := 0; i < numElements; i++ {
        dequeued = append(dequeued, q.Data[q.Start])
        q.Start = (q.Start + 1) % q.Size
        q.IsFull = false
    }
		if q.Start == q.End && !q.IsFull {
        return dequeued, true // Queue is empty
    }

    return dequeued, false
}

// Print all elements in the queue.
// Does not pop any elements
func (q *ringQueue) Print() {
    fmt.Print("Queue elements: ")
    if q.IsFull {
        fmt.Print(q.Data[q.Start])
        for i := (q.Start + 1) % q.Size; i != q.End; i = (i + 1) % q.Size {
            fmt.Printf(" %v", q.Data[i])
        }
    } else {
        for i := q.Start; i != q.End; i = (i + 1) % q.Size {
            fmt.Printf(" %v", q.Data[i])
        }
    }
    fmt.Println()
}