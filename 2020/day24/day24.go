package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "24"
	dayTitle = "Lobby Layout"
)

type Axial struct {
	q, r int
}

func (a Axial) S() int {
	return -a.q - a.r
}

func (a Axial) Move(m string) Axial {
	switch m {
	case "e":
		return Axial{a.q + 1, a.r}
	case "ne":
		return Axial{a.q + 1, a.r - 1}
	case "se":
		return Axial{a.q, a.r + 1}
	case "w":
		return Axial{a.q - 1, a.r}
	case "nw":
		return Axial{a.q, a.r - 1}
	case "sw":
		return Axial{a.q - 1, a.r + 1}
	}
	panic("Illegal movement! " + m)
}

func (a Axial) Adj() []Axial {
	return []Axial{
		a.Move("e"),
		a.Move("ne"),
		a.Move("se"),
		a.Move("w"),
		a.Move("nw"),
		a.Move("sw"),
	}
}

type Floor map[Axial]bool

func countBlack(floor Floor) int {
	count := 0
	for _, isBlack := range floor {
		if isBlack {
			count++
		}
	}
	return count
}

func paintFloor(floor Floor, reference Axial, steps []string) {
	for _, step := range steps {
		curr := reference
		i := 0
		for i < len(step) {
			move := step[i : i+1]
			switch move {
			case "e":
				curr = curr.Move(move)
				i++
				continue
			case "w":
				curr = curr.Move(move)
				i++
				continue
			}

			move = step[i : i+2]
			switch move {
			case "ne":
				curr = curr.Move(move)
				i += 2
				continue
			case "se":
				curr = curr.Move(move)
				i += 2
				continue
			case "nw":
				curr = curr.Move(move)
				i += 2
				continue
			case "sw":
				curr = curr.Move(move)
				i += 2
				continue
			}

			panic("What kind of movement is this? " + move)
		}

		// flip the item, false is default value for bool. so treat this bool as IsBlack
		floor[curr] = !floor[curr]
	}
}

func GameOfPaint(floor Floor, days int) Floor {
	curr := floor
	for i := 0; i < days; i++ {
		next := make(Floor)

		for axial := range curr {
			check := make(map[Axial]struct{})
			adj := axial.Adj()
			for _, a := range adj {
				check[a] = struct{}{}
			}
			// need the adj of the adj too
			for _, a := range adj {
				for _, aa := range a.Adj() {
					check[aa] = struct{}{}
				}
			}

			// apply the rules for all of the ones that we want to check
			for hex := range check {
				hexAdj := hex.Adj()
				isBlack := curr[hex]
				black := 0
				for _, a := range hexAdj {
					if curr[a] {
						black++
					}
				}

				if isBlack && black == 0 {
					// pass, since we are flipping black, and not add it to the floor map
				} else if isBlack && black > 2 {
					// pass, since we are flipping black, and not add it to the floor map
				} else if !isBlack && black == 2 {
					next[hex] = true
				} else { // carry over the prior value, only if its black
					if curr[hex] {
						next[hex] = curr[hex]
					}
				}
			}
		}

		curr = next
	}
	return curr
}

func part1() {
	lines, _ := util.ReadInput("input.txt")
	reference := Axial{0, 0}
	floor := make(Floor)
	paintFloor(floor, reference, lines)
	fmt.Printf("Part 1: %v\n", countBlack(floor))
}

func part2() {
	lines, _ := util.ReadInput("input.txt")
	reference := Axial{0, 0}
	floor := make(Floor)
	paintFloor(floor, reference, lines)
	floor = GameOfPaint(floor, 100)
	fmt.Printf("Part 2: %v\n", countBlack(floor))
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
