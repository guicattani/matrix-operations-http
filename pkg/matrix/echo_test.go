package matrix

import "testing"

func TestEcho(t *testing.T) {
	mtx := Matrix{
		{"1", "2"},
		{"3", "4"},
	}

	response := mtx.Echo()
	if response != "1,2\n3,4\n" {
		t.Errorf("Expected response to be the same as inputted.")
	}
}
