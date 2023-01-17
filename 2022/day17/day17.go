package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "17"
	dayTitle = "Pyroclastic Flow"
)

const (
	TowerLeft      = 0
	TowerRight     = 6
	NextRockHeight = 4
)

var JetMove map[rune]int = map[rune]int{
	'<': -1,
	'>': 1,
}

type WindGen struct {
	input []rune
	index int
}

func (w *WindGen) Next() (rune, int) {
	r, i := w.input[w.index], w.index
	w.index = (w.index + 1) % len(w.input)
	return r, i
}

type RockSet map[util.Point]struct{}

type Rock struct {
	Points RockSet
	Left, Right,
	Up, Down int
}

func (r Rock) Jet(rr rune) Rock {
	i := JetMove[rr]
	points := make(RockSet)
	for p := range r.Points {
		points[util.NewPoint(p.X+i, p.Y)] = struct{}{}
	}
	return Rock{points, r.Left + i, r.Right + i, r.Up, r.Down}
}

func (r Rock) Fall() Rock {
	points := make(RockSet)
	for p := range r.Points {
		points[util.NewPoint(p.X, p.Y-1)] = struct{}{}
	}
	return Rock{points, r.Left, r.Right, r.Up - 1, r.Down - 1}
}

func (r *Rock) Add(a Rock) {
	for p := range a.Points {
		r.Points[p] = struct{}{}
	}
	if a.Up > r.Up {
		r.Up = a.Up
	}
	if a.Down < r.Down {
		r.Down = a.Down
	}
	if a.Left < r.Left {
		r.Left = a.Left
	}
	if a.Right > r.Right {
		r.Right = a.Right
	}
}

func (r Rock) Collision(a Rock) bool {
	for p := range a.Points {
		if _, ok := r.Points[p]; ok {
			return true
		}
	}
	return false
}

func (r Rock) Heights() [7]int {
	var h [7]int
	for i := 0; i < 7; i++ {
		max := 0
		for p := range r.Points {
			if p.X == i && p.Y > max {
				max = p.Y
			}
		}
		h[i] = max
	}

	// now normalize it
	max := h[0]
	for i := 1; i < len(h); i++ {
		if h[i] > max {
			max = h[i]
		}
	}
	for i := 0; i < len(h); i++ {
		h[i] = h[i] - max
	}
	return h
}

func (r Rock) Print() {
	for y := r.Up + 3; y >= r.Down; y-- {
		fmt.Printf("|")
		for x := TowerLeft; x <= TowerRight; x++ {
			if _, ok := r.Points[util.NewPoint(x, y)]; ok {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("|  %v\n", y)
	}
	fmt.Println()
}

// Each rock appears so that its left edge is two units away from the left wall
// bottom edge is three units above the highest rock in the room (or the floor, if there isn't one)
type RockConstructor func(int, int) Rock

type RockGen struct {
	constructors []RockConstructor
	index        int
}

func (rg *RockGen) Next(a, b int) (Rock, int) {
	rock, idx := rg.constructors[rg.index](a, b), rg.index
	rg.index = (rg.index + 1) % len(rg.constructors)
	return rock, idx
}

func NewRockMinus(left, bottom int) Rock {
	points := make(RockSet)
	for i := left; i < left+4; i++ {
		points[util.NewPoint(i, bottom)] = struct{}{}
	}
	return Rock{points, left, left + 3, bottom, bottom}
}
func NewRockPlus(left, bottom int) Rock {
	points := make(RockSet)
	points[util.NewPoint(left+1, bottom+2)] = struct{}{}
	points[util.NewPoint(left, bottom+1)] = struct{}{}
	points[util.NewPoint(left+1, bottom+1)] = struct{}{}
	points[util.NewPoint(left+2, bottom+1)] = struct{}{}
	points[util.NewPoint(left+1, bottom)] = struct{}{}
	return Rock{points, left, left + 2, bottom + 2, bottom}
}
func NewRockElbow(left, bottom int) Rock {
	points := make(RockSet)
	points[util.NewPoint(left, bottom)] = struct{}{}
	points[util.NewPoint(left+1, bottom)] = struct{}{}
	points[util.NewPoint(left+2, bottom)] = struct{}{}
	points[util.NewPoint(left+2, bottom+1)] = struct{}{}
	points[util.NewPoint(left+2, bottom+2)] = struct{}{}
	return Rock{points, left, left + 2, bottom + 2, bottom}
}
func NewRockSchlong(left, bottom int) Rock {
	points := make(RockSet)
	points[util.NewPoint(left, bottom)] = struct{}{}
	points[util.NewPoint(left, bottom+1)] = struct{}{}
	points[util.NewPoint(left, bottom+2)] = struct{}{}
	points[util.NewPoint(left, bottom+3)] = struct{}{}
	return Rock{points, left, left, bottom + 3, bottom}
}
func NewRockBox(left, bottom int) Rock {
	points := make(RockSet)
	points[util.NewPoint(left, bottom)] = struct{}{}
	points[util.NewPoint(left+1, bottom)] = struct{}{}
	points[util.NewPoint(left, bottom+1)] = struct{}{}
	points[util.NewPoint(left+1, bottom+1)] = struct{}{}
	return Rock{points, left, left + 1, bottom + 1, bottom}
}

func NewRockGen() RockGen {
	return RockGen{[]RockConstructor{NewRockMinus, NewRockPlus, NewRockElbow, NewRockSchlong, NewRockBox}, 0}
}

func ParseWindGen(filename string) WindGen {
	input, _ := util.ReadInput(filename)
	return WindGen{[]rune(input[0]), 0}
}

func NewTower() Rock {
	floorPoints := make(RockSet)
	for x := TowerLeft; x <= TowerRight; x++ {
		floorPoints[util.NewPoint(x, 0)] = struct{}{}
	}
	return Rock{floorPoints, TowerLeft, TowerRight, 0, 0}
}

func part1() {
	windGen := ParseWindGen("input.txt")
	rockGen := NewRockGen()
	tower := NewTower()
	numDrop := 2022

	// blow the wind and drop the rock until it stops falling
	for rockDrop := 0; rockDrop < numDrop; rockDrop++ {
		rock, _ := rockGen.Next(2, tower.Up+4)
		for true {
			wind, _ := windGen.Next()
			windRock := rock.Jet(wind)
			if windRock.Left >= TowerLeft && windRock.Right <= TowerRight && !tower.Collision(windRock) {
				rock = windRock
			}

			dropRock := rock.Fall()
			if tower.Collision(dropRock) {
				tower.Add(rock)
				break
			}
			rock = dropRock
		}
	}
	fmt.Printf("Part 1: %v\n", tower.Up)
}

type State struct {
	wind, rock int
	relH       [7]int
}

func part2() {
	windGen := ParseWindGen("input.txt")
	rockGen := NewRockGen()
	tower := NewTower()
	numDrop := 1000000000000

	states := make([]State, 0)
	statesRock := make(map[State]int)
	statesHeight := make(map[State]int)
	loopStart, loopEnd, loopEndHeight := 0, 0, 0
	// blow the wind and drop the rock until it stops falling
search:
	for rockDrop := 0; rockDrop < numDrop; rockDrop++ {
		rock, rockIdx := rockGen.Next(2, tower.Up+4)

		for true {
			wind, windIdx := windGen.Next()
			windRock := rock.Jet(wind)
			if windRock.Left >= TowerLeft && windRock.Right <= TowerRight && !tower.Collision(windRock) {
				rock = windRock
			}

			dropRock := rock.Fall()
			if tower.Collision(dropRock) {
				tower.Add(rock)
				// add state too
				heights := tower.Heights()
				state := State{windIdx, rockIdx, heights}
				if n, ok := statesRock[state]; ok { // loop detected
					loopStart = n
					loopEnd = rockDrop
					loopEndHeight = tower.Up
					break search
				} else {
					states = append(states, state)
					statesRock[state] = rockDrop
					statesHeight[state] = tower.Up
				}
				break
			}
			rock = dropRock
		}
	}

	// do some math to calculate the correct answer
	goTo := numDrop
	height := statesHeight[states[loopStart]]
	goTo -= loopStart

	loopLen := loopEnd - loopStart
	cycles := goTo / loopLen
	remainder := goTo % loopLen
	cycleInc := loopEndHeight - statesHeight[states[loopStart]]
	height += cycles * cycleInc

	remainderIdx := loopStart + remainder
	remainderHeight := statesHeight[states[remainderIdx]] - statesHeight[states[loopStart]]
	height += remainderHeight - 1
	fmt.Printf("Part 2: %v\n", height)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
