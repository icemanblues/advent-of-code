package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "08"
	dayTitle = "Seven Segment Search"
)

type runeset map[rune]struct{}

type Segment struct {
	in, out []runeset
}

var siglenToNum map[int]int = map[int]int{
	2: 1,
	4: 4,
	3: 7,
	7: 8,
}

func Parse(filename string) []Segment {
	input, _ := util.ReadInput(filename)
	var segments []Segment
	for _, in := range input {
		parts := strings.Split(in, " | ")
		i := strings.Split(parts[0], " ")
		ii := make([]runeset, 0, len(i))
		for _, s := range i {
			rs := make(runeset)
			for _, r := range s {
				rs[r] = struct{}{}
			}
			ii = append(ii, rs)
		}

		j := strings.Split(parts[1], " ")
		jj := make([]runeset, 0, len(j))
		for _, s := range j {
			rs := make(runeset)
			for _, r := range s {
				rs[r] = struct{}{}
			}
			jj = append(jj, rs)
		}
		segments = append(segments, Segment{ii, jj})
	}
	return segments
}

func part1() {
	input := Parse("input.txt")
	count := 0
	for _, segment := range input {
		for _, out := range segment.out {
			if _, ok := siglenToNum[len(out)]; ok {
				count++
			}
		}
	}

	fmt.Printf("Part 1: %v\n", count)
}

func part2() {
	input := Parse("input.txt")
	sum := 0
	for _, segment := range input {
		numbers := make([]runeset, 10, 10)
		// find 1 // unique, len is 2
		for i, s := range segment.in {
			if len(s) == 2 {
				numbers[1] = s
				segment.in = append(segment.in[:i], segment.in[i+1:]...)
				break
			}
		}

		// find 4 // unique, len is 4
		for i, s := range segment.in {
			if len(s) == 4 {
				numbers[4] = s
				segment.in = append(segment.in[:i], segment.in[i+1:]...)
				break
			}
		}

		// find 7 // unique, len is 3
		for i, s := range segment.in {
			if len(s) == 3 {
				numbers[7] = s
				segment.in = append(segment.in[:i], segment.in[i+1:]...)
				break
			}
		}

		// find 8 // unique, len is 8
		for i, s := range segment.in {
			if len(s) == 7 {
				numbers[8] = s
				segment.in = append(segment.in[:i], segment.in[i+1:]...)
				break
			}
		}

		// find 3 // len is 5, overlaps with 1
		for i, s := range segment.in {
			if len(s) == 5 {
				overlap := true
				for r := range numbers[1] {
					_, ok := s[r]
					overlap = overlap && ok
				}
				if overlap {
					numbers[3] = s
					segment.in = append(segment.in[:i], segment.in[i+1:]...)
					break
				}
			}
		}

		// find 9 // len is 6, overlaps with 3
		for i, s := range segment.in {
			if len(s) == 6 {
				overlap := true
				for r := range numbers[3] {
					_, ok := s[r]
					overlap = overlap && ok
				}
				if overlap {
					numbers[9] = s
					segment.in = append(segment.in[:i], segment.in[i+1:]...)
					break
				}
			}
		}

		// find 0 // len is 6, overlaps with 1
		for i, s := range segment.in {
			if len(s) == 6 {
				overlap := true
				for r := range numbers[1] {
					_, ok := s[r]
					overlap = overlap && ok
				}
				if overlap {
					numbers[0] = s
					segment.in = append(segment.in[:i], segment.in[i+1:]...)
					break
				}
			}
		}

		// find 6 // len is 6 (only one remaining)
		for i, s := range segment.in {
			if len(s) == 6 {
				numbers[6] = s
				segment.in = append(segment.in[:i], segment.in[i+1:]...)
				break
			}
		}

		// find 5 // len is 5, missing one item from 9
		for i, s := range segment.in {
			if len(s) == 5 {
				// my runeset intersection (actually exclusion)
				count := 0
				for r := range s {
					if _, ok := numbers[9][r]; ok {
						count++
					}
				}
				if count == 5 {
					numbers[5] = s
					segment.in = append(segment.in[:i], segment.in[i+1:]...)
					break
				}
			}
		}

		// find 2 // len is 5 (only one remaining)
		for i, s := range segment.in {
			if len(s) == 5 {
				numbers[2] = s
				segment.in = append(segment.in[:i], segment.in[i+1:]...)
				break
			}
		}

		// decode the out value
		pos := 1000
		acc := 0
		for _, rs := range segment.out {
			for i, n := range numbers {
				if reflect.DeepEqual(rs, n) {
					acc = acc + (i * pos)
					pos = pos / 10
					break
				}
			}
		}
		sum += acc
	}

	fmt.Printf("Part 2: %v\n", sum)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
