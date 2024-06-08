package jqueue

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRingQueuePerformance(t *testing.T) {
    initialSize := 1
    iterations := 10

    rq := NewRingQueue(initialSize, true)

    start := time.Now()
    for i := 0; i < iterations; i++ {
        rq.Enqueue(rand.Int())
    }
    duration := time.Since(start)
    fmt.Printf("Enqueue %d elements took %v\n", iterations, duration)

    start = time.Now()
    for i := 0; i < iterations; i++ {
        rq.Dequeue()
    }
    duration = time.Since(start)
    fmt.Printf("Dequeue %d elements took %v\n", iterations, duration)

    rq = NewRingQueue(initialSize, true)
    start = time.Now()
    for i := 0; i < iterations; i++ {
        rq.ResizeQueue(1.5)
    }
    duration = time.Since(start)
    fmt.Printf("Resize queue %d times took %v\n", iterations, duration)

    for i := 0; i < initialSize; i++ {
        rq.Enqueue(rand.Int())
    }
    start = time.Now()
    for i := 0; i < iterations; i++ {
        rq.ShrinkQueue(0.5)
    }
    duration = time.Since(start)
    fmt.Printf("Shrink queue %d times took %v\n", iterations, duration)
}

func BenchmarkRingQueue(b *testing.B) {
    initialSize := 10

    b.Run("Enqueue", func(b *testing.B) {
        rq := NewRingQueue(initialSize, true)
        for i := 0; i < b.N; i++ {
            rq.Enqueue(rand.Int())
        }
    })

    b.Run("Dequeue", func(b *testing.B) {
        rq := NewRingQueue(initialSize, true)
        for i := 0; i < initialSize; i++ {
            rq.Enqueue(rand.Int())
        }
        for i := 0; i < b.N; i++ {
            rq.Dequeue()
        }
    })

    b.Run("ResizeQueue", func(b *testing.B) {
        rq := NewRingQueue(initialSize, true)
        for i := 0; i < b.N; i++ {
            rq.ResizeQueue(1.25)
        }
    })

    b.Run("ShrinkQueue", func(b *testing.B) {
        rq := NewRingQueue(initialSize, true)
        for i := 0; i < initialSize; i++ {
            rq.Enqueue(rand.Int())
        }
        for i := 0; i < b.N; i++ {
            rq.ShrinkQueue(0.5)
        }
    })
}
