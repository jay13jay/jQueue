# jqueue
golang queue implementation

# Usage

## Create a queue
```go
// create a resizable queue with
// starting capacity of 10
q = NewRingQueue(10, true)
```

```go
// create a fixed size queue with
// capacity of 20
q = NewRingQueue(20, false)
```

## Resize a queue
The queue will automatically resize when trying to enque if the queue is full. 
The resize function takes 1 argument, factor, which is the factor to expand the queue by.
If you attempt to Enqueue when the queue is full, the automatic resize will grow the queue by 1.25, or 25%

Ex:
```go
// doubles the queue size
q.ResizeQueue(2)

// grows queue by 25%
// note: this is the default behavior
q.ResizeQueue(1.25)

// resize before enqueue
if q.isFull {
  q.Resize(2)
  q.Enqueue(data)
}
```
*Note: Queue capacity is always an int value, so value of capacity will round down to nearest whole number*

## Get element(s)
```go
// pop a single element
element, isEmpty := q.Dequeue()

// pop 5 elements
elements, isEmpty := q.DequeueAmount(5)
```

## Example implementation
```go
func main() {
    rq := jqueue.NewRingQueue(5, true)

    rq.Enqueue(1)
    rq.Enqueue(2)
    rq.Enqueue(3)
    rq.Enqueue(4)
    rq.Enqueue(5)

    dequeued, empty := rq.DequeueAmount(3)
    fmt.Println("Dequeued:", dequeued, "Empty:", empty) // Output: Dequeued: [1 2 3] Empty: false

    rq.Enqueue(6)
    dequeued, isEmpty = rq.DequeueAmount(4)
    fmt.Println("Dequeued:", dequeued, "Empty:", empty) // Output: Dequeued: [4 5 6] Empty: true

    dequeued, isEmpty = rq.DequeueAmount(5)
    fmt.Println("Dequeued:", dequeued, "Empty:", empty) // Output: Dequeued: [] Empty: true

    rq.Print() // Output: Queue elements: Queue is empty
}
```
