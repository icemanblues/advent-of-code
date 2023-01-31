package main

import (
	"testing"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

// TestWrapCube this test is very specific to the input file, since the WrapCube function is very specific to it
func TestWrapCube(t *testing.T) {
	grid, _, _, _ := Parse("input.txt")
	off := 10
	tests := []struct {
		name     string
		start    util.Point
		dir      Direction
		expected util.Point
		eDir     Direction
	}{
		{"1->6 edge start", util.NewPoint(50, 0), U, util.NewPoint(0, 150), R},
		{"1->6 edge mid", util.NewPoint(50+off, 0), U, util.NewPoint(0, 150+off), R},
		{"1->6 edge end", util.NewPoint(99, 0), U, util.NewPoint(0, 199), R},

		{"1->4 edge start", util.NewPoint(50, 0), L, util.NewPoint(0, 149), R},
		{"1->4 edge mid", util.NewPoint(50, 0+off), L, util.NewPoint(0, 149-off), R},
		{"1->4 edge end", util.NewPoint(50, 49), L, util.NewPoint(0, 100), R},

		{"2->6 edge start", util.NewPoint(100, 0), U, util.NewPoint(0, 199), U},
		{"2->6 edge mid", util.NewPoint(100+off, 0), U, util.NewPoint(off, 199), U},
		{"2->6 edge end", util.NewPoint(149, 0), U, util.NewPoint(49, 199), U},

		{"2->5 edge mid", util.NewPoint(149, 0+off), R, util.NewPoint(99, 149-off), L},

		{"2->3 edge mid", util.NewPoint(100+off, 49), D, util.NewPoint(99, 50+off), L},

		{"3->2 edge mid", util.NewPoint(99, 50+off), R, util.NewPoint(100+off, 49), U},

		{"3->4 edge mid", util.NewPoint(50, 50+off), L, util.NewPoint(off, 100), D},

		{"4->1 edge mid", util.NewPoint(0, 100+off), L, util.NewPoint(50, 49-off), R},

		{"4->3 edge mid", util.NewPoint(0+off, 100), U, util.NewPoint(50, 50+off), R},

		{"5->2 edge mid", util.NewPoint(99, 100+off), R, util.NewPoint(149, 49-off), L},

		{"5->6 edge mid", util.NewPoint(50+off, 149), D, util.NewPoint(49, 150+off), L},

		{"6->1 edge mid", util.NewPoint(0, 150+off), L, util.NewPoint(50+off, 0), D},
		{"6->2 edge mid", util.NewPoint(0+off, 199), D, util.NewPoint(100+off, 0), D},
		{"6->5 edge mid", util.NewPoint(49, 150+off), R, util.NewPoint(50+off, 149), U},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			end, endDir := WrapCube(test.start, test.start, test.dir, grid, test.start) // we dont need all of these // I know the function
			if end != test.expected {
				t.Errorf("Incorrect Cube Wrap! actual: %v expected: %v\n", end, test.expected)
			}
			if endDir != test.eDir {
				t.Errorf("Incorrect Facing! actual: %c expected: %c\n", endDir, test.eDir)
			}
		})
	}
}
