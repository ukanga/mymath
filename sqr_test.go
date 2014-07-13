package mymath

import "testing"

func TestSqr(t *testing.T) {
	const in, out = 2, 4

	if x := Sqr(in); x != out {
		t.Errorf("Sqr(%v) = %v. want %v", in, x, out)
	}
}
