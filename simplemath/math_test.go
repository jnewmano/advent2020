package simplemath

import (
	"testing"
)

func TestGCD(t *testing.T) {
	if GCD(3, 9) != 3 {
		t.Fatalf("expected GCD to be 3\n")
	}
	if GCD(7, 19) != 1 {
		t.Fatalf("expected GCD to be 1\n")
	}
	if GCD(500, 5, 1000) != 5 {
		t.Fatalf("expected GCD to be 5\n")
	}
	if GCD(9, 27, 900, 27000) != 9 {
		t.Fatalf("expected GCD to be 9\n")
	}
}
