package simplemath

import "testing"

func TestSqrt(t *testing.T) {
	r := Sqrt(4)
	if r != 4 {
		t.Errorf("Sqrt(4) failed,Got %v,excepted 2.")
	}
}
