package day17

import "testing"

func TestDenom(t *testing.T) {
	if denom(0) != 1 ||
		denom(1) != 2 ||
		denom(2) != 4 ||
		denom(3) != 8 {
		t.Fail()
	}
}
