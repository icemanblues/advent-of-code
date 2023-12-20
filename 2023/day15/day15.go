package main

import (
	"fmt"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "15"
	dayTitle = "Lens Library"
)

func parse(filename string) []string {
	return strings.Split(util.MustRead(filename)[0], ",")
}

func Hash(s string) (hash int) {
	for _, r := range s {
		hash += int(r)
		hash *= 17
		hash = hash % 256
	}
	return
}

func part1() {
	sum := 0
	for _, s := range parse("input.txt") {
		sum += Hash(s)
	}
	fmt.Printf("Part 1: %v\n", sum)
}

type Lens struct {
	label string
	focal int
}

func part2() {
	hashmap := make(map[int][]Lens)
	for _, s := range parse("input.txt") {
		if strings.HasSuffix(s, "-") {
			label := string(string(s[:len(s)-1]))
			hash := Hash(label)
			list := hashmap[hash]
			result := make([]Lens, 0, util.Max(len(list)-1, 0))
			for _, lens := range list {
				if lens.label != label {
					result = append(result, lens)
				}
			}
			hashmap[hash] = result
			continue
		}

		values := strings.Split(s, "=")
		label, value := values[0], util.MustAtoi(values[1])
		hash := Hash(label)
		list := hashmap[hash]
		result := make([]Lens, 0, len(list)+1)
		isFound := false
		for _, l := range list {
			if l.label == label {
				result = append(result, Lens{label, value})
				isFound = true
			} else {
				result = append(result, l)
			}
		}
		if !isFound {
			result = append(result, Lens{label, value})
		}
		hashmap[hash] = result
	}

	sum := 0
	for h := 0; h < 256; h++ {
		list := hashmap[h]
		for i, lens := range list {
			sum += (h + 1) * (i + 1) * lens.focal
		}
	}
	fmt.Printf("Part 2: %v\n", sum)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
