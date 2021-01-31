package heap

import (
	"reflect"
	"testing"
)

func TestMinHeapBuild(t *testing.T) {
	arr := []int{1, 12, 3, 14, 5, 16, 7, 8}
	h := NewHeap(Min, 8)
	h.Build(&arr)
	expectedArray := []int{1, 5, 3, 8, 12, 16, 7, 14}
	gotArray := h.Print()
	if !reflect.DeepEqual(expectedArray, gotArray) {
		t.Errorf("Build heap failed, got: arr=%v, want: arr=%v.", gotArray, expectedArray)
	}
}

func TestMinHeapTop(t *testing.T) {
	arr := []int{1, 12, 3, 14, 5, 16, 7, 8}
	h := NewHeap(Min, 8)
	h.Build(&arr)
	expectedValue := 1
	gotValue := h.Top()
	if expectedValue != gotValue {
		t.Errorf("Build heap failed, got: value=%d, want: value=%d.", gotValue, expectedValue)
	}
}

func TestMinHeapPop(t *testing.T) {
	arr := []int{1, 12, 3, 14, 5, 16, 7, 8}
	h := NewHeap(Min, 8)
	h.Build(&arr)
	expectedValue := 1
	gotValue := h.Pop()
	if expectedValue != gotValue {
		t.Errorf("Build heap failed, got: value=%d, want: value=%d.", gotValue, expectedValue)
	}
	expectedArray := []int{3, 5, 7, 8, 12, 16, 14, 0}
	gotArray := h.Print()
	if !reflect.DeepEqual(expectedArray, gotArray) {
		t.Errorf("Build heap failed, got: arr=%v, want: arr=%v.", gotArray, expectedArray)
	}
}

func TestMinHeapPopEmpty(t *testing.T) {
	arr := []int{1}
	h := NewHeap(Min, 1)
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
