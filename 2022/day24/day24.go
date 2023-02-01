package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "24"
	dayTitle = "Blizzard Basin"
)

var directions map[rune]util.Point = map[rune]util.Point{
	'^': util.NewPoint(0, -1),
	'v': util.NewPoint(0, 1),
	'<': util.NewPoint(-1, 0),
	'>': util.NewPoint(1, 0),
	' ': util.NewPoint(0, 0),
}

type WindSet map[util.Point]struct{}

type WindState struct {
	up, down, left, right WindSet
}

func (ws WindState) contains(p util.Point) bool {
	_, uOk := ws.up[p]
	_, dOk := ws.down[p]
	_, lOk := ws.left[p]
	_, rOk := ws.right[p]
	return uOk || dOk || lOk || rOk
}

func StormEqual(s1, s2 WindState) bool {
	// up
	for p := range s1.up {
		if _, ok := s2.up[p]; !ok {
			return false
		}
	}

	// down
	for p := range s1.down {
		if _, ok := s2.down[p]; !ok {
			return false
		}
	}

	// left
	for p := range s1.left {
		if _, ok := s2.left[p]; !ok {
			return false
		}
	}

	// right
	for p := range s1.right {
		if _, ok := s2.right[p]; !ok {
			return false
		}
	}

	return true
}

type State struct {
	round     int
	curr      util.Point
	blizzards WindState
}

type StormGen struct {
	grid  [][]rune
	cache []WindState // round number to WindState
	wrap  int
}

func NewStormGen(grid [][]rune, start WindState) *StormGen {
	return &StormGen{grid, []WindState{start}, 0}
}

func (sg *StormGen) Next(rnd int) (WindState, int) {
	// check cache first
	if rnd < len(sg.cache) {
		return sg.cache[rnd], rnd
	}

	// loop lookup?
	if sg.wrap != 0 {
		idx := rnd % sg.wrap
		return sg.cache[idx], idx
	}

	// we assume that the prior one is here. it should due how we call it.
	// if you call it in advance thats on you
	if rnd == len(sg.cache) {
		prev := sg.cache[rnd-1]
		storm := WindState{
			Blow(prev.up, '^', sg.grid),
			Blow(prev.down, 'v', sg.grid),
			Blow(prev.left, '<', sg.grid),
			Blow(prev.right, '>', sg.grid),
		}

		if StormEqual(sg.cache[0], storm) {
			sg.wrap = len(sg.cache)
			return sg.cache[0], 0
		}

		sg.cache = append(sg.cache, storm)
		return storm, rnd
	}

	fmt.Printf("You are calling out of order. Looking for round %v when latest is %v\n", rnd, len(sg.cache)-1)
	return sg.cache[0], 0
}

func Blow(wind WindSet, dir rune, grid [][]rune) WindSet {
	next := make(WindSet)
	offset := directions[dir]
	for w := range wind {
		nw := util.NewPoint(w.X+offset.X, w.Y+offset.Y)
		// need to check for wrap around
		if lookup(grid, nw) == '#' {
			switch dir {
			case '^':
				nw.Y = len(grid) - 2
			case 'v':
				nw.Y = 1
			case '<':
				nw.X = len(grid[nw.X]) - 2
			case '>':
				nw.X = 1
			}
		}
		next[nw] = struct{}{}
	}
	return next
}

func lookup(grid [][]rune, p util.Point) rune {
	if p.Y < 0 || p.Y >= len(grid) || p.X < 0 || p.X >= len(grid[p.Y]) {
		return '#'
	}
	return grid[p.Y][p.X]
}

func Parse(filename string) ([][]rune, WindState, util.Point, util.Point) {
	grid, _ := util.ReadRuneput(filename)
	// find start and end
	var start, end util.Point
	for x, r := range grid[0] {
		if r == '.' {
			start = util.NewPoint(x, 0)
			break
		}
	}
	for x, r := range grid[len(grid)-1] {
		if r == '.' {
			end = util.NewPoint(x, len(grid)-1)
			break
		}
	}

	// find all blizzards
	up := make(WindSet)
	down := make(WindSet)
	left := make(WindSet)
	right := make(WindSet)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			switch grid[y][x] {
			case '^':
				up[util.NewPoint(x, y)] = struct{}{}
			case '>':
				right[util.NewPoint(x, y)] = struct{}{}
			case '<':
				left[util.NewPoint(x, y)] = struct{}{}
			case 'v':
				down[util.NewPoint(x, y)] = struct{}{}
			}
		}
	}
	windState := WindState{up, down, left, right}
	return grid, windState, start, end
}

type Node struct {
	curr util.Point
	rnd  int
}

// Traverse BFS
func Traverse(grid [][]rune, startState State, end util.Point, gen *StormGen) State {
	visited := make(map[Node]struct{})
	q := []State{startState}
	for len(q) != 0 {
		state := q[0]
		q = q[1:]

		// are we at the end point
		if state.curr == end {
			return state
		}

		// advance the round and wind (remember to wrap around)
		rnd := state.round + 1
		nextBlizz, nidx := gen.Next(rnd)

		// choose one of 5 movement options
		// enqueue all 5 if valid
		for _, offset := range directions {
			next := util.NewPoint(state.curr.X+offset.X, state.curr.Y+offset.Y)
			if nextBlizz.contains(next) || lookup(grid, next) == '#' {
				continue
			}
			if _, ok := visited[Node{next, nidx}]; ok {
				continue // we've been here before
			}
			visited[Node{next, nidx}] = struct{}{}
			q = append(q, State{rnd, next, nextBlizz})
		}
	}

	// should never reach here
	fmt.Printf("Unable to solve the blizzard basin. Impossible!\n")
	return startState
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)

	grid, windState, start, end := Parse("input.txt")
	gen := NewStormGen(grid, windState)
	startState := State{0, start, windState}
	endState := Traverse(grid, startState, end, gen)
	fmt.Printf("Part 1: %v\n", endState.round)

	// go back to start
	startAgain := Traverse(grid, endState, start, gen)

	// go back to end
	endAgain := Traverse(grid, startAgain, end, gen)
	fmt.Printf("Part 2: %v\n", endAgain.round)

}
