package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readInput(filename string) (map[int]rune, map[string]rune) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("something bad with file: %v\n", err)
		return nil, nil
	}
	defer file.Close()

	lineNum := 0
	transitions := make(map[string]rune)
	state := make(map[int]rune)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineNum++

		// create the initial state
		if lineNum == 1 {
			inputState := strings.Split(line, " ")[2]
			for i, r := range inputState {
				if r == '#' {
					state[i] = '#'
				}
			}

			continue
		}
		if lineNum == 2 {
			continue
		}

		// build the transition map
		words := strings.Split(line, " ")
		transitions[words[0]] = []rune(words[2])[0]

	}
	return state, transitions
}

func main() {
	fmt.Println("Day 12: Subterranean Sustainability")

	part1()
	part2()
}

func part1() {
	fmt.Println("Part 1")

	state, transitions := readInput("input12.txt")

	numGen := 20
	// numGen := 50000000000
	for gen := 0; gen < numGen; gen++ {
		// printState(gen, state)

		// get the min & max idx so we know where to start and end
		minIdx := len(state)
		maxIdx := 0
		for idx, _ := range state {
			if idx < minIdx {
				minIdx = idx
			}
			if maxIdx < idx {
				maxIdx = idx
			}
		}

		nextState := make(map[int]rune)
		for i := minIdx - 2; i <= maxIdx+2; i++ {
			// check if a transition matches

			for pattern, trans := range transitions {
				match := true
				for j, r := range pattern {
					p, ok := state[i+j-2]
					if !ok {
						p = '.'
					}
					match = match && p == r
				}

				if match && trans == '#' {
					nextState[i] = trans
				}
				// if !match {
				// 	if r, ok := state[i]; ok {
				// 		nextState[i] = r
				// 	}
				// }
			}
		}

		state = nextState
	}

	// printState(numGen, state)
	fmt.Printf("number of plants: %v\n", len(state))

	sum := 0
	for k, _ := range state {
		sum += k
	}
	fmt.Println(sum)
}

func stringState(state map[int]rune) string {
	// get the min & max idx so we know where to start and end
	maxIdx, minIdx := 0, 0
	for idx, _ := range state {
		if idx < minIdx {
			minIdx = idx
		}
		if maxIdx < idx {
			maxIdx = idx
		}
	}

	sb := strings.Builder{}
	for i := minIdx - 1; i <= maxIdx+1; i++ {

		if r, ok := state[i]; !ok {
			sb.WriteRune('.')
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

func printState(gen int, state map[int]rune) {

	fmt.Printf("%2d: %v\n", gen, stringState(state))
}

func part2() {
	fmt.Println("Part 2")

	state, transitions := readInput("input12.txt")

	// numGen := 20
	numGen := 50000000000

	genMap := make(map[string]int)

	for gen := 0; gen < numGen; gen++ {
		// printState(gen, state)
		if gen%1000 == 0 {
			fmt.Printf("generation iteration: %v %v\n", gen, len(state))
		}

		// get the min & max idx so we know where to start and end
		minIdx := len(state)
		maxIdx := 0
		for idx, _ := range state {
			if idx < minIdx {
				minIdx = idx
			}
			if maxIdx < idx {
				maxIdx = idx
			}
		}

		nextState := make(map[int]rune)
		for i := minIdx - 1; i <= maxIdx+1; i++ {
			// check if a transition matches

			for pattern, trans := range transitions {
				match := true
				for j, r := range pattern {
					p, ok := state[i+j-2]
					if !ok {
						p = '.'
					}
					match = match && p == r
				}

				if match && trans == '#' {
					nextState[i] = trans
				}
			}
		}

		state = nextState
		s := stringState(state)
		if v, ok := genMap[s]; ok {
			fmt.Printf("loop detected. curr gen %v, old gen %v\n", gen, v)
			break
		}
		genMap[s] = gen
	}

	// printState(numGen, state)
	fmt.Printf("number of plants: %v\n", len(state))

	sum := 0
	for k, _ := range state {
		sum += k
	}
	fmt.Println(sum)
}
