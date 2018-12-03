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

	// f1 := NewFabric("#1 @ 1,3: 4x4")
	// f2 := NewFabric("#2 @ 3,1: 4x4")
	// f3 := NewFabric("#3 @ 5,5: 2x2")

	part1()
}

type Fabric struct {
	ID int
	X  int
	Y  int
	W  int
	H  int
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

	counter := [1000][1000]int{}
	for _, k := range f {
		for i := 0; i < k.W; i++ {
			for j := 0; j < k.H; j++ {
				counter[k.X+i][k.Y+j]++
			}
		}
	}

	sum := 0
	for i := range counter {
		for j := range counter[i] {
			if counter[i][j] > 1 {
				sum++
			}
		}
	}
	fmt.Printf("sum of counts for overlaps: %v\n", sum)

	claims := [1000][1000]map[int]bool{}
	for _, k := range f {
		for i := 0; i < k.W; i++ {
			for j := 0; j < k.H; j++ {
				mapm := claims[k.X+i][k.Y+j]
				if mapm == nil {
					mapm = make(map[int]bool)
					claims[k.X+i][k.Y+j] = mapm
				}
				claims[k.X+i][k.Y+j][k.ID] = true
			}
		}
	}

	overlapID := make(map[int]bool)
	for i := range claims {
		for j := range claims[i] {
			if len(claims[i][j]) > 1 {
				for k, _ := range claims[i][j] {
					overlapID[k] = true
				}
			}
		}
	}

	fmt.Printf("num fabrics %v, num overlapped %v\n", len(f), len(overlapID))

	for _, k := range f {
		_, ok := overlapID[k.ID]
		if !ok {
			fmt.Printf("does not overlap with anyone: %v\n", k.ID)
			break
		}
	}
}
