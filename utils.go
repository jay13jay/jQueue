package jqueue

// Create new ring queue.
// accepts cap and resize as params
// cap is an int, the size of the queue
// resize is a flag, if true the queue will
// resize when enqueing while full.
// if false, the queue will overwrite elements
// in a circular fashion.
func NewRingQueue(cap int, allowResize bool) *ringQueue {
	return &ringQueue{
		data: make([]interface{}, cap),
		start: 0,
		end: 0,
		size: cap,
		isFull: false,
		resize: allowResize,
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

// resize the queue takes 1 argument, m
// which is the multiple to increase the size by
func (q *ringQueue) ResizeQueue(factor float64) {
    newSize := int(float64(q.size) * factor)
    newData := make([]interface{}, newSize)

		// Check if size is smaller than 1, if so, shrink the queue
		if factor < 1 {
			q.ShrinkQueue(factor)
		}

    if q.start < q.end {
        copy(newData, q.data[q.start:q.end])
    } else {
        n := copy(newData, q.data[q.start:])
        copy(newData[n:], q.data[:q.end])
    }

    q.start = 0
    q.end = q.size
    q.size = newSize
    q.data = newData
    q.isFull = false
}

// Shrink the queue by a factor of 'factor'.
// If factor is >= 1, the queue is not shrunk.
// If factor is < 1, the queue is shrunk by that factor.
// For example, if factor is 0.75, the queue will be shrunk to 75% of its current size.
func (q *ringQueue) ShrinkQueue(factor float64) {
    if factor >= 1 {
        return // Factor should be less than 1 to shrink the queue
    }

    newSize := int(float64(q.size) * factor)
    if newSize < 1 {
        newSize = 1 // Ensure the new size is at least 1
    }

    // Only shrink if the new size is less than the number of elements in the queue
    numElements := q.size
    if q.start < q.end {
        numElements = q.end - q.start
    } else if q.start > q.end {
        numElements = q.size - q.start + q.end
    }

    if newSize < numElements {
        newSize = numElements
    }

    newData := make([]interface{}, newSize)

    if q.start < q.end {
        copy(newData, q.data[q.start:q.end])
    } else {
        n := copy(newData, q.data[q.start:])
        copy(newData[n:], q.data[:q.end])
    }

    q.start = 0
    q.end = numElements
    q.size = newSize
    q.data = newData
}
