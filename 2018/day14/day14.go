package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("something bad with file: %v\n", err)
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

const (
	test    = 9
	input14 = 430971
	test2   = 59414
)

func main() {
	fmt.Println("Day 14: Chocolate Charts")

	// part1(input14)
	part2(input14)
}

func part1(input int) ([]int, []int) {
	fmt.Println("Part 1")

	scores := []int{3, 7}

	e1idx := 0
	e2idx := 1

	// create new recipes
	tick := 0
	for len(scores) < input+10 {
		tick++
		// fmt.Printf("%v: elf: %v elf2: %v\n", tick, e1idx, e2idx)
		// fmt.Printf("%v: %v\n", tick, scores)

		// sum the last two recipes
		recipes := scores[e1idx] + scores[e2idx]
		if recipes >= 10 {
			r1 := recipes / 10
			r2 := recipes % 10
			scores = append(scores, r1, r2)
		} else {
			scores = append(scores, recipes)
		}

		//advance the elves
		e1idx += (1 + scores[e1idx])
		e1idx = e1idx % len(scores)
		e2idx += (1 + scores[e2idx])
		e2idx = e2idx % len(scores)
	}

	// fmt.Println(scores)
	answer := scores[len(scores)-10:]
	fmt.Printf("last 10: %v\n", answer)

	return scores, answer
}

func testEq(a, b []int) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func part2(input int) {
	fmt.Println("Part 2")

	var t []int
	s := strconv.Itoa(input)
	for i := range s {
		r, _ := strconv.Atoi(string(s[i]))
		t = append(t, r)
	}
	l := len(t)
	i := 0

	scores := []int{3, 7}
	e1idx := 0
	e2idx := 1
	// create new recipes
	tick := 0
	for true { // len(scores) < input+10
		if tick%1000 == 0 {
			fmt.Printf("tick: %v i: %v t: %v\n", tick, i, t)
		}

		// sum the last two recipes
		recipes := scores[e1idx] + scores[e2idx]
		if recipes >= 10 {
			r1 := recipes / 10
			r2 := recipes % 10
			scores = append(scores, r1, r2)
		} else {
			scores = append(scores, recipes)
		}

		//advance the elves
		e1idx += (1 + scores[e1idx])
		e1idx = e1idx % len(scores)
		e2idx += (1 + scores[e2idx])
		e2idx = e2idx % len(scores)

		// check if we have what we need
		if len(scores)-i > l {
			if testEq(scores[i:i+l], t) {
				fmt.Printf("found it at idx: %v\n", i)
				break
			}
			i++
		}

		tick++
	}
}
