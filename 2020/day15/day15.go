package main

import (
	"fmt"
)

const (
	dayNum   = "15"
	dayTitle = "Rambunctious Recitation"
)

var input = []int{0, 13, 1, 16, 6, 17}

func searchBackwards(numbers []int, num int) int {
	for i := len(numbers) - 1; i >= 0; i-- {
		if numbers[i] == num {
			return i + 1 // the turn is index+1
		}
	}
	return 0 // not found
}

func MemGame(starting []int, haltTurnNum int) int {
	for turnNum := len(starting) + 1; turnNum <= haltTurnNum; turnNum++ {
		lastTurnNum := turnNum - 1
		lastSpoken := starting[lastTurnNum-1]
		lastSpokenTurn := searchBackwards(starting[:lastTurnNum-1], lastSpoken)

		age := 0
		if lastSpokenTurn != 0 {
			age = lastTurnNum - lastSpokenTurn
		}
		starting = append(starting, age)
	}

	return starting[haltTurnNum-1]
}

func MemGameImproved(starting []int, haltTurnNum int) int {
	mem := make(map[int]int)
	for i, n := range starting {
		mem[n] = i + 1
	}

	var deferTurn int = -1
	var deferSpoken int = -1

	lastSpoken := starting[len(starting)-1]
	for turnNum := len(starting) + 1; turnNum <= haltTurnNum; turnNum++ {
		lastTurnNum := turnNum - 1
		lastSpokenTurn := mem[lastSpoken]
		age := 0
		if lastSpokenTurn != 0 {
			age = lastTurnNum - lastSpokenTurn
		}
		starting = append(starting, age)

		// add it to memory (to defer, then memory)
		mem[deferSpoken] = deferTurn
		deferSpoken = age
		deferTurn = turnNum

		lastSpoken = age
	}

	return lastSpoken
}

func part1() {
	fmt.Printf("Part 1: %v\n", MemGame(input, 2020))
}

func part2() {
	fmt.Printf("Part 2: %v\n", MemGameImproved(input, 30000000))
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	testPart1()
	part1()
	testPart2()
	part2()
}

func testPart1() {
	tests := []struct {
		numbers  []int
		expected int
	}{
		{[]int{1, 3, 2}, 1},
		{[]int{2, 1, 3}, 10},
		{[]int{1, 2, 3}, 27},
		{[]int{2, 3, 1}, 78},
		{[]int{3, 2, 1}, 438},
		{[]int{3, 1, 2}, 1836},
	}

	fmt.Println("testPart")
	for _, test := range tests {
		actual := MemGame(test.numbers, 2020)
		if actual == test.expected {
			fmt.Printf("PASS! actual: %v expected: %v input: %v\n", actual, test.expected, test.numbers)
		} else {
			fmt.Printf("FAIL! actual: %v expected: %v input: %v\n", actual, test.expected, test.numbers)
		}

	}
}

func testPart2() {
	tests := []struct {
		numbers  []int
		expected int
	}{
		{[]int{0, 3, 6}, 175594},
		{[]int{1, 3, 2}, 2578},
		{[]int{2, 1, 3}, 3544142},
		{[]int{1, 2, 3}, 261214},
		{[]int{2, 3, 1}, 6895259},
		{[]int{3, 2, 1}, 18},
		{[]int{3, 1, 2}, 362},
	}

	fmt.Println("testPart2")
	for _, test := range tests {
		actual := MemGameImproved(test.numbers, 30000000)
		if actual == test.expected {
			fmt.Printf("Pass! %v == %v for %v \n", actual, test.expected, test.numbers)
		} else {
			fmt.Printf("FAIL! %v == %v for %v \n", actual, test.expected, test.numbers)
		}
	}
}
