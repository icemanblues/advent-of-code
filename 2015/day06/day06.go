package main

import (
	"fmt"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "06"
	dayTitle = "Probably a Fire Hazard"
)

type Point struct {
	x, y int
}

func NewPoint(s string) Point {
	xy := strings.Split(s, ",")
	return Point{util.MustAtoi(xy[0]), util.MustAtoi(xy[1])}
}

type Action int

const (
	on Action = iota
	off
	toggle
)

type Rect struct {
	top, bot Point
	action   Action
}

func (r Rect) contains(p Point) bool {
	return r.top.x <= p.x && r.top.y <= p.y &&
		r.bot.x >= p.x && r.bot.y >= p.y
}

func ReadRect(filename string) []Rect {
	var rects []Rect
	for _, line := range util.MustRead(filename) {
		parts := strings.Fields(line)
		action := toggle
		if parts[0] == "turn" && parts[1] == "on" {
			action = on
		} else if parts[0] == "turn" && parts[1] == "off" {
			action = off
		}
		bot := parts[len(parts)-1]
		top := parts[len(parts)-3]
		rects = append(rects, Rect{NewPoint(top), NewPoint(bot), action})
	}
	return rects
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	rects := ReadRect("input.txt")
	lit := 0
	lumens := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			p := Point{x, y}
			isOn := false
			lumen := 0
			for _, rect := range rects {
				if rect.contains(p) {
					switch rect.action {
					case on:
						isOn = true
						lumen++
					case off:
						isOn = false
						lumen--
						if lumen < 0 {
							lumen = 0
						}
					case toggle:
						isOn = !isOn
						lumen += 2
					}
				}
			}
			if isOn {
				lit++
			}
			lumens += lumen
		}
	}
	fmt.Printf("Part 1: %v\n", lit)
	fmt.Printf("Part 2: %v\n", lumens)
}
