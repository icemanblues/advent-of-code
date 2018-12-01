package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readFile(fn string) {
	file, _ := os.Open(fn)
	defer file.Close()

	ln := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("line number %v: %s\n", ln, line)

		ln++
	}
}

func main() {
	fmt.Println("Day 1")

	file, _ := os.Open("input01.txt")
	defer file.Close()

	var deltas []int

	ln := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		i, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("not a valid int. %v\n", i)
		}
		deltas = append(deltas, i)
		ln++
	}

	freq := 0
	freqSet := make(map[int]bool)

	index := 0
	_, ok := freqSet[freq]
	for !ok {
		freqSet[freq] = true

		safeIdx := index % len(deltas)
		freq += deltas[safeIdx]

		index++
		_, ok = freqSet[freq]
	}

	fmt.Printf("frequency: %v in %v iterations %v deltas\n", freq, index, len(deltas))
}
