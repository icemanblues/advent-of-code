package main

import (
	"testing"
)

// OtherTool(region, tool int, toolBadRegion map[int]int) int
func TestOtherTool(t *testing.T) {
	tests := []struct {
		region   int
		tool     int
		expected int
	}{
		{Rocky, Torch, Climbing},
		{Rocky, Climbing, Torch},

		{Wet, Neither, Climbing},
		{Wet, Climbing, Neither},

		{Narrow, Torch, Neither},
		{Narrow, Neither, Torch},
	}

	for _, test := range tests {
		newTool := OtherTool(test.region, test.tool)
		if newTool != test.expected {
			t.Errorf("Expected %v but Actual %v with start tool %v", test.expected, newTool, test.tool)
		}
	}
}
