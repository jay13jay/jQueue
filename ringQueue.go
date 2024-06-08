package jqueue

type ringQueue struct {
	data []interface{}
	start int
	end int
	size int
	isFull bool
	resize bool
}

func myFunc() {
	rq := NewRingQueue(10, true)

	rq.ResizeQueue(2)
}