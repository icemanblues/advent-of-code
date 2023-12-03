package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "03"
	dayTitle = "Gear Ratios"
)

func isInt(r rune) bool {
	switch r {
	case '0':
		return true
	case '1':
		return true
	case '2':
		return true
	case '3':
		return true
	case '4':
		return true
	case '5':
		return true
	case '6':
		return true
	case '7':
		return true
	case '8':
		return true
	case '9':
		return true
	default:
		return false
	}
}

func isSymbol(r rune) bool {
	return r != '.' && !isInt(r)
}

func adjPoints(p util.Point) []util.Point {
	adj := make([]util.Point, 0, 8)
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			adj = append(adj, util.Point{X: p.X + dx, Y: p.Y + dy})
		}
	}
	return adj
}

type Span struct {
	Y, Xmin, Xmax int
}

func (s Span) contains(p util.Point) bool {
	return p.Y == s.Y && p.X >= s.Xmin && p.X <= s.Xmax
}

func GetIntSpans(grid [][]rune) map[Span]int {
	spans := make(map[Span]int)
	for y := 0; y < len(grid); y++ {
		x := 0
		for x < len(grid[y]) {
			if isInt(grid[y][x]) {
				start := x
				for x = x + 1; x < len(grid[y]); x++ {
					if !isInt(grid[y][x]) {
						break
					}
				}
				spans[Span{y, start, x - 1}] = util.MustAtoi(string(grid[y][start:x]))
			} else {
				x++
			}
		}
	}
	return spans
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	input, _ := util.ReadRuneput("input.txt")
	intSpans := GetIntSpans(input)
	spans := make(map[Span]struct{})
	gearRatio := 0
	// find symbols, scan around them for matching spans
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if isSymbol(input[y][x]) {
				spanCount := make(map[Span]struct{})
				for _, p := range adjPoints(util.Point{X: x, Y: y}) {
					for span := range intSpans {
						if span.contains(p) {
							spans[span] = struct{}{}
							spanCount[span] = struct{}{}
						}
					}
				}
				// check if gear
				if len(spanCount) == 2 {
					product := 1
					for span := range spanCount {
						product *= intSpans[span]
					}
					gearRatio += product
				}
			}
		}
	}

	sum := 0
	for span := range spans {
		sum += intSpans[span]
	}
	fmt.Printf("Part 1: %v\n", sum)
	fmt.Printf("Part 2: %v\n", gearRatio)
}
