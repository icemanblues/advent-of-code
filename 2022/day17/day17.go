package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "17"
	dayTitle = "Pyroclastic Flow"
)

type WindGen struct {
	input []rune
	index int
}

func (w *WindGen) Next() rune {
	r := w.input[w.index]
	w.index = (w.index + 1) % len(w.input)
	return r
}

var JetMove map[rune]int = map[rune]int{
	'<': -1,
	'>': 1,
}

type RockSet map[util.Point]struct{}

type Rock struct {
	Points RockSet
	Left, Right,
	Up, Down int
}

func (r Rock) Jet(rr rune) Rock {
	i, ok := JetMove[rr]
	if !ok {
		panic(fmt.Sprintf("ERROR ERROR ERROR: unknown rune %c\n", rr))
	}
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

func (r Rock) Print() {
	for y := r.Up + 3; y >= r.Down; y-- {
		fmt.Printf("|")
		for x := 0; x <= 6; x++ {
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

func (rg *RockGen) Next(a, b int) Rock {
	rock := rg.constructors[rg.index%len(rg.constructors)](a, b)
	rg.index++
	return rock
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

func part1() {
	windGen := ParseWindGen("input.txt")
	rockGen := NewRockGen()

	// create the tower with the floor added
	towerLeft, towerRight, numDrop := 0, 6, 2022
	floorPoints := make(RockSet)
	for x := towerLeft; x <= towerRight; x++ {
		floorPoints[util.NewPoint(x, 0)] = struct{}{}
	}
	tower := Rock{floorPoints, towerLeft, towerRight, 0, 0}

	for rockDrop := 0; rockDrop < numDrop; rockDrop++ {
		rock := rockGen.Next(2, tower.Up+4)
		//fmt.Printf("Rock drop: %v\n", rockDrop)
		//rock.Print()
		//tower.Print()
		// blow the wind and drop the rock until it stops falling
		for true {
			windRock := rock.Jet(windGen.Next())
			if windRock.Left >= towerLeft && windRock.Right <= towerRight && !tower.Collision(windRock) {
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
	//tower.Print()
}

func part2() {
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
