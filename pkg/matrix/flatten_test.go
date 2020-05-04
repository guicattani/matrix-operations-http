package matrix

import "testing"

func TestFlatten(t *testing.T) {
	m := Matrix{
		{"1", "2"},
		{"3", "4"},
	}

	response := m.Flatten()
	if response.Message != "1,2,3,4\n" {
		t.Errorf("Expected response printed in one line: expected\n%s got\n%s", "1,2,3,4\n", response.Message)
	}
}
