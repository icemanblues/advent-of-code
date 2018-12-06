package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Day 5: Alchemical Reduction")

	file, err := os.Open("input05.txt")
	if err != nil {
		fmt.Printf("something bad with file: %v\n", err)
		return
	}
	defer file.Close()

	var polymer string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		polymer = scanner.Text()
	}
	// fmt.Printf("polym: %v\n", polymer)

	// part1(polymer)
	part2(polymer)
}

func part2(polymer string) {
	alpha := "abcdefghijklmnopqrstuvwxyz"
	fmt.Println(len(alpha))

	letterLen := make(map[string]int)
	for i, _ := range alpha {
		a := string(alpha[i])
		s := strings.Replace(polymer, a, "", -1)
		s = strings.Replace(s, strings.ToUpper(a), "", -1)
		// fmt.Printf("removed %v : %v\n", a, s)
		letterLen[a] = part1(s)
	}

	min := 10000000
	for _, v := range letterLen {
		if v < min {
			min = v
		}
	}

	fmt.Printf("shortest possible: %v\n", min)
}

func pair(b byte) byte {
	s := string(b)
	if s == strings.ToUpper(s) {
		return strings.ToLower(s)[0]
	}

	return strings.ToUpper(s)[0]
}

func part1(polymer string) int {
	fmt.Println("Day 5 - Part 1")

	var s []byte
	for i, _ := range polymer {
		if len(s) == 0 {
			s = append(s, polymer[i])
			continue
		}

		if pair(polymer[i]) == s[len(s)-1] {
			s = s[:len(s)-1]
			continue
		}

		s = append(s, polymer[i])
	}

	// fmt.Printf("part1: %v\n", len(s))
	return len(s)
}
