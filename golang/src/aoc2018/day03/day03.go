package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 3: No Matter How You Slice It")

	f1 := NewFabric("#1 @ 1,3: 4x4")
	f2 := NewFabric("#2 @ 3,1: 4x4")
	f3 := NewFabric("#3 @ 5,5: 2x2")
	fmt.Printf("f1 overlap f2?  true : %v : overlap %v\n", IsOverlap(f1, f2), Overlap(f1, f2))
	fmt.Printf("f2 overlap f3? false : %v\n", IsOverlap(f2, f3))
	fmt.Printf("f1 overlap f3? false : %v\n", IsOverlap(f1, f3))

	part1()
	// part2()
}

func part2() {
	fmt.Println("Day 03 Part 2")
	// file, _ := os.Open("input03.txt")
	// defer file.Close()

	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// }
}

type Fabric struct {
	ID int
	X  int
	Y  int
	W  int
	H  int
}

func (f *Fabric) Xend() int {
	return f.X + f.W - 1
}

func (f *Fabric) Yend() int {
	return f.Y + f.H - 1
}

func NewFabric(s string) Fabric {
	// #1 @ 1,3: 4x4
	parts := strings.Split(s, " ")
	id, _ := strconv.Atoi(strings.Split(parts[0], "#")[1])

	xy := strings.Split(strings.TrimSuffix(parts[2], ":"), ",")
	x, _ := strconv.Atoi(xy[0])
	y, _ := strconv.Atoi(xy[1])

	wh := strings.Split(parts[3], "x")
	w, _ := strconv.Atoi(wh[0])
	h, _ := strconv.Atoi(wh[1])

	return Fabric{
		ID: id,
		X:  x,
		Y:  y,
		W:  w,
		H:  h,
	}
}

func part1() {
	fmt.Println("Day 03 Part 1")

	file, _ := os.Open("input03.txt")
	defer file.Close()

	var f []Fabric

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		f = append(f, NewFabric(line))
	}

	// track the indices that are bad
	bad := make(map[int]bool)
	points := make(map[Point]bool)
	var sqin int = 0

	for i := 0; i < len(f)-1; i++ {
		for j := i + 1; j < len(f); j++ {
			if IsOverlap(f[i], f[j]) {
				bad[i] = true
				bad[j] = true

				p := Overlap(f[i], f[j])
				sqin += len(p)
				for _, v := range p {
					points[v] = true
				}
			}
		}
	}

	fmt.Printf("number of fabrics: %v\n", len(f))
	fmt.Printf("number of overlaps: %v\n", len(bad))
	fmt.Printf("square inches overlapped %v\n", len(points))
	fmt.Printf("sqin: %v\n", sqin)
}

type Point struct {
	X int
	Y int
}

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func Overlap(f1, f2 Fabric) []Point {
	var p []Point

	xmin := Max(f2.X, f1.X)
	xmax := Min(f2.Xend(), f1.Xend())

	ymin := Max(f2.Y, f1.Y)
	ymax := Min(f2.Yend(), f1.Yend())

	for i := xmin; i <= xmax; i++ {
		for j := ymin; j <= ymax; j++ {
			p = append(p, Point{
				X: i,
				Y: j,
			})
		}
	}

	return p
}

func IsOverlap(f1, f2 Fabric) bool {

	xmin := f2.X
	xmax := f2.Xend()
	xOverlap := (f1.X <= xmin && f1.Xend() >= xmin) || (f1.X <= xmax && f1.Xend() >= xmax)

	ymin := f2.Y
	ymax := f2.Yend()
	yOverlap := (f1.Y <= ymin && f1.Yend() >= ymin) || (f1.Y <= ymax && f1.Yend() >= ymax)

	return xOverlap && yOverlap

}
