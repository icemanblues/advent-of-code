package main

import (
	"fmt"
	"strconv"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "22"
	dayTitle = "Monkey Map"
)

type Grid map[util.Point]rune

type Direction rune

const (
	U Direction = 'U'
	R Direction = 'R'
	D Direction = 'D'
	L Direction = 'L'
)

func (d Direction) Score() int {
	switch d {
	case U:
		return 3
	case R:
		return 0
	case D:
		return 1
	case L:
		return 2
	default:
		fmt.Printf("Unknown direction for Score: %c\n", d)
		return 0
	}
}

var Movement map[Direction]util.Point = map[Direction]util.Point{
	U: util.NewPoint(0, -1),
	R: util.NewPoint(1, 0),
	D: util.NewPoint(0, 1),
	L: util.NewPoint(-1, 0),
}

func Move(p util.Point, d Direction) util.Point {
	offset, ok := Movement[d]
	if !ok {
		fmt.Printf("Unknown direction for Move: %c\n", d)
		return p
	}
	return util.NewPoint(p.X+offset.X, p.Y+offset.Y)
}

type Inst struct {
	Step int
	Turn rune
}

func (inst Inst) String() string {
	return fmt.Sprintf("{Step: %v Turn: %c}", inst.Step, inst.Turn)
}

const (
	LTurn  = 'L'
	RTurn  = 'R'
	NoTurn = ' '
)

func Turn(d Direction, t rune) Direction {
	switch t {
	case RTurn:
		switch d {
		case U:
			return R
		case R:
			return D
		case D:
			return L
		case L:
			return U
		default:
			fmt.Printf("Unknown direction for Turn: %c\n", d)
			return d
		}
	case LTurn:
		switch d {
		case U:
			return L
		case R:
			return U
		case D:
			return R
		case L:
			return D
		default:
			fmt.Printf("Unknown direction for Turn: %c\n", d)
			return d
		}
	case NoTurn:
		return d
	default:
		fmt.Printf("Unknown turn type for Turn: %c\n", t)
		return d
	}
}

func Parse(filename string) (Grid, []Inst, util.Point, util.Point) {
	parseMap := true
	grid := make(Grid)
	startX := -1
	maxX, maxY := 0, 0
	insts := make([]Inst, 0)
	numbers := make([]rune, 0)

	input, _ := util.ReadInput(filename)
	for y, line := range input {
		if line == "" {
			parseMap = false
			continue
		}

		// add it to the grid
		if parseMap {
			if y > maxY {
				maxY = y
			}
			for x, r := range line {
				if r == '.' || r == '#' {
					p := util.NewPoint(x, y)
					grid[p] = r
				}
				if x > maxX {
					maxX = x
				}
				if startX == -1 && r == '.' {
					startX = x
				}
			}
			continue
		}

		// parse the movement instructions
		for _, r := range line {
			if r == 'L' || r == 'R' {
				i, _ := strconv.Atoi(string(numbers))
				numbers = nil
				insts = append(insts, Inst{i, ' '})
				insts = append(insts, Inst{0, r})
			} else {
				numbers = append(numbers, r)
			}
		}
		if len(numbers) != 0 {
			i, _ := strconv.Atoi(string(numbers))
			numbers = nil
			insts = append(insts, Inst{i, ' '})
		}
	}

	return grid, insts, util.NewPoint(startX, 0), util.NewPoint(maxX, maxY)
}

func Score(p util.Point, d Direction) int {
	return 1000*(p.Y+1) + 4*(p.X+1) + d.Score()
}

// WrapFunc Function on how to handle wrapping around the map. It should always return a point thats in the grid
// prev is in the grid, curr is NOT in the grid. We moved off the map, triggering the need to wrap around
type WrapFunc func(prev, curr util.Point, d Direction, grid Grid, bounds util.Point) (util.Point, Direction)

func Traverse(grid Grid, insts []Inst, start util.Point, dir Direction, bounds util.Point, wrap WrapFunc) (util.Point, Direction) {
	curr := start
	for _, inst := range insts {
		// move
		for i := 0; i < inst.Step; i++ {
			next, nd := Move(curr, dir), dir
			// need to check that next is valid, (in grid, otherwise wrap)
			_, ok := grid[next]
			if !ok {
				next, nd = wrap(curr, next, dir, grid, bounds)
			}
			tile := grid[next]
			if tile == '#' { // stop moving at a wall
				break
			}

			curr, dir = next, nd // all good so make the move
		}
		// turn
		dir = Turn(dir, inst.Turn)
	}
	return curr, dir
}

func WrapGrid(prev, curr util.Point, d Direction, grid Grid, bounds util.Point) (util.Point, Direction) {
	_, ok := grid[curr]
	for ; !ok; _, ok = grid[curr] {
		curr = Move(curr, d)
		if curr.X < 0 {
			curr.X = bounds.X
		}
		if curr.X > bounds.X {
			curr.X = 0
		}
		if curr.Y < 0 {
			curr.Y = bounds.Y
		}
		if curr.Y > bounds.Y {
			curr.Y = 0
		}
	}
	return curr, d
}

func part1() {
	grid, insts, start, bounds := Parse("input.txt")
	dir := R
	p, d := Traverse(grid, insts, start, dir, bounds, WrapGrid)
	fmt.Printf("Part 1: %v\n", Score(p, d))
}

// cube sides are 50 x 50
const cubeSize = 50

// WrapCube a WrapFunc specific for the cube orientation of my input
// my input.txt cube has this shape
//
// .12
// .3.
// 45.
// 6..
func WrapCube(prev, curr util.Point, dir Direction, grid Grid, bounds util.Point) (util.Point, Direction) {
	// figure out which cube curr is in
	cube := 0
	cx, cy := prev.X/cubeSize, prev.Y/cubeSize
	if cx == 1 && cy == 0 {
		cube = 1
	} else if cy == 0 && cx == 2 {
		cube = 2
	} else if cy == 1 && cx == 1 {
		cube = 3
	} else if cy == 2 && cx == 0 {
		cube = 4
	} else if cy == 2 && cx == 1 {
		cube = 5
	} else if cy == 3 && cx == 0 {
		cube = 6
	} else {
		fmt.Printf("WrapCube: Unable to determine which cube for this point: %v\n", prev)
		return prev, dir
	}

	if cube == 1 && dir == U { // go to cube 6
		d := prev.X - cubeSize
		return util.NewPoint(0, 3*cubeSize+d), R
	} else if cube == 1 && dir == L { // go to cube 4
		d := prev.Y
		return util.NewPoint(0, 3*cubeSize-1-d), R
	} else if cube == 2 && dir == U { // go to cube 6
		d := prev.X - 2*cubeSize
		return util.NewPoint(d, 4*cubeSize-1), U
	} else if cube == 2 && dir == R { // go to cube 5
		d := prev.Y
		return util.NewPoint(2*cubeSize-1, 3*cubeSize-1-d), L
	} else if cube == 2 && dir == D { // go to cube 3
		d := prev.X - 2*cubeSize
		return util.NewPoint(2*cubeSize-1, cubeSize+d), L
	} else if cube == 3 && dir == L { // go to cube 4
		d := prev.Y - cubeSize
		return util.NewPoint(d, 2*cubeSize), D
	} else if cube == 3 && dir == R { // go to cube 2
		d := prev.Y - cubeSize
		return util.NewPoint(2*cubeSize+d, cubeSize-1), U
	} else if cube == 4 && dir == U { // go to cube 3
		d := prev.X
		return util.NewPoint(cubeSize, cubeSize+d), R
	} else if cube == 4 && dir == L { // go to cube 1
		d := prev.Y - 2*cubeSize
		return util.NewPoint(cubeSize, cubeSize-1-d), R
	} else if cube == 5 && dir == R { // go to cube 2
		d := prev.Y - 2*cubeSize
		return util.NewPoint(3*cubeSize-1, cubeSize-1-d), L
	} else if cube == 5 && dir == D { // go to cube 6
		d := prev.X - cubeSize
		return util.NewPoint(cubeSize-1, 3*cubeSize+d), L
	} else if cube == 6 && dir == L { // go to cube 1
		d := prev.Y - 3*cubeSize
		return util.NewPoint(cubeSize+d, 0), D
	} else if cube == 6 && dir == D { // go to cube 2
		d := prev.X
		return util.NewPoint(2*cubeSize+d, 0), D
	} else if cube == 6 && dir == R { // go to cube 5
		d := prev.Y - 3*cubeSize
		return util.NewPoint(cubeSize+d, 3*cubeSize-1), U
	}

	fmt.Printf("WrapCube: Missing a cube wrap. This point should never be reached: %v\n", prev)
	return prev, dir
}

// 10276 too low
func part2() {
	grid, insts, start, bounds := Parse("input.txt")
	dir := R
	p, d := Traverse(grid, insts, start, dir, bounds, WrapCube)
	fmt.Printf("Part 2: %v\n", Score(p, d))
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
