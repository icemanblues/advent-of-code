package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/icemanblues/advent-of-code/2020/util"
)

const (
	dayNum   = "02"
	dayTitle = "Password Philosophy"
)

type PasswordPolicy struct {
	Min    int
	Max    int
	Char   rune
	Passwd []rune
}

func parsePasswordPolicy(line string) PasswordPolicy {
	rule := strings.Split(line, ":")
	password := strings.TrimSpace(rule[1])
	rule = strings.Split(rule[0], " ")
	char := rule[1]
	times := strings.Split(rule[0], "-")
	min, _ := strconv.Atoi(times[0])
	max, _ := strconv.Atoi(times[1])
	charRune := []rune(char)[0]

	return PasswordPolicy{
		min, max, charRune, []rune(password),
	}
}

func numPasswd(input []string) int {
	num := 0
	for _, line := range input {
		policy := parsePasswordPolicy(line)
		count := 0
		for _, r := range policy.Passwd {
			if policy.Char == r {
				count++
			}
		}
		if count >= policy.Min && count <= policy.Max {
			num++
		}
	}
	return num
}

func xorPasswd(input []string) int {
	num := 0
	for _, line := range input {
		policy := parsePasswordPolicy(line)
		if (policy.Passwd[policy.Min-1] == policy.Char) != (policy.Passwd[policy.Max-1] == policy.Char) {
			num++
		}
	}
	return num
}

func test1() {
	input, _ := util.ReadInput("test1.txt")
	fmt.Printf("Test 1: %v\n", numPasswd(input))
}

func test2() {
	input, _ := util.ReadInput("test1.txt")
	fmt.Printf("Test 1: %v\n", xorPasswd(input))
}

func part1() {
	input, _ := util.ReadInput("input.txt")
	fmt.Printf("Part 1: %v\n", numPasswd(input))
}

func part2() {
	input, _ := util.ReadInput("input.txt")
	fmt.Printf("Part 2: %v\n", xorPasswd(input))
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	test1()
	part1()
	test2()
	part2()
}
