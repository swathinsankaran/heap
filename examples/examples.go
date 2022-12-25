package main

import (
	"fmt"

	heap "github.com/swathinsankaran/heap"
)

func main() {

	fmt.Println("Max heap - int8")
	fmt.Println("===============")

	int8Arr := []int8{5, 3, 6, 11, 2, 4, 10, 22, 34, 7}

	h := heap.NewHeap[int8](heap.Max, len(int8Arr))
	h.Build(&int8Arr)

	fmt.Println("Before heapify", int8Arr)
	fmt.Println("After heapify", h.Print())
	fmt.Println("Max: ", h.Top())
	fmt.Println("Popped: ", h.Pop(), h.Pop(), h.Pop(), h.Pop(), h.Pop(), h.Pop(), h.Pop(), h.Pop(), h.Pop(), h.Pop(), h.Pop(), h.Pop())
	fmt.Println("After pop", h.Print())
	h.Build(&int8Arr)
	fmt.Println("After new", h.Print())
	h.Insert(100)
	fmt.Println("After insert", h.Print())

	fmt.Println("Min heap - int16")
	fmt.Println("===============")

	int16Arr := []int16{5, 3, 6, 11, 2, 4, 10, 22, 34, 7}

	h1 := heap.NewHeap[int16](heap.Min, len(int16Arr))
	h1.Build(&int16Arr)

	fmt.Println("Before heapify", int16Arr)
	fmt.Println("After heapify", h1.Print())
	fmt.Println("Min: ", h1.Top())
	fmt.Println("Popped: ", h1.Pop(), h1.Pop(), h1.Pop(), h1.Pop(), h1.Pop(), h1.Pop(), h1.Pop(), h1.Pop(), h1.Pop(), h1.Pop(), h1.Pop(), h1.Pop())
	fmt.Println("After pop", h1.Print())

	h1.Build(&int16Arr)
	fmt.Println("After new", h1.Print())

	h1.Insert(1)
	fmt.Println("After insert", h1.Print())

	fmt.Println("Min heap - float32")
	fmt.Println("===============")

	float32Arr := []float32{5.1, 3.2, 6.4, 11.4, 2.56, 4.78, 10.66, 22.33, 34.78, 7.98, 2.32}

	h2 := heap.NewHeap[float32](heap.Min, len(float32Arr))
	h2.Build(&float32Arr)

	fmt.Println("Before heapify", float32Arr)
	fmt.Println("After heapify", h2.Print())
	fmt.Println("Min: ", h2.Top())
	fmt.Println("Popped: ", h2.Pop(), h2.Pop(), h2.Pop(), h2.Pop(), h2.Pop(), h2.Pop(), h2.Pop(), h2.Pop(), h2.Pop(), h2.Pop(), h2.Pop(), h2.Pop())
	fmt.Println("After pop", h2.Print())

	h2.Build(&float32Arr)
	fmt.Println("After new", h2.Print())

	h2.Insert(1)
	fmt.Println("After insert", h2.Print())
}
