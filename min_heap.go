package heap

import "sync"

type minHeap[T SupportedTypes] struct {
	data     []T
	capacity int
	lock     sync.RWMutex
}

func (h *minHeap[T]) heapify(i int) {
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

func (h *minHeap[T]) Build(arr *[]T) {
	h.lock.Lock()
	defer h.lock.Unlock()
	if len(h.data) == 0 {
		h.data = make([]T, len(*arr))
	}
	copy(h.data, *arr)
	h.capacity = len(*arr)
	for i := h.capacity/2 - 1; i >= 0; i-- {
		h.heapify(i)
	}
}

func (h *minHeap[T]) Print() []T {
	h.lock.RLock()
	defer h.lock.RUnlock()
	return h.data
}

func (h *minHeap[T]) Top() T {
	h.lock.RLock()
	defer h.lock.RUnlock()
	return h.data[0]
}

func (h *minHeap[T]) Pop() T {
	h.lock.Lock()
	defer h.lock.Unlock()
	var v T
	if h.capacity != 0 {
		v = h.data[0]
		h.data[0] = h.data[h.capacity-1]
		h.data = h.data[:h.capacity-1]
		h.capacity--
		h.heapify(0)
	}
	return v
}

func (h *minHeap[T]) Insert(value T) {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.data = append(h.data, value)
	h.capacity++
	for i := h.capacity - 1; i > 0 && h.data[i/2] > h.data[i]; i /= 2 {
		swap(&h.data[i/2], &h.data[i])
	}
}

func (h *minHeap[T]) getMinIndex(arr []T, a, b int) int {
	if a < h.capacity {
		if arr[a] < arr[b] {
			return a
		}
	}
	return b
}
