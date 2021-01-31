package heap

import (
	"reflect"
	"testing"
)

func TestNewMaxHeap(t *testing.T) {
	h := NewHeap(Max, 8)
	if reflect.TypeOf(h).String() != "*heap.maxHeap" {
		t.Errorf("New heap failed, got: %v, want: *heap.maxHeap.", reflect.TypeOf(h).String())
	}
}

func TestNewMinHeap(t *testing.T) {
	h := NewHeap(Min, 8)
	if reflect.TypeOf(h).String() != "*heap.minHeap" {
		t.Errorf("New heap failed, got: %v, want: *heap.minHeap.", reflect.TypeOf(h).String())
	}
}

func TestNewNilHeap(t *testing.T) {
	h := NewHeap(HeapType(2), 8)
	if h != nil {
		t.Errorf("New heap failed, got: %v, want: <nil>.", h)
	}
}
