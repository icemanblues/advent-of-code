package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "02"
	dayTitle = "I Was Told There Would Be No Math"
)

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	lines, _ := util.ReadInput("input.txt")
	wrappingPaper := 0
	ribbon := 0
	for _, line := range lines {
		dim := make([]int, 0, 3)
		for _, s := range strings.Split(line, "x") {
			i, _ := strconv.Atoi(s)
			dim = append(dim, i)
		}
		sort.Slice(dim, func(i, j int) bool {
			return dim[i] < dim[j]
		})

		lw := dim[0] * dim[1]
		wh := dim[0] * dim[2]
		lh := dim[1] * dim[2]
		min := lw
		if wh < min {
			min = wh
		}
		if lh < min {
			min = lh
		}
		area := 2 * (lw + wh + lh)
		wrappingPaper += area + min

		ribbonWrap := 2 * (dim[0] + dim[1])
		bow := dim[0] * dim[1] * dim[2]
		ribbon += ribbonWrap + bow
	}
	fmt.Printf("Part 1: %v\n", wrappingPaper)
	fmt.Printf("Part 2: %v\n", ribbon)
}
