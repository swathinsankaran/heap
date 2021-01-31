package main

type HeapType byte

const (
	MIN HeapType = iota
	MAX
)

type Heap interface {
	heapify(int)
	Build(*[]int)
	Print() []int
	Top() int
}

func newHeap(heapType HeapType) Heap {
	switch heapType {
	case MIN:
		return &minHeap{}
	case MAX:
		return &maxHeap{}
	}
	return nil
}
