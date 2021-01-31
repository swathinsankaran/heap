package heap

import "testing"

func TestSwap(t *testing.T) {
	a, b := 10, 9
	swap(&a, &b)
	if a != 9 || b != 10 {
		t.Errorf("Swap was incorrect, got: a=%d b=%d, want: a=%d b=%d.", a, b, 9, 10)
	}
}
