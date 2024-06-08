package jqueue

import "fmt"

// Create new ring queue.
// accepts cap and Resize as params
// cap is an int, the Size of the queue
// Resize is a flag, if true the queue will
// Resize when enqueing while full.
// if false, the queue will overwrite elements
// in a circular fashion.
func NewRingQueue(cap int, allowResize bool) *ringQueue {
	return &ringQueue{
		Data: make([]interface{}, cap),
		Start: 0,
		End: 0,
		Size: cap,
		IsFull: false,
		Resize: allowResize,
	}
}

// Utility function to get the minimum of two integers
// used for dequeing multiple elements to ensure
// more elements than are present aren't popped
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

// Resize the queue takes 1 argument, m
// which is the multiple to increase the Size by
func (q *ringQueue) ResizeQueue(factor float64) {
	if factor < 1 {
		q.ShrinkQueue(factor)
		return
	}

	newSize := int(float64(q.Size) * factor)
	if newSize - q.Size < 2 {
		newSize = q.Size + 2 // Ensure the new Size is at least 2 greater than the old Size
	}
	fmt.Printf("Resizing queue to %d\n", newSize)
	newData := make([]interface{}, newSize)

	if q.Start < q.End {
		copy(newData, q.Data[q.Start:q.End])
	} else {
		n := copy(newData, q.Data[q.Start:])
		copy(newData[n:], q.Data[:q.End])
	}

	var numElements int
	if q.Start <= q.End {
		numElements = q.End - q.Start
	} else {
		numElements = q.Size - q.Start + q.End
	}

	q.Data = newData
	q.Start = 0
	q.End = numElements % newSize
	q.Size = newSize
	q.IsFull = (numElements == newSize)
}

// Shrink the queue by a factor of 'factor'.
// If factor is >= 1, the queue is not shrunk.
// If factor is < 1, the queue is shrunk by that factor.
// For example, if factor is 0.75, the queue will be shrunk to 75% of its current Size.
func (q *ringQueue) ShrinkQueue(factor float64) {
	if factor >= 1 {
		return // Factor should be less than 1 to shrink the queue
	}

	newSize := int(float64(q.Size) * factor)
	if newSize < 1 {
		newSize = 1 // Ensure the new Size is at least 1
	}

	numElements := q.Size
	if q.Start < q.End {
		numElements = q.End - q.Start
	} else if q.Start > q.End {
		numElements = q.Size - q.Start + q.End
	}

	if newSize < numElements {
		newSize = numElements
	}

	newData := make([]interface{}, newSize)

	if q.Start < q.End {
		copy(newData, q.Data[q.Start:q.End])
	} else {
		n := copy(newData, q.Data[q.Start:])
		copy(newData[n:], q.Data[:q.End])
	}

	q.Data = newData
	q.Start = 0
	q.End = numElements
	q.Size = newSize
	q.IsFull = false
}
