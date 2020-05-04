package matrix

import "testing"

func TestEcho(t *testing.T) {
	m := Matrix{
		{"1", "2"},
		{"3", "4"},
	}

	response := m.Echo()
	if response.Message != "1,2\n3,4\n" {
		t.Errorf("Expected response to be the same as inputted: expected\n%s got\n%s", "1,2\n3,4\n", response.Message)
	}
}
