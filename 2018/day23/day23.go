package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	X, Y, Z int
}

type Nanobot struct {
	point  Point
	radius int
}

func readInput(filename string) []Nanobot {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("something bad with file: %v\n", err)
		return nil
	}
	defer file.Close()

	var bots []Nanobot
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// pos=<40090364,79904452,-1333054>, r=90438706
		line := scanner.Text()
		posr := strings.Split(line, ">")

		points := strings.Split(posr[0], "<")
		values := strings.Split(points[1], ",")

		p := Point{}
		for i, s := range values {
			if i == 0 {
				p.X, err = strconv.Atoi(s)
				if err != nil {
					fmt.Printf("Unable to parse Point X: %v err: %v\n", s, err)
				}
			} else if i == 1 {
				p.Y, err = strconv.Atoi(s)
				if err != nil {
					fmt.Printf("Unable to parse Point Y: %v err: %v\n", s, err)
				}
			} else {
				p.Z, err = strconv.Atoi(s)
				if err != nil {
					fmt.Printf("Unable to parse Point Z: %v err: %v\n", s, err)
				}
			}
		}

		r := strings.Split(posr[1], "=")
		radius, err := strconv.Atoi(r[1])
		if err != nil {
			fmt.Printf("Unable to determine radius: %v err: %v\n", radius, err)
		}

		bots = append(bots, Nanobot{p, radius})

	}
	return bots
}

func main() {
	fmt.Println("Day 23: Experimental Emergency Teleportation")

	// part1("test.txt")
	// part1("input23.txt")
	part2("test2.txt")
	part2("input23.txt")
}

// Abs .
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Distance .
func Distance(a, b Point) int {
	return Abs(a.X-b.X) + Abs(a.Y-b.Y) + Abs(a.Z-b.Z)
}

// Manhattan distance
func Manhattan(a Point) int {
	return Distance(a, Zero)
}

// InRadius true if b is in the radius of a (not necessarily vice versa)
func InRadius(a, b Nanobot) bool {
	d := Distance(a.point, b.point)
	return d <= a.radius
}

func part1(fn string) {
	fmt.Println("Part 1")

	nanobots := readInput(fn)

	// determine the strongest signal bot
	strongBot := nanobots[0]
	strongRadius := strongBot.radius
	for _, n := range nanobots {
		if n.radius > strongRadius {
			strongRadius = n.radius
			strongBot = n
		}
	}

	fmt.Printf("strongest nanobot: %v\n", strongBot)

	count := 0
	for _, bot := range nanobots {
		if InRadius(strongBot, bot) {
			count++
		}
	}
	fmt.Printf("In range of the strongest: %d\n", count)
}

// Zero .
var Zero = Point{0, 0, 0}

func bruteForce(nanobots []Nanobot, minX, maxX, minY, maxY, minZ, maxZ int) Point {
	// minX, maxX, minY, maxY, minZ, maxZ := xyzRange(nanobots)

	total := (maxX - minX) * (maxY - minY) * (maxZ - minZ)

	// find all of the points with the best count of points
	iter := 0
	bestCount := 0
	var bestPoints []Point
	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			for k := minZ; k <= maxZ; k++ {
				if iter%10000 == 0 {
					fmt.Printf("iter %d out of %d\n", iter, total)
				}

				p := Point{i, j, k}

				count := 0
				for _, n := range nanobots {
					if Distance(p, n.point) <= n.radius {
						count++
					}
				}

				if count > bestCount {
					bestCount = count
					bestPoints = []Point{p}
				} else if count == bestCount {
					bestPoints = append(bestPoints, p)
				}

				iter++
			}
		}
	}

	// from the best points, find the one with the shortest distance
	zero := Point{0, 0, 0}
	bp := bestPoints[0]
	bd := Distance(zero, bp)
	for _, p := range bestPoints {
		if d := Distance(zero, p); d < bd {
			bd = d
			bp = p
		}
	}

	fmt.Printf("Choose point: %v\n", bp)
	return bp
}

func xyzRange(nanobots []Nanobot) (minX, maxX, minY, maxY, minZ, maxZ int) {
	// find the min/max point
	minX, maxX = math.MaxInt32, -1
	minY, maxY = math.MaxInt32, -1
	minZ, maxZ = math.MaxInt32, -1
	for _, n := range nanobots {
		p := n.point
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}

		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}

		if p.Z < minZ {
			minZ = p.Z
		}
		if p.Z > maxZ {
			maxZ = p.Z
		}
	}

	return
}

func avg(points []Point) Point {
	l := len(points)
	sumX, sumY, sumZ := 0, 0, 0
	for _, p := range points {
		sumX += p.X
		sumY += p.Y
		sumZ += p.Z
	}

	avg := Point{sumX / l, sumY / l, sumZ / l}
	return avg
}

func avgBot(bots []Nanobot) Point {
	points := make([]Point, len(bots), len(bots))
	for i, n := range bots {
		points[i] = n.point
	}

	return avg(points)
}

func score(nanobots []Nanobot, point Point) (c, m int, in, out []Nanobot) {
	m = Manhattan(point)

	c = 0
	for _, n := range nanobots {
		if d := Distance(point, n.point); d <= n.radius {
			c++
			in = append(in, n)
		} else {
			out = append(out, n)
		}
	}
	return
}

func delta(a, b Point) (d, e, f int) {
	if b.X > a.X {
		d = 1
	} else {
		d = -1
	}

	if b.Y > a.Y {
		e = 1
	} else {
		e = -1
	}

	if b.Z > a.Z {
		f = 1
	} else {
		f = -1
	}

	return d, e, f
}

// returns a new point 1 step closer to b (from a)
func inc(a, b Point) Point {
	xdiff := Abs(b.X - a.X)
	ydiff := Abs(b.Y - a.Y)
	zdiff := Abs(b.Z - a.Z)

	dx, dy, dz := delta(a, b)

	if xdiff > ydiff && xdiff > zdiff {
		return Point{a.X + dx, a.Y, a.Z}
	} else if ydiff > zdiff {
		return Point{a.X, a.Y + dy, a.Z}
	}
	return Point{a.X, a.Y, a.Z + dz}
}

func cluster(start, end Point, nanobots []Nanobot) (Point, int, int) {
	// walk from starting point to the average, moving by largest distance, one step at a time
	// want to maximize count and minimize length
	distance := Distance(start, end)
	sc, sm, _, _ := score(nanobots, start)
	currPoint := start
	bestPoint := start
	bestCount := sc
	bestLen := sm
	iter := 0
	for {
		// !!! increment the point
		// move closer to average
		currPoint = inc(currPoint, end)

		// score the increment
		c, m, _, _ := score(nanobots, currPoint)
		// fmt.Printf("iter %d: curr %v %d %d\n", iter, currPoint, c, m)
		if c == bestCount && m < bestLen {
			bestPoint = currPoint
			bestLen = m
			bestCount = c
		}
		if c > bestCount {
			bestPoint = currPoint
			bestCount = c
			bestLen = m
		}

		if iter%100000 == 0 {
			fmt.Printf("iter %d: curr %v best point, count, length %v %d %d out of total %d\n", iter, currPoint, bestPoint, bestCount, bestLen, distance)
		}

		if currPoint == end {
			break
		}

		iter++
	}

	fmt.Printf("iter %d: best point, count, length: %v, %d, %d\n", iter, bestPoint, bestCount, bestLen)
	return bestPoint, bestCount, bestLen
}

// Interval .
type Interval struct {
	lo, hi int
}

// Accessor .
type Accessor func(n Nanobot) int

// In .
func (it Interval) In(x int) bool {
	return it.lo <= x && it.hi >= x
}

func interval(nanobots []Nanobot, accessor Accessor) []Interval {
	ints := make([]Interval, len(nanobots), len(nanobots))
	for i, n := range nanobots {
		x := accessor(n)
		ints[i] = Interval{x - n.radius, x + n.radius}
	}
	return ints
}

func minMaxInts(ints []Interval) (lo, hi int) {
	lo, hi = ints[0].lo, ints[0].hi
	for _, it := range ints {
		if l := it.lo; l < lo {
			lo = l
		}
		if h := it.hi; h > hi {
			hi = h
		}
	}
	return
}

func search(ints []Interval, acc Accessor, min, max int) (int, int) {
	iCount := 0
	iBest := 0

	for i := min; i <= max; i++ {
		c := 0
		for _, n := range ints {
			if n.In(i) {
				c++
			}
		}

		if iCount < c {
			iCount = c
			iBest = i
		}
	}

	return iBest, iCount
}

func intervalSearch(ints []Interval) (int, int, int) {
	n := len(ints)
	los := make([]int, n, n)
	his := make([]int, n, n)
	for i, it := range ints {
		los[i] = it.lo
		his[i] = it.hi
	}

	// sort the upper and lower bounds
	sort.Ints(los)
	sort.Ints(his)

	count := 1
	maxCount := 1
	time := los[0]
	timeOut := his[0]

	i, j := 0, 0
	for i < n && j < n {
		if los[i] <= his[j] {
			count++

			//reset the max value here
			if count > maxCount {
				maxCount = count
				time = los[i]
			}

			i++
		} else {
			if count == maxCount {
				timeOut = his[j]
				// fmt.Printf("count: %d Good window: %d - %d total:%d\n", count, time, timeOut, timeOut-time)
			}

			count--
			j++
		}
	}

	return time, timeOut, maxCount
}

// 142831677 too high
// 141205672 too high
// 143324916 {44440530 53965408 44918978}, 840,
// 107137375 too low
func part2(fn string) {
	fmt.Println("Part 2")

	nanobots := readInput(fn)
	fmt.Printf("nanobot count: %d\n", len(nanobots))

	getX := func(n Nanobot) int { return n.point.X }
	getY := func(n Nanobot) int { return n.point.Y }
	getZ := func(n Nanobot) int { return n.point.Z }
	xInts := interval(nanobots, getX)
	yInts := interval(nanobots, getY)
	zInts := interval(nanobots, getZ)

	// find the highest and lowest x value

	// xmin, xmax := minMaxInts(xInts)
	// ymin, ymax := minMaxInts(yInts)
	// zmin, zmax := minMaxInts(zInts)

	// search for best ranges
	// fmt.Println("-- X Axis --")
	xBest, xHigh, xCount := intervalSearch(xInts)
	// fmt.Println("-- Y Axis --")
	yBest, yHigh, yCount := intervalSearch(yInts)
	// fmt.Println("-- Z Axis --")
	zBest, zHigh, zCount := intervalSearch(zInts)

	fmt.Printf("x-axis: best: %d high: %d count: %d\n", xBest, xHigh, xCount)
	fmt.Printf("y-axis: best: %d high: %d count: %d\n", yBest, yHigh, yCount)
	fmt.Printf("z-axis: best: %d high: %d count: %d\n", zBest, zHigh, zCount)

	seed := Point{xBest, yBest, zBest}
	seedCount, seedLen, _, _ := score(nanobots, seed)
	fmt.Printf("seed %v count %d length %d\n", seed, seedCount, seedLen)

	best := bruteForce(nanobots, xBest, xHigh, yBest, yHigh, zBest, zHigh)
	answer := Manhattan(best)
	fmt.Printf("Point is best %v with distance %d\n", best, answer)

}
