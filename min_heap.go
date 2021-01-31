package heap

import "fmt"

type minHeap struct {
	data     []int
	capacity int
}

func (h *minHeap) heapify(i int) {
	l := (2 * i) + 1
	r := (2 * i) + 2
	index := i
	index = h.getMinIndex(h.data, l, index)
	index = h.getMinIndex(h.data, r, index)
	if i != index {
		swap(&h.data[index], &h.data[i])
		h.heapify(index)
	}
}

func (h *minHeap) Build(arr *[]int) {
	copy(h.data, *arr)
	h.capacity = len(*arr)
	for i := h.capacity/2 - 1; i >= 0; i-- {
		h.heapify(i)
	}
}

func (h *minHeap) Print() []int {
	return h.data
}

func (h *minHeap) Top() int {
	return h.data[0]
}

func (h *minHeap) Pop() int {
	if h.capacity == 0 {
		fmt.Println("heap is empty")
		return 0
	}
	v := h.data[0]
	h.data[0] = h.data[h.capacity-1]
	h.data[h.capacity-1] = 0
	h.capacity--
	h.heapify(0)
	return v
}

func (h *minHeap) getMinIndex(arr []int, a, b int) int {
	if a < h.capacity {
		if arr[a] < arr[b] {
			return a
		}
	}
	return b
}
