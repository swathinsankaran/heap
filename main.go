package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 3, 6, 1, 2, 4, 10, 22, 34, 7}
	fmt.Println("Before heapify", arr)

	h := newHeap(MAX)
	h.Build(&arr)

	fmt.Println("After heapify", h.Print())
	fmt.Println("Max: ", h.Top())

	h1 := newHeap(MIN)
	h1.Build(&arr)

	fmt.Println("After heapify", h1.Print())
	fmt.Println("Min: ", h1.Top())

}
