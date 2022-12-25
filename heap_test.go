package heap

import (
	"reflect"
	"testing"
)

func TestNewMaxHeapInt(t *testing.T) {
	testcases := []struct {
		heap         interface{}
		expectedType string
	}{
		{
			heap:         NewHeap[int](Max, 1),
			expectedType: "*heap.maxHeap[int]",
		}, {
			heap:         NewHeap[int8](Max, 1),
			expectedType: "*heap.maxHeap[int8]",
		}, {
			heap:         NewHeap[int16](Max, 1),
			expectedType: "*heap.maxHeap[int16]",
		}, {
			heap:         NewHeap[int32](Max, 1),
			expectedType: "*heap.maxHeap[int32]",
		}, {
			heap:         NewHeap[int64](Max, 1),
			expectedType: "*heap.maxHeap[int64]",
		},
	}
	for _, tc := range testcases {
		if reflect.TypeOf(tc.heap).String() != tc.expectedType {
			t.Errorf("New heap failed, got: %v, want: %v.", reflect.TypeOf(tc.heap).String(), tc.expectedType)
		}
	}
}

func TestNewMaxHeapFloat(t *testing.T) {
	testcases := []struct {
		heap         interface{}
		expectedType string
	}{
		{
			heap:         NewHeap[float32](Max, 1),
			expectedType: "*heap.maxHeap[float32]",
		}, {
			heap:         NewHeap[float64](Max, 1),
			expectedType: "*heap.maxHeap[float64]",
		},
	}
	for _, tc := range testcases {
		if reflect.TypeOf(tc.heap).String() != tc.expectedType {
			t.Errorf("New heap failed, got: %v, want: %v.", reflect.TypeOf(tc.heap).String(), tc.expectedType)
		}
	}
}

func TestNewMinHeapInt(t *testing.T) {
	testcases := []struct {
		heap         interface{}
		expectedType string
	}{
		{
			heap:         NewHeap[int](Min, 1),
			expectedType: "*heap.minHeap[int]",
		}, {
			heap:         NewHeap[int8](Min, 1),
			expectedType: "*heap.minHeap[int8]",
		}, {
			heap:         NewHeap[int16](Min, 1),
			expectedType: "*heap.minHeap[int16]",
		}, {
			heap:         NewHeap[int32](Min, 1),
			expectedType: "*heap.minHeap[int32]",
		}, {
			heap:         NewHeap[int64](Min, 1),
			expectedType: "*heap.minHeap[int64]",
		},
	}
	for _, tc := range testcases {
		if reflect.TypeOf(tc.heap).String() != tc.expectedType {
			t.Errorf("New heap failed, got: %v, want: %v.", reflect.TypeOf(tc.heap).String(), tc.expectedType)
		}
	}
}

func TestNewMinHeapFloat(t *testing.T) {
	testcases := []struct {
		heap         interface{}
		expectedType string
	}{
		{
			heap:         NewHeap[float32](Min, 1),
			expectedType: "*heap.minHeap[float32]",
		}, {
			heap:         NewHeap[float64](Min, 1),
			expectedType: "*heap.minHeap[float64]",
		},
	}
	for _, tc := range testcases {
		if reflect.TypeOf(tc.heap).String() != tc.expectedType {
			t.Errorf("New heap failed, got: %v, want: %v.", reflect.TypeOf(tc.heap).String(), tc.expectedType)
		}
	}
}

func TestNewNilHeap(t *testing.T) {
	h := NewHeap[int](HeapType(2), 8)
	if h != nil {
		t.Errorf("New heap failed, got: %v, want: <nil>.", h)
	}
}
