package sdk

import "testing"

func TestCreatesClient(t *testing.T) {
	var result = PrintHelloWorld()

	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}
