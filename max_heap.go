package heap

import "fmt"

type maxHeap struct {
	data     []int
	capacity int
}

func (h *maxHeap) heapify(i int) {
	l := (2 * i) + 1
	r := (2 * i) + 2
	index := i
	index = h.getMaxIndex(h.data, l, index)
	index = h.getMaxIndex(h.data, r, index)
	if i != index {
		swap(&h.data[index], &h.data[i])
		h.heapify(index)
	}
}

func (h *maxHeap) Build(arr *[]int) {
	if len(h.data) == 0 {
		h.data = make([]int, len(*arr))
	}
	copy(h.data, *arr)
	h.capacity = len(*arr)
	for i := h.capacity/2 - 1; i >= 0; i-- {
		h.heapify(i)
	}
}

func (h *maxHeap) Print() []int {
	return h.data
}

func (h *maxHeap) Top() int {
	return h.data[0]
}

func (h *maxHeap) Pop() int {
	if h.capacity == 0 {
		fmt.Println("heap is empty")
		return 0
	}
	v := h.data[0]
	h.data[0] = h.data[h.capacity-1]
	h.data = h.data[:h.capacity-1]
	h.capacity--
	h.heapify(0)
	return v
}

func (h *maxHeap) Insert(value int) {
	h.data = append(h.data, value)
	h.capacity++
	for i := h.capacity - 1; i > 0 && h.data[i/2] < h.data[i]; i /= 2 {
		swap(&h.data[i/2], &h.data[i])
	}
}

func (h *maxHeap) getMaxIndex(arr []int, a, b int) int {
	if a < h.capacity {
		if arr[a] > arr[b] {
			return a
		}
	}
	return b
}
