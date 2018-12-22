package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("something bad with file: %v\n", err)
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

type Point struct{ X, Y int }

type Input struct {
	depth  int
	target Point
}

func main() {
	fmt.Println("Day 22: Mode Maze")

	test := Input{depth: 510, target: Point{10, 10}}
	part1(test)
	input := Input{depth: 4848, target: Point{15, 700}}
	part1(input)

	// 983 too high
	// part2()
}

var Mouth = Point{0, 0}

func computeGeologic(p, target Point, depth int, geoCache, eroCache map[Point]int) int {
	// check cache first
	if s, ok := geoCache[p]; ok {
		return s
	}

	if p == Mouth {
		geoCache[p] = 0
		return geoCache[p]
	}
	if p == target {
		geoCache[p] = 0
		return geoCache[p]
	}
	if p.Y == 0 {
		geoCache[p] = p.X * 16807
		return geoCache[p]
	}
	if p.X == 0 {
		geoCache[p] = p.Y * 48271
		return geoCache[p]
	}

	// should call compute erosion here
	geoCache[p] = computeErosion(Point{p.X - 1, p.Y}, target, depth, eroCache, geoCache) * computeErosion(Point{p.X, p.Y - 1}, target, depth, eroCache, geoCache)
	return geoCache[p]
}

func computeErosion(p, target Point, depth int, eroCache, geoCache map[Point]int) int {
	if s, ok := eroCache[p]; ok {
		return s
	}

	geo := computeGeologic(p, target, depth, geoCache, eroCache)
	eroCache[p] = (geo + depth) % 20183

	return eroCache[p]
}

func computeRegion(p, target Point, depth int, regCache, eroCache, geoCache map[Point]int) int {
	if s, ok := regCache[p]; ok {
		return s
	}

	e := computeErosion(p, target, depth, eroCache, geoCache)
	regCache[p] = e % 3
	return regCache[p]

}

func computeAll(p, target Point, depth int, geoCache, eroCache, regionCache map[Point]int) (int, int, int) {
	geo := computeGeologic(p, target, depth, geoCache, eroCache)
	ero := computeErosion(p, target, depth, eroCache, geoCache)
	reg := computeRegion(p, target, depth, regionCache, eroCache, geoCache)

	return geo, ero, reg
}

// prolly will want to memoize this too
func computeRisk(target Point, regCache map[Point]int) int {
	sum := 0
	for i := 0; i <= target.X; i++ {
		for j := 0; j <= target.Y; j++ {
			sum += regCache[Point{i, j}]
		}
	}
	return sum
}

func adj(p Point) []Point {
	var points []Point

	if p.X-1 >= 0 {
		points = append(points, Point{p.X - 1, p.Y})
	}
	if p.Y-1 >= 0 {
		points = append(points, Point{p.X, p.Y - 1})
	}

	points = append(points, Point{p.X + 1, p.Y})
	points = append(points, Point{p.X, p.Y + 1})

	return points
}

func part1(input Input) {
	fmt.Println("Part 1")

	geoCache := make(map[Point]int)
	erosionCache := make(map[Point]int)
	regionCache := make(map[Point]int)

	// compute all values for the top row and top column
	// this will seed the
	for i := 0; i <= input.target.X; i++ {
		computeAll(Point{i, 0}, input.target, input.depth, geoCache, erosionCache, regionCache)
	}
	for j := 0; j <= input.target.Y; j++ {
		computeAll(Point{0, j}, input.target, input.depth, geoCache, erosionCache, regionCache)
	}
	// now do it for all points up to the target
	for i := 1; i <= input.target.X; i++ {
		for j := 1; j <= input.target.Y; j++ {
			computeAll(Point{i, j}, input.target, input.depth, geoCache, erosionCache, regionCache)
		}
	}

	sum := computeRisk(input.target, regionCache)
	fmt.Printf("Sum to %v is %d\n", input.target, sum)

	// region types
	rocky := 0  // torch, climbing
	wet := 1    // climbing, neither
	narrow := 2 // torch, neither
	// tools
	torch := 1
	climbing := 1
	neither := 2

	delay := 7

	toolBadRegion := map[int]int{
		torch:    wet,    // 0, 2
		climbing: narrow, // 0 1
		neither:  rocky,  // 1 2
	}

	// BFS from start to target, and back
	timer := -1
	start := Node{Mouth, 0, torch, 0, 0}
	queue := []Node{start}
	visited := make(map[Point]int)
	for len(queue) != 0 {
		n := queue[0]
		queue = queue[1:]

		// found the target
		// must equip the torch when target is reached
		if n.p == input.target {
			goal := n.time
			if n.tool != torch {
				// swap to torch
				goal += delay
			}

			if timer == -1 {
				timer = goal
			}
			if goal < timer {
				timer = goal
			}

			fmt.Printf("Target Arrived. time %d, Best: %d Node %v\n", goal, timer, n)
			continue
		}

		// if you have taken longer than our best time, you need to stop
		if timer != -1 && n.time > timer {
			continue
		}

		// have I been here before with better time
		if t, ok := visited[n.p]; ok && t <= n.time {
			continue
		}
		visited[n.p] = n.time

		computeAll(n.p, input.target, input.depth, geoCache, erosionCache, regionCache)
		curr := regionCache[n.p]
		t := n.tool

		// figure out the other tool to use for this region
		for tool, badReg := range toolBadRegion {
			if tool == n.tool {
				continue
			}
			if curr == badReg {
				continue
			}
			t = tool
		}

		// keep advancing towards the goal
		// get adj points and check their
		adjs := adj(n.p)
		for _, a := range adjs {
			r := computeRegion(a, input.target, input.depth, regionCache, erosionCache, geoCache)
			// enqueue with this tool
			if toolBadRegion[n.tool] != r {
				queue = append(queue, Node{a, n.time + 1, n.tool, n.steps + 1, n.swaps})
			}
			// enqueue with the other tool (plus swap)
			if toolBadRegion[t] != r {
				queue = append(queue, Node{a, n.time + 1 + delay, t, n.steps + 1, n.swaps + 1})
			}
		}
	}

	fmt.Printf("fewest minutes: %d\n", timer)
}

type Node struct {
	p     Point
	time  int
	tool  int
	steps int
	swaps int
}
