package main

import (
	"strings"
	"testing"
)

func TestGreet(t *testing.T) {
	// Call the Greet function to get the output
	output := Greet()

	// Check if the output contains "your_name_here"
	if strings.Contains(output, "your_name_here") {
		t.Errorf("Expected the greeting message to not contain 'your_name_here', but got: %s", output)
	}
}
