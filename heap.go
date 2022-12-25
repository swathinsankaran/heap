package heap

type HeapType byte

const (
	Min HeapType = iota
	Max
)

type SupportedTypes interface {
	int | int8 | int16 | int32 | int64 | float64 | float32
}

type Heap[T SupportedTypes] interface {
	heapify(int)
	Build(*[]T)
	Print() []T
	Top() T
	Pop() T
	Insert(T)
}

func NewHeap[T SupportedTypes](heapType HeapType, capacity int) Heap[T] {
	switch heapType {
	case Min:
		return &minHeap[T]{capacity: capacity, data: make([]T, capacity)}
	case Max:
		return &maxHeap[T]{capacity: capacity, data: make([]T, capacity)}
	}
	return nil
}
