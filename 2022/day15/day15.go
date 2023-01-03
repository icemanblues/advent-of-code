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
		sensor := util.NewPoint(sX, sY)

		right := strings.Split(tokens[1], ", ")
		beaconXCoord := strings.Split(right[0], "=")[1]
		beaconYCoord := strings.Split(right[1], "=")[1]
		bX, _ := strconv.Atoi(beaconXCoord)
		bY, _ := strconv.Atoi(beaconYCoord)
		beacon := util.NewPoint(bX, bY)

		dist := util.ManhattanDist(sensor, beacon)
		sb = append(sb, SensorBeacon{sensor, beacon, dist})
	}
	return sb
}

func RangeRow(sensor util.Point, dist, row int) ([2]int, bool) {
	delta := dist - util.Abs(sensor.Y-row)
	if delta <= 0 {
		return [2]int{0, 0}, false
	}
	x1, x2 := sensor.X-delta, sensor.X+delta
	return [2]int{x1, x2}, true
}

func RangeReduce(list [][2]int) [][2]int {
	i, j := 0, 1
	l := len(list)
	for true {
		if j >= len(list) {
			i++
			j = i + 1
		}
		if i >= len(list)-1 {
			break
		}

		seg1 := list[i]
		seg2 := list[j]
		if seg1[0] <= seg2[0] && seg1[1] >= seg2[1] {
			// seg 1 encompasses seg2
			list = append(list[:j], list[j+1:]...)
		} else if seg2[0] <= seg1[0] && seg2[1] >= seg1[1] {
			// seg 2 encompasses seg1
			list = append(list[:i], list[i+1:]...)
		} else if seg1[0] >= seg2[0] && seg1[0] <= seg2[1] {
			// seg1 overlaps seg2
			m := util.Max(seg1[1], seg2[1])
			list = append(list[:j], list[j+1:]...)
			list[i] = [2]int{seg2[0], m}
		} else if seg2[0] >= seg1[0] && seg2[0] <= seg1[1] {
			// seg2 overlaps seg1
			m := util.Max(seg1[1], seg2[1])
			list = append(list[:j], list[j+1:]...)
			list[i] = [2]int{seg1[0], m}
		} else {
			// nothing
			j++
			continue
		}

		if l == len(list) { // no change
			break
		}
		l = len(list)
	}
	return list
}

func Coverage(input []SensorBeacon, row int) [][2]int {
	list := make([][2]int, 0, len(input))
	for _, item := range input {
		if r, ok := RangeRow(item.Sensor, item.dist, row); ok {
			list = append(list, r)
		}
	}

	// now need to reduce this list of arrays
	// TODO: call RangeReduce twice to handle an issue. len(list)==2 and indices are off
	list = RangeReduce(list)
	list = RangeReduce(list)
	return list
}

func Nope(input []SensorBeacon, row int) int {
	ranges := Coverage(input, row)

	// any sensors or beacons on this row
	grid := make(map[util.Point]struct{})
	for _, sb := range input {
		if sb.Beacon.Y == row {
			grid[sb.Beacon] = struct{}{}
		}
		if sb.Sensor.Y == row {
			grid[sb.Sensor] = struct{}{}
		}
	}

	count := 0
	for _, seg := range ranges {
		count += seg[1] - seg[0] + 1
	}
	return count - len(grid)
}

func part1() {
	fmt.Printf("Part 1: %v\n", Nope(parse("input.txt"), 2000000))
}

func part2() {
	input := parse("input.txt")
	for row := 0; row <= 4000000; row++ {
		ranges := Coverage(input, row)
		if len(ranges) == 2 && ranges[0][1]+2 == ranges[1][0] {
			x := ranges[0][1] + 1
			freq := 4000000*x + row
			fmt.Printf("Part 2: %v (%v,%v)\n", freq, x, row)
			break
		}
	}
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
