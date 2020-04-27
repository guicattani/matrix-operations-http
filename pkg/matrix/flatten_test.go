package matrix

import "testing"

func TestFlatten(t *testing.T) {
	mtx := Matrix{
		{"1", "2"},
		{"3", "4"},
	}

	response := mtx.Flatten()
	if response != "1,2,3,4\n" {
		t.Errorf("Expected response printed in one line.")
	}
}
