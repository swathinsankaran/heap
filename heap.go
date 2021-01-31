package heap

type HeapType byte

const (
	Min HeapType = iota
	Max
)

type Heap interface {
	heapify(int)
	Build(*[]int)
	Print() []int
	Top() int
	Pop() int
	Insert(int)
}

func NewHeap(heapType HeapType, capacity int) Heap {
	switch heapType {
	case Min:
		return &minHeap{capacity: capacity, data: make([]int, capacity)}
	case Max:
		return &maxHeap{capacity: capacity, data: make([]int, capacity)}
	}
	return nil
}
