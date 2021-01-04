package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

const (
	dayNum   = "04"
	dayTitle = "The Ideal Stocking Stuffer"
)

const INPUT = "bgvyzdsv"

func AdventCoin(base string, num int) string {
	s := base + strconv.Itoa(num)
	md5sum := md5.Sum([]byte(s))
	return hex.EncodeToString(md5sum[:])
}

func SearchZero(base string, numZeroes int) int {
	runes := make([]rune, 0, numZeroes)
	for i := 0; i < numZeroes; i++ {
		runes = append(runes, '0')
	}
	zeroes := string(runes)

	i := 1
	coin := AdventCoin(base, i)
	for !strings.HasPrefix(coin, zeroes) {
		i++
		coin = AdventCoin(INPUT, i)
	}
	return i
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	fmt.Printf("Part 1: %v\n", SearchZero(INPUT, 5))
	fmt.Printf("Part 1: %v\n", SearchZero(INPUT, 6))
}
