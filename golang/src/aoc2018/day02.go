package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Day 2: Inventory Management System")

	//part1()
	part2()
}

func part2() {
	file, _ := os.Open("input02.txt")
	defer file.Close()

	var boxIDs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		boxID := scanner.Text()
		boxIDs = append(boxIDs, boxID)
	}

	for i := 0; i < len(boxIDs)-1; i++ {
		for j := i + 1; j < len(boxIDs); j++ {
			if diff1(boxIDs[i], boxIDs[j]) {
				solution := intersect(boxIDs[i], boxIDs[j])
				fmt.Printf("solution: %v\n", solution)
			}
		}
	}
}

func diff1(s, t string) bool {
	c := 0

	for i := range s {
		if s[i] != t[i] {
			c++
		}
	}

	return c == 1
}

func intersect(s, t string) string {
	var same []byte

	for i := range s {
		if s[i] == t[i] {
			same = append(same, s[i])
		}
	}

	return string(same)
}

func part1() {
	file, _ := os.Open("input02.txt")
	defer file.Close()

	count2 := 0
	count3 := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		boxID := scanner.Text()
		c2, c3 := count23(boxID)
		if c2 {
			count2++
		}
		if c3 {
			count3++
		}
	}

	checksum := count2 * count3
	fmt.Printf("checksum: %v\n", checksum)
}

func count23(s string) (bool, bool) {
	count := make(map[rune]int)
	for _, r := range s {
		c, ok := count[r]
		if !ok {
			count[r] = 1
		} else {
			count[r] = c + 1
		}
	}

	c2 := false
	c3 := false
	for _, v := range count {
		if v == 2 {
			c2 = true
		} else if v == 3 {
			c3 = true
		}
	}

	return c2, c3
}
