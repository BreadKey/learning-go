package operator

import "testing"

func TestToNegative(t *testing.T) {
	i := 10
	if i + ToNegative(i) != 0 {
		t.Error("Wrong result")
	}
}