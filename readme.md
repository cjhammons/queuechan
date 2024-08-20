# Qchan

A simple queue implementation in Go using channels.

## Installation

To install the package, use the following command:

```sh
go get github.com/cjhammons/qchan
```

## Usage

Here is an example of how to use the `qchan` package:

```go
package main

import (
    "fmt"
    "github.com/cjhammons/qchan"
)

func main() {
    q := qchan.NewQueue(10) // Create a new queue with a buffer size of 10
    
    // Enqueue items
    q.Enqueue(1)
    q.Enqueue(2)
    
    // Dequeue items
    value, ok := q.Dequeue()
    if ok {
        fmt.Println(value) // Output: 1
    }
    
    value, ok = q.Dequeue()
    if ok {
        fmt.Println(value) // Output: 2
    }
    
    // Check if the queue is empty
    if q.IsEmpty() {
        fmt.Println("Queue is empty")
    }
    
    // Enqueue and wait for dequeue
    go func() {
        q.Enqueue(3)
    }()
    
    value, ok = q.WaitAndDequeue()
    if ok {
        fmt.Println(value) // Output: 3
    }
}
```

## Methods

### `NewQueue(bufferSize int) *Queue`
Creates and returns a new `Queue` with a specified buffer size.

### `Enqueue(value any)`
Adds a new value to the end of the queue.

### `WaitAndDequeue() (any, bool)`
Waits for an item to be available and then dequeues it. Returns the dequeued value and a boolean indicating if the operation was successful.

### `Dequeue() (any, bool)`
Removes and returns the value from the front of the queue. Returns the dequeued value and a boolean indicating if the operation was successful.

### `IsEmpty() bool`
Checks if the queue is empty. Returns true if the queue is empty, otherwise false.

### `Size() int`
Returns the number of items in the queue.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.