package main

import (
	"fmt"
	"strconv"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

func parse(filename string, presize int) (map[int]struct{}, []int) {
	lines, _ := util.ReadInput(filename)
	preamble := make(map[int]struct{})
	message := make([]int, 0)

	for i, line := range lines {
		num, _ := strconv.Atoi(line)
		if i < presize {
			preamble[num] = struct{}{}
		}
		message = append(message, num)
	}

	return preamble, message
}

func isSumPreamble(preamble map[int]struct{}, target int) bool {
	for k := range preamble {
		t := target - k
		if _, ok := preamble[t]; ok && k != t {
			return true
		}
	}
	return false
}

func findBadSum(premable map[int]struct{}, message []int) int {
	for i := len(premable); i < len(message); i++ {
		target := message[i]
		if !isSumPreamble(premable, target) {
			return target
		}

		delete(premable, message[i-len(premable)])
		premable[target] = struct{}{}
	}

	panic("no bad sum")
}

func findMinMaxForBadSum(msg []int, target int) int {
search:
	for i, n := range msg {
		sum := n
		for j := i + 1; j < len(msg)-1; j++ {
			num := msg[j]
			sum += num
			if sum > target {
				continue search
			}
			if sum == target {
				min, max := n, n
				for k := i; k <= j; k++ {
					if msg[k] < min {
						min = msg[k]
					}
					if msg[k] > max {
						max = msg[k]
					}
				}
				return min + max
			}
		}
	}

	panic("no sum found for target")
}

func part1() {
	pre, msg := parse("input.txt", 25)
	fmt.Printf("Part 1: %v\n", findBadSum(pre, msg))
}

func part2() {
	pre, msg := parse("input.txt", 25)
	fmt.Printf("Part 2: %v\n", findMinMaxForBadSum(msg, findBadSum(pre, msg)))
}

func main() {
	fmt.Println("Day 09: Encoding Error")
	part1()
	part2()
}
