package main

type maxHeap struct {
	data []int
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
	h.data = make([]int, len(*arr))
	copy(h.data, *arr)
	for i := len(h.data)/2 - 1; i >= 0; i-- {
		h.heapify(i)
	}
}

func (h *maxHeap) Print() []int {
	return h.data
}

func (h *maxHeap) Top() int {
	return h.data[0]
}

func (h *maxHeap) getMaxIndex(arr []int, a, b int) int {
	if a < len(arr) {
		if arr[a] > arr[b] {
			return a
		}
	}
	return b
}
