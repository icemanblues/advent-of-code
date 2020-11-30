package util

import (
	"bufio"
	"os"
	"strconv"
)

func ReadInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func ReadIntput(filename string) ([]int, error) {
	lines, err := ReadInput(filename)
	if err != nil {
		return nil, err
	}

	nums := make([]int, 0, len(lines))
	for _, e := range lines {
		n, err := strconv.Atoi(e)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}
