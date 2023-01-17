package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "18"
	dayTitle = "Boiling Boulders"
)

func border(a, b util.Point3D) bool {
	return (a.X == b.X && a.Y == b.Y && util.Abs(a.Z-b.Z) == 1) ||
		(a.X == b.X && a.Z == b.Z && util.Abs(a.Y-b.Y) == 1) ||
		(a.Z == b.Z && a.Y == b.Y && util.Abs(a.X-b.X) == 1)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	cubes, _ := util.Read3D("input.txt", ",")

	sidesShared := make(map[util.Point3D]int)
	for i := 0; i < len(cubes)-1; i++ {
		for j := i + 1; j < len(cubes); j++ {
			a, b := cubes[i], cubes[j]
			if border(a, b) {
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

	minX, maxX := cubes[0].X, cubes[0].X
	minY, maxY := cubes[0].Y, cubes[0].Y
	minZ, maxZ := cubes[0].Z, cubes[0].Z
	cubeSet := make(map[util.Point3D]struct{})
	for _, c := range cubes {
		cubeSet[c] = struct{}{}
		if c.X > maxX {
			maxX = c.X
		}
		if c.X < minX {
			minX = c.X
		}
		if c.Y > maxY {
			maxY = c.Y
		}
		if c.Y < minY {
			minY = c.Y
		}
		if c.Z > maxZ {
			maxZ = c.Z
		}
		if c.Z < minZ {
			minZ = c.Z
		}
	}

	// all possible air pockets, surrounding our known cubes
	airCube := make(map[util.Point3D]bool)
	for x := minX - 1; x <= maxX+1; x++ {
		for y := minY - 1; y <= maxY+1; y++ {
			for z := minZ - 1; z <= maxZ+1; z++ {
				a := util.NewPoint3D(x, y, z)
				if _, ok := cubeSet[a]; !ok {
					airCube[a] = false
				}
			}
		}
	}

	// need to filter out the air pockets are reachable from the exterior
	// start at the minimum, then check to see if the neighbors are in the set of air pockets
	// BFS to see what else is outside air
	q := []util.Point3D{util.NewPoint3D(minX-1, minY-1, minZ-1)}
	for len(q) != 0 {
		c := q[0]
		q = q[1:]

		visited, ok := airCube[c]
		if visited || !ok { // already visited or not viable air
			continue
		}
		airCube[c] = true
		next := []util.Point3D{
			util.NewPoint3D(c.X-1, c.Y, c.Z),
			util.NewPoint3D(c.X+1, c.Y, c.Z),
			util.NewPoint3D(c.X, c.Y-1, c.Z),
			util.NewPoint3D(c.X, c.Y+1, c.Z),
			util.NewPoint3D(c.X, c.Y, c.Z-1),
			util.NewPoint3D(c.X, c.Y, c.Z+1),
		}
		q = append(q, next...)
	}

	airShared := make(map[util.Point3D]int)
	for a, b := range airCube {
		if b {
			continue
		}
		for _, c := range cubes {
			if border(a, c) {
				airShared[a]++
			}
		}
	}
	airSum := 0
	for _, v := range airShared {
		airSum += v
	}

	fmt.Printf("Part 2: %v\n", surfaceArea-airSum)
}
