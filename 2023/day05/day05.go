package main

import (
	"fmt"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "05"
	dayTitle = "If You Give A Seed A Fertilizer"
)

type MapRange struct {
	DestStart, SourceStart, Length int
}

func (mr MapRange) lookup(code int) (int, bool) {
	if diff := code - mr.SourceStart; mr.SourceStart <= code && diff < mr.Length {
		return mr.DestStart + diff, true
	}
	return 0, false
}

func Lookup(ranges []MapRange, code int) int {
	for _, mr := range ranges {
		v, ok := mr.lookup(code)
		if ok {
			return v
		}
	}
	return code
}

type Almanac struct {
	SeedToSoil            []MapRange
	SoilToFertilizer      []MapRange
	FertilizerToWater     []MapRange
	WaterToLight          []MapRange
	LightToTemperature    []MapRange
	TemperatureToHumidity []MapRange
	HumidityToLocation    []MapRange
}

func SliceAtoI(ss []string) []int {
	ints := make([]int, len(ss), len(ss))
	for i, s := range ss {
		ints[i] = util.MustAtoi(s)
	}
	return ints
}

func parseInput(filename string) (seeds []int, almanac Almanac) {
	input, _ := util.ReadInput(filename)
	seedTitleLine := strings.Split(input[0], ": ")
	seeds = SliceAtoI(strings.Fields(seedTitleLine[1]))

	var currMR []MapRange
	for _, line := range input[1:] {
		if line == "" {
			continue
		}

		switch line {
		case "seed-to-soil map:":
			currMR = almanac.SeedToSoil
			continue
		case "soil-to-fertilizer map:":
			almanac.SeedToSoil = currMR
			currMR = almanac.SoilToFertilizer
			continue
		case "fertilizer-to-water map:":
			almanac.SoilToFertilizer = currMR
			currMR = almanac.FertilizerToWater
			continue
		case "water-to-light map:":
			almanac.FertilizerToWater = currMR
			currMR = almanac.WaterToLight
			continue
		case "light-to-temperature map:":
			almanac.WaterToLight = currMR
			currMR = almanac.LightToTemperature
			continue
		case "temperature-to-humidity map:":
			almanac.LightToTemperature = currMR
			currMR = almanac.TemperatureToHumidity
			continue
		case "humidity-to-location map:":
			almanac.TemperatureToHumidity = currMR
			currMR = almanac.HumidityToLocation
			continue
		}

		rangeFields := strings.Fields(line)
		currMR = append(currMR, MapRange{
			DestStart:   util.MustAtoi(rangeFields[0]),
			SourceStart: util.MustAtoi(rangeFields[1]),
			Length:      util.MustAtoi(rangeFields[2]),
		})

	}
	almanac.HumidityToLocation = currMR
	return
}

func Garden(almanac Almanac, seed int) int {
	soil := Lookup(almanac.SeedToSoil, seed)
	fert := Lookup(almanac.SoilToFertilizer, soil)
	water := Lookup(almanac.FertilizerToWater, fert)
	light := Lookup(almanac.WaterToLight, water)
	temp := Lookup(almanac.LightToTemperature, light)
	humid := Lookup(almanac.TemperatureToHumidity, temp)
	loc := Lookup(almanac.HumidityToLocation, humid)
	return loc
}

func part1() {
	seeds, almanac := parseInput("input.txt")
	low := 0
	for _, seed := range seeds {
		loc := Garden(almanac, seed)
		if low == 0 || loc < low {
			low = loc
		}
	}
	fmt.Printf("Part 1: %v\n", low)
}

func search(a Almanac, s, e int) int {
	sloc, eloc := Garden(a, s), Garden(a, e)
	if s == e-1 {
		return util.Min(sloc, eloc)
	}

	m := (s + e) / 2
	mloc := Garden(a, m)
	// only if there is a trend line between start, mid, end can I be sure this is where the min lies
	// otherwise need to partition and find the min between the two partitions
	if sloc < mloc && mloc < eloc {
		return search(a, s, m)
	} else if eloc < mloc && mloc < sloc {
		return search(a, m, e)
	} else {
		return util.Min(search(a, s, m), search(a, m, e))
	}
}

func part2BinarySearch() {
	seeds, almanac := parseInput("input.txt")
	low := 0

	for i := 0; i < len(seeds)-1; i = i + 2 {
		start, len := seeds[i], seeds[i+1]
		end := start + len - 1
		loc := search(almanac, start, end)
		if low == 0 || loc < low {
			low = loc
		}
	}
	fmt.Printf("Part 2 Binary  Search: %v\n", low)
}

func (mr MapRange) rlookup(code int) (int, bool) {
	if diff := code - mr.DestStart; mr.DestStart <= code && diff < mr.Length {
		return mr.SourceStart + diff, true
	}
	return 0, false
}

func ReverseLookup(ranges []MapRange, code int) int {
	for _, mr := range ranges {
		if v, ok := mr.rlookup(code); ok {
			return v
		}
	}
	return code
}

func reverse(almanac Almanac, loc int) int {
	humid := ReverseLookup(almanac.HumidityToLocation, loc)
	temp := ReverseLookup(almanac.TemperatureToHumidity, humid)
	light := ReverseLookup(almanac.LightToTemperature, temp)
	water := ReverseLookup(almanac.WaterToLight, light)
	fert := ReverseLookup(almanac.FertilizerToWater, water)
	soil := ReverseLookup(almanac.SoilToFertilizer, fert)
	seed := ReverseLookup(almanac.SeedToSoil, soil)
	return seed
}

func part2ReverseLookup() {
	seeds, almanac := parseInput("input.txt")
	var pairs [][2]int
	for i := 0; i < len(seeds)-1; i = i + 2 {
		start, len := seeds[i], seeds[i+1]
		end := start + len - 1
		pairs = append(pairs, [2]int{start, end})
	}
	low, seek := -1, true
	for seek {
		low++
		seed := reverse(almanac, low)
		for _, pair := range pairs {
			if seed >= pair[0] && seed <= pair[1] {
				seek = false
				break
			}
		}
	}
	fmt.Printf("Part 2 Reverse Lookup: %v\n", low)
}

func part2() {
	part2ReverseLookup()
	part2BinarySearch()
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
