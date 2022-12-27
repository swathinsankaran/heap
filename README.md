# heap

[![GoDoc](https://godoc.org/github.com/swathinsankaran/heap?status.svg)](https://godoc.org/github.com/swathinsankaran/heap)
[![Go Report Card](https://goreportcard.com/badge/github.com/swathinsankaran/heap)](https://goreportcard.com/report/github.com/swathinsankaran/heap)

> A heap data structure library which utilizes `Generics`. This library currently supports heaps of type int, int8, int16, int32, int64, float32 and float64.

## Install

```
# using go get
$ go get -u github.com/swathinsankaran/heap
```

## Usage

```go
package main

import (
	"fmt"

	heap "github.com/swathinsankaran/heap"
)

func main() {

	int8Arr := []int8{5, 3, 6, 11, 2, 4, 10, 22, 34, 7}

	h := heap.NewHeap[int8](heap.Max, len(int8Arr))
	h.Build(&int8Arr)             // Builds the heap with the array passed.
	fmt.Println(h.Print())        // Prints the values in the heap after heapify.
	fmt.Println("Max: ", h.Top()) // Gets the max value.
	fmt.Println(h.Pop())          // Pops the max value from the heap.
	h.Insert(100)                 // Inserts 100 into the heap.
```

## License

MIT © [Swathin Sankaran](https://github.com/swathinsankaran)
