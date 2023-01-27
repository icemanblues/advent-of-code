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

type WrapFunc func(util.Point, Direction, Grid, util.Point) (util.Point, Direction)

func Traverse(grid Grid, insts []Inst, start util.Point, dir Direction, bounds util.Point, wrap WrapFunc) (util.Point, Direction) {
	curr := start
	for _, inst := range insts {
		// move
		for i := 0; i < inst.Step; i++ {
			next := Move(curr, dir)
			// need to check that next is valid
			_, ok := grid[next]
			if !ok {
				next, dir = wrap(next, dir, grid, bounds)
				// TODO: if this is a wall, then need to rollback
				// dir is lost, I didn't save it
				//next = grid.Wrap(next, dir, bounds)
			}
			tile := grid[next]
			if tile == '#' { // stop moving at a wall
				break
			}

			curr = next // all good so make the move
		}
		// turn
		dir = Turn(dir, inst.Turn)
	}
	return curr, dir
}

// type WrapFunc func(util.Point, Direction, Grid, util.Point)
func WrapGrid(curr util.Point, d Direction, grid Grid, bounds util.Point) (util.Point, Direction) {
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
//
// .12
// .3.
// 45.
// 6..

func WrapCube(curr util.Point, d Direction, grid Grid, bounds util.Point) (util.Point, Direction) {
	return curr, d
}

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
