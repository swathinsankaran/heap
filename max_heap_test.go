package heap

import (
	"reflect"
	"testing"
)

func TestMaxHeapBuild(t *testing.T) {
	arr := []int{1, 12, 3, 14, 5, 16, 7, 8}
	h := NewHeap[int](Max, 8)
	h.Build(&arr)
	expectedArray := []int{16, 14, 7, 12, 5, 3, 1, 8}
	gotArray := h.Print()
	if !reflect.DeepEqual(expectedArray, gotArray) {
		t.Errorf("Build heap failed, got: arr=%v, want: arr=%v.", gotArray, expectedArray)
	}
}

func TestMaxHeapTop(t *testing.T) {
	arr := []int{1, 12, 3, 14, 5, 16, 7, 8}
	h := NewHeap[int](Max, 8)
	h.Build(&arr)
	expectedValue := 16
	gotValue := h.Top()
	if expectedValue != gotValue {
		t.Errorf("Build heap failed, got: value=%d, want: value=%d.", gotValue, expectedValue)
	}
}

func TestMaxHeapPop(t *testing.T) {
	arr := []int{1, 12, 3, 14, 5, 16, 7, 8}
	h := NewHeap[int](Max, 8)
	h.Build(&arr)
	expectedValue := 16
	gotValue := h.Pop()
	if expectedValue != gotValue {
		t.Errorf("Build heap failed, got: value=%d, want: value=%d.", gotValue, expectedValue)
	}
	expectedArray := []int{14, 12, 7, 8, 5, 3, 1}
	gotArray := h.Print()
	if !reflect.DeepEqual(expectedArray, gotArray) {
		t.Errorf("Build heap failed, got: arr=%v, want: arr=%v.", gotArray, expectedArray)
	}
}

func TestMaxHeapPopEmpty(t *testing.T) {
	arr := []int{1}
	h := NewHeap[int](Max, 1)
	h.Build(&arr)
	expectedValue := 1
	gotValue := h.Pop()
	if expectedValue != gotValue {
		t.Errorf("Build heap failed, got: value=%d, want: value=%d.", gotValue, expectedValue)
	}
	expectedValue = 0
	gotValue = h.Pop()
	if expectedValue != gotValue {
		t.Errorf("Build heap failed, got: value=%d, want: value=%d.", gotValue, expectedValue)
	}
}

func TestMaxHeapInsert(t *testing.T) {
	arr := []int{1}
	h := NewHeap[int](Max, 1)
	h.Build(&arr)
	expectedValue := 1
	gotValue := h.Pop()
	if expectedValue != gotValue {
		t.Errorf("Build heap failed, got: value=%d, want: value=%d.", gotValue, expectedValue)
	}
	h.Insert(10)
	expectedValue = 10
	gotValue = h.Pop()
	if expectedValue != gotValue {
		t.Errorf("Build heap failed, got: value=%d, want: value=%d.", gotValue, expectedValue)
	}
}
