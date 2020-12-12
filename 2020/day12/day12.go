package main

import (
	"fmt"
	"strconv"

	"github.com/icemanblues/advent-of-code/2020/util"
)

type Coord struct {
	x, y int
	dir  rune
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sail(directions [][]rune) int {
	pos := Coord{0, 0, 'E'}
	for _, direction := range directions {
		num, _ := strconv.Atoi(string(direction[1:]))
		pos = mov(direction[0], num, pos)
	}
	return abs(pos.x) + abs(pos.y)
}

func rotateLeft(face rune, num int) rune {
	for num != 0 {
		switch face {
		case 'N':
			face = 'W'
		case 'W':
			face = 'S'
		case 'S':
			face = 'E'
		case 'E':
			face = 'N'
		default:
			panic("unknown face: " + string(face))
		}
		num -= 90
	}
	return face
}

func mov(r rune, num int, c Coord) Coord {
	switch r {
	case 'N':
		return Coord{c.x, c.y - num, c.dir}
	case 'S':
		return Coord{c.x, c.y + num, c.dir}
	case 'E':
		return Coord{c.x + num, c.y, c.dir}
	case 'W':
		return Coord{c.x - num, c.y, c.dir}
	case 'L':
		return Coord{c.x, c.y, rotateLeft(c.dir, num)}
	case 'R':
		return Coord{c.x, c.y, rotateLeft(c.dir, 360-num)}
	case 'F':
		return mov(c.dir, num, c)
	default:
		panic("unknown instruction: " + string(r))
	}
}

func rotateDegreesLeft(c Coord, deg int) Coord {
	if deg == 90 {
		return Coord{c.y, -c.x, c.dir}
	} else if deg == 180 {
		return Coord{-c.x, -c.y, c.dir}
	} else if deg == 270 {
		return Coord{-c.y, c.x, c.dir}
	}
	panic("unknown degree:")
}

func waypoint(directions [][]rune) int {
	ship := Coord{0, 0, 'E'}
	wp := Coord{10, -1, 'W'}
	for _, direction := range directions {
		inst := direction[0]
		num, _ := strconv.Atoi(string(direction[1:]))
		switch inst {
		case 'N':
			wp = mov(inst, num, wp)
		case 'S':
			wp = mov(inst, num, wp)
		case 'E':
			wp = mov(inst, num, wp)
		case 'W':
			wp = mov(inst, num, wp)
		case 'L':
			wp = rotateDegreesLeft(wp, num)
		case 'R':
			wp = rotateDegreesLeft(wp, 360-num)
		case 'F':
			ship = Coord{ship.x + num*wp.x, ship.y + num*wp.y, ship.dir}
		default:
			panic("unknown instruction: " + string(inst))
		}
	}
	return abs(ship.x) + abs(ship.y)
}

func part1() {
	directions, _ := util.ReadRuneput("input.txt")
	fmt.Printf("Part 1: %v\n", sail(directions))
}

func part2() {
	directions, _ := util.ReadRuneput("input.txt")
	fmt.Printf("Part 2: %v\n", waypoint(directions))
}

func main() {
	fmt.Println("Day 12: Rain Risk")
	part1()
	part2()
}
