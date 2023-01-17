package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "18"
	dayTitle = "Boiling Boulders"
)

func part1() {
	cubes, _ := util.Read3D("input.txt", ",")

	sidesShared := make(map[util.Point3D]int)
	for i := 0; i < len(cubes)-1; i++ {
		for j := i + 1; j < len(cubes); j++ {
			a, b := cubes[i], cubes[j]
			// check z
			if a.X == b.X && a.Y == b.Y && util.Abs(a.Z-b.Z) == 1 {
				sidesShared[a]++
				sidesShared[b]++
			}
			// check y
			if a.X == b.X && a.Z == b.Z && util.Abs(a.Y-b.Y) == 1 {
				sidesShared[a]++
				sidesShared[b]++
			}
			// check x
			if a.Z == b.Z && a.Y == b.Y && util.Abs(a.X-b.X) == 1 {
				sidesShared[a]++
				sidesShared[b]++
			}
		}
	}

	sum := 0
	for _, v := range sidesShared {
		sum += v
	}
	surfaceArea := 6*len(cubes) - sum
	fmt.Printf("Part 1: %v\n", surfaceArea)
}

func part2() {
	fmt.Println("Part 2")
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
