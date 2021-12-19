package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "17"
	dayTitle = "Trick Shot"
)

type TargetArea struct {
	MinX, MaxX, MinY, MaxY int
}

var TestArea = TargetArea{20, 30, -10, -5}

var InputArea = TargetArea{257, 286, -101, -57}

func trench(pos util.Point, area TargetArea) bool {
	return pos.X >= area.MinX && pos.X <= area.MaxX && pos.Y >= area.MinY && pos.Y <= area.MaxY
}

func tick(pos, vel util.Point) (util.Point, util.Point) {
	pos.X = pos.X + vel.X
	pos.Y = pos.Y + vel.Y

	if vel.X > 0 {
		vel.X = vel.X - 1
	} else if vel.X < 0 {
		vel.X = vel.X + 1
	}
	vel.Y = vel.Y - 1

	return pos, vel
}

func launch(vel util.Point, area TargetArea) (int, bool) {
	pos := util.Point{X: 0, Y: 0}
	h := 0
	for true {
		if pos.Y > h {
			h = pos.Y
		}
		if trench(pos, area) {
			return h, true
		}

		// increment to iterate
		pos, vel = tick(pos, vel)

		// check to terminate
		// Y is lower than the min Y
		// X isn't in range and vel X is zero
		if pos.Y < area.MinY {
			break
		}
		if vel.X == 0 && (pos.X > area.MaxX || pos.X < area.MinY) {
			break
		}
	}
	return h, false
}

func search(area TargetArea) (int, int) {
	min := util.Min(area.MinX, area.MinY)
	max := util.Max(area.MaxX, area.MaxY)
	best := 0
	count := 0
	for x := min; x <= max; x++ {
		for y := min; y <= max; y++ {
			vel := util.Point{X: x, Y: y}
			h, ok := launch(vel, area)
			if best == 0 {
				best = h
			}
			if ok && h > best {
				best = h
			}
			if ok {
				count++
			}
		}
	}
	return best, count
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	height, count := search(InputArea)
	fmt.Printf("Part 1: %v\n", height)
	fmt.Printf("Part 2: %v\n", count)
}
