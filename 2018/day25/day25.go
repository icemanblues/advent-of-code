package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Point .
type Point struct {
	X, Y, Z, T int
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
	return Abs(a.X-b.X) + Abs(a.Y-b.Y) + Abs(a.Z-b.Z) + Abs(a.T-b.T)
}

// InConstellation .
func InConstellation(a, b Point) bool {
	return Distance(a, b) <= 3
}

func readInput(filename string) []Point {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("something bad with file: %v\n", err)
		return nil
	}
	defer file.Close()

	var points []Point
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, ",")

		x, err := strconv.Atoi(nums[0])
		if err != nil {
			fmt.Printf("Unable to parse X %v on line: %v\n", nums[0], line)
		}

		y, err := strconv.Atoi(nums[1])
		if err != nil {
			fmt.Printf("Unable to parse Y %v on line: %v\n", nums[1], line)
		}
		z, err := strconv.Atoi(nums[2])
		if err != nil {
			fmt.Printf("Unable to parse Z %v on line: %v\n", nums[2], line)
		}
		t, err := strconv.Atoi(nums[3])
		if err != nil {
			fmt.Printf("Unable to parse T %v on line: %v\n", nums[3], line)
		}

		points = append(points, Point{x, y, z, t})
	}
	return points
}

func main() {
	fmt.Println("Day 25: Four-Dimensional Adventure")

	// part1("test1-4.txt")
	// part1("test1-3.txt")
	// part1("test1-8.txt")
	part1("input25.txt")
	// part2()
}

func part1(fn string) {
	fmt.Println("Part 1")

	points := readInput(fn)

	// first pass, build a slice of sets
	var constellations []map[Point]struct{}
	for i := 0; i < len(points)-1; i++ {
		set := make(map[Point]struct{})
		set[points[i]] = struct{}{}
		for j := i + 1; j < len(points); j++ {
			if InConstellation(points[i], points[j]) {
				set[points[j]] = struct{}{}
			}
		}
		constellations = append(constellations, set)
	}

	fmt.Printf("After first pass, we have %v constellations\n", len(constellations))

	// continue to reduce the slice until there is no change, then stop
	prev := len(constellations)

	// reduce the constellations together

	for {
		var deleteMe []int
		deleteSet := make(map[int]struct{})

		// merge set j into set i, and then skipping j and deleting it before iterating
		for i := 0; i < len(constellations)-1; i++ {
			if _, ok := deleteSet[i]; ok {
				continue
			}
			for j := i + 1; j < len(constellations); j++ {
				if _, ok := deleteSet[j]; ok {
					continue
				}

			intersection:
				for pi := range constellations[i] {
					for pj := range constellations[j] {
						if InConstellation(pj, pi) {
							// add to be deleted
							deleteMe = append(deleteMe, j)
							deleteSet[j] = struct{}{}
							// now combine into i
							for k, v := range constellations[j] {
								constellations[i][k] = v
							}
							break intersection
						}
					}
				}
			}
		}

		// now delete all of indices marked for deletion. largest index first (back to front)
		sort.Ints(deleteMe)
		// fmt.Printf("Want to delete these indices: %v\n", deleteMe)
		for i := len(deleteMe) - 1; i >= 0; i-- {
			d := deleteMe[i]
			// fmt.Printf("deleting index %d from constellations len: %d\n", d, len(constellations))
			constellations = append(constellations[:d], constellations[d+1:]...)
		}

		// no change, so stop processing
		if len(constellations) == prev {
			break
		}
		prev = len(constellations)
	}

	fmt.Printf("We have this many constellations: %d\n", len(constellations))
}

func part2() {
	fmt.Println("Part 2")
}
