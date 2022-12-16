package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "15"
	dayTitle = "Beacon Exclusion Zone"
)

type SensorBeacon struct {
	Sensor, Beacon util.Point
	dist           int
}

func parse(filename string) []SensorBeacon {
	input, _ := util.ReadInput(filename)
	sb := make([]SensorBeacon, 0, len(input))
	for _, line := range input {
		tokens := strings.Split(line, ": ")
		left := strings.Split(tokens[0], ", ")
		sensorXCoord := strings.Split(left[0], "=")[1]
		sensorYCoord := strings.Split(left[1], "=")[1]
		sX, _ := strconv.Atoi(sensorXCoord)
		sY, _ := strconv.Atoi(sensorYCoord)
		sensor := util.Point{sX, sY}

		right := strings.Split(tokens[1], ", ")
		beaconXCoord := strings.Split(right[0], "=")[1]
		beaconYCoord := strings.Split(right[1], "=")[1]
		bX, _ := strconv.Atoi(beaconXCoord)
		bY, _ := strconv.Atoi(beaconYCoord)
		beacon := util.Point{bX, bY}

		dist := util.ManhattanDist(sensor, beacon)
		sb = append(sb, SensorBeacon{sensor, beacon, dist})
	}
	return sb
}

type Boundary struct {
	Xmin, Xmax, Ymin, Ymax int
	Dist                   int
}

func GetBounds(input []SensorBeacon) (map[util.Point]rune, Boundary) {
	grid := make(map[util.Point]rune)
	grid[input[0].Sensor] = 'S'
	grid[input[0].Beacon] = 'B'
	xmin, xmax := input[0].Sensor.X, input[0].Sensor.X
	ymin, ymax := input[0].Sensor.Y, input[0].Sensor.Y
	maxDist := input[0].dist
	for i := 1; i < len(input); i++ {
		item := input[i]
		grid[item.Sensor] = 'S'
		grid[item.Beacon] = 'B'

		if item.Sensor.X < xmin {
			xmin = item.Sensor.X
		}
		if item.Sensor.X > xmax {
			xmax = item.Sensor.X
		}
		if item.Sensor.Y < ymin {
			ymin = item.Sensor.Y
		}
		if item.Sensor.Y > ymax {
			ymax = item.Sensor.Y
		}
		if item.dist > maxDist {
			maxDist = item.dist
		}
	}
	return grid, Boundary{xmin, xmax, ymin, ymax, maxDist}
}

func noBeaconCount(input []SensorBeacon, row int) int {
	grid, b := GetBounds(input)
	count := 0
	for x := b.Xmin - b.Dist; x <= b.Xmax+b.Dist; x++ {
		p := util.Point{x, row}
		if _, ok := grid[p]; ok {
			continue
		}
		for _, item := range input {
			if util.ManhattanDist(item.Sensor, p) <= item.dist {
				count++
				break
			}
		}
	}
	return count
}

func part1() {
	fmt.Printf("Part 1: %v\n", noBeaconCount(parse("input.txt"), 2000000))
}

func part2() {
	input := parse("input.txt")
	grid, bounds := GetBounds(input)
	fmt.Println(len(grid))
	fmt.Println(bounds)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
