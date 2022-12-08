package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "08"
	dayTitle = "Treetop Tree House"
)

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	countVisibleTrees, maxView := 0, 0
	grid, _ := util.ReadIntGrid("input.txt", "")
	for i:=0; i<len(grid);i++ {
		for j:=0; j<len(grid[i]); j++ {
			left, lv := true, j
			for k:= j-1; k>=0; k-- {
				if grid[i][j] <= grid[i][k] {
					left, lv = false, j-k
					break
				}
			}

			right, rv := true, len(grid[i])-1-j
			for k:= j+1; k<len(grid[i]); k++ {
				if grid[i][j] <= grid[i][k] {
					right, rv = false, k-j
					break
				}
			}

			up, uv := true, i
			for k:= i-1; k>=0; k-- {
				if grid[i][j] <= grid[k][j] {
					up, uv = false, i-k
					break
				}
			}

			down, dv := true, len(grid)-1-i
			for k:= i+1; k<len(grid); k++ {
				if grid[i][j] <= grid[k][j] {
					down, dv = false, k-i
					break
				}
			}

			if up || down || left || right {
				countVisibleTrees++
			}

			if viewScore := lv * rv * uv * dv; viewScore > maxView {
				maxView = viewScore
			}
		}
	}
	fmt.Printf("Part 1: %v\n", countVisibleTrees)
	fmt.Printf("Part 2: %v\n", maxView)
}
