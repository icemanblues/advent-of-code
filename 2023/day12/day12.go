package main

import (
	"fmt"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "12"
	dayTitle = "Hot Springs"
)

func parse(filename string) ([]string, [][]int) {
	input, _ := util.ReadInput(filename)
	springs := make([]string, 0, len(input))
	blocks := make([][]int, 0, len(input))

	for _, line := range input {
		leftRight := strings.Fields(line)
		springs = append(springs, leftRight[0])
		blocks = append(blocks, util.SliceAtoI(strings.Split(leftRight[1], ",")))
	}
	return springs, blocks
}

func questionCount(spring string) int {
	questions := 0
	for _, r := range spring {
		if r == '?' {
			questions++
		}
	}
	return questions
}

func blockCount(spring []rune) []int {
	var counts []int
	c := 0
	for _, r := range spring {
		switch r {
		case '#':
			c++
		case '.':
			if c != 0 {
				counts = append(counts, c)
				c = 0
			}
		default:
			c = 0
			break
		}
	}
	if c != 0 {
		counts = append(counts, c)
	}
	return counts
}

func isValid(spring string, blocks []int, guess []rune, count int) bool {
	arr := make([]rune, 0, len(spring))
	j := 0
	for _, r := range spring {
		if r != '?' {
			arr = append(arr, r)
			continue
		}
		if j < len(guess) {
			arr = append(arr, guess[j])
			j++
		} else {
			arr = append(arr, '?')
			break
		}
	}
	counts := blockCount(arr)
	//fmt.Printf("Checking Validity: %c %v %v %c\n", arr, blocks, counts, guess)
	if count == len(guess) { // do a full slice compare
		if len(counts) != len(blocks) {
			return false
		}
		for i, c := range counts {
			if c != blocks[i] {
				return false
			}
		}
		return true
	}

	for i, c := range counts {
		if i >= len(blocks) {
			return false
		}
		if c != blocks[i] {
			return false
		}
	}
	return true
}

func arrangements(spring string, blocks []int, guess []rune, count int) int {
	// check if it is valid
	if !isValid(spring, blocks, guess, count) {
		//fmt.Printf("Not Valid: %c\n", guess)
		return 0
	}
	// made it this far, so must be a valid arrangement
	if len(guess) == count {
		//fmt.Printf("Valid: %c\n", guess)
		return 1
	}

	sum := 0
	l := len(guess)
	for _, r := range "#." {
		guess = append(guess, r)
		sum += arrangements(spring, blocks, guess, count)
		guess = guess[:l]
	}
	return sum
}

func part1() {
	springs, blocks := parse("input.txt")
	sum := 0
	for i, spring := range springs {
		q := questionCount(spring)
		guess := make([]rune, 0, q)
		arrange := arrangements(spring, blocks[i], guess, q)
		sum += arrange
	}
	fmt.Printf("Part 1: %v\n", sum)
}

func unfold(spring string, blocks []int) (string, []int) {
	sb := strings.Builder{}
	sb.WriteString(spring)
	sb.WriteRune('?')
	sb.WriteString(spring)
	sb.WriteRune('?')
	sb.WriteString(spring)
	sb.WriteRune('?')
	sb.WriteString(spring)
	sb.WriteRune('?')
	sb.WriteString(spring)

	unfoldedBlocks := make([]int, 0, 5*len(blocks))
	for i := 0; i < 5; i++ {
		unfoldedBlocks = append(unfoldedBlocks, blocks...)
	}

	return sb.String(), unfoldedBlocks
}

func part2() {
	springs, blocks := parse("input.txt")
	sum := 0
	for i, spring := range springs {
		unfoldedSpring, unfoldedBlocks := unfold(spring, blocks[i])
		q := questionCount(unfoldedSpring)
		guess := make([]rune, 0, q)
		arrange := arrangements(unfoldedSpring, unfoldedBlocks, guess, q)
		fmt.Printf("starting on %v of %v. %v Progress %.2f\n", i, len(springs), arrange, float64(i)/float64(len(springs)))
		sum += arrange
	}
	fmt.Printf("Part 2: %v\n", sum)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
