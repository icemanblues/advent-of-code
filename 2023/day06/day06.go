package main

import (
	"fmt"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "06"
	dayTitle = "Wait For It"
)

func parseRaw(filename string) ([]string, []string) {
	input, _ := util.ReadInput(filename)
	timeLine := strings.Split(input[0], ": ")
	times := strings.Fields(timeLine[1])

	distLine := strings.Split(input[1], ": ")
	dists := strings.Fields(distLine[1])
	return times, dists
}

func parseBoatRaces(filename string) ([]int, []int) {
	time, dist := parseRaw(filename)
	return util.SliceAtoI(time), util.SliceAtoI(dist)
}

func parseKerning(filename string) (int, int) {
	time, dist := parseRaw(filename)
	t, d := strings.Builder{}, strings.Builder{}
	for i, v := range time {
		t.WriteString(v)
		d.WriteString(dist[i])
	}
	return util.MustAtoi(t.String()), util.MustAtoi(d.String())
}

func dist(charge, limit int) int {
	return charge * (limit - charge)
}

func winningCharge(t, d int) (min, max int) {
	for c := 1; c < t; c++ {
		if dist(c, t) > d {
			min = c
			break
		}
	}

	for c := t - 1; c > 0; c-- {
		if dist(c, t) > d {
			max = c
			break
		}
	}
	return
}

func part1() {
	times, dists := parseBoatRaces("input.txt")
	product := 1
	for i, t := range times {
		d := dists[i]
		min, max := winningCharge(t, d)
		product *= (max - min + 1)
	}
	fmt.Printf("Part 1: %v\n", product)
}

func part2() {
	time, distance := parseKerning("input.txt")
	min, max := winningCharge(time, distance)
	fmt.Printf("Part 2: %v\n", (max - min + 1))
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
