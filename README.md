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
The resize function takes 1 argument, m, which is the multiple to expand the queue by.

Ex:
```go
// doubles the queue size
q.ResizeQueue(2)

// grows queue by 25%
q.ResizeQueue(1.25)
```
Note: Queue capacity is always an int value, so value of resize will round down to nearest whole number



