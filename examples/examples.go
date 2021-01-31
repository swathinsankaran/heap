package main

import (
	"fmt"

	heap "github.com/swathins079/heap"
)

func main() {
	arr := []int{5, 3, 6, 1, 2, 4, 10, 22, 34, 7}
	fmt.Println("Before heapify", arr)

	h := heap.NewHeap(heap.Max, 10)
	h.Build(&arr)

	fmt.Println("After heapify", h.Print())
	fmt.Println("Max: ", h.Top())
	fmt.Println("Popped: ", h.Pop(), h.Pop(), h.Pop(), h.Pop(), h.Pop(), h.Pop(), h.Pop(), h.Pop(), h.Pop(), h.Pop(), h.Pop(), h.Pop())
	fmt.Println("After pop", h.Print())
	h.Build(&arr)
	fmt.Println("After new", h.Print())

	h1 := heap.NewHeap(heap.Min, 10)
	h1.Build(&arr)

	fmt.Println("After heapify", h1.Print())
	fmt.Println("Min: ", h1.Top())
	fmt.Println("Popped: ", h1.Pop(), h1.Pop(), h1.Pop(), h1.Pop(), h1.Pop(), h1.Pop(), h1.Pop(), h1.Pop(), h1.Pop(), h1.Pop(), h1.Pop(), h1.Pop())
	fmt.Println("After pop", h1.Print())

	h1.Build(&arr)
	fmt.Println("After new", h1.Print())
}
