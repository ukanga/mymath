package mymath

import "testing"

func TestCube(t *testing.T) {
	const in, out = 2, 8

	if x := Cube(in); x != out {
		t.Errorf("Cube(%v) = %v. want %v", in, x, out)
	}
}
