package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "13"
	dayTitle = "Distress Signal"
)

type Pair struct {
	one, two []any
}

// True if in proper order. False if not
// True if the first return is real. If false, then its a draw
func Compare(one, two []any) (bool, bool) {
	l := len(one)
	if ll := len(two); ll > l {
		l = ll
	}

	for i := 0; i < l; i++ {
		if len(two) <= i {
			return false, true
		}
		if len(one) <= i {
			return true, true
		}

		a, b := one[i], two[i]
		aType, bType := reflect.ValueOf(a).Kind(), reflect.ValueOf(b).Kind()
		if aType == reflect.Float64 && bType == reflect.Float64 { // both floats
			af := a.(float64)
			bf := b.(float64)
			if af < bf {
				return true, true
			}
			if af > bf {
				return false, true
			}
		} else if aType == reflect.Slice && bType == reflect.Slice { // both slices
			as := a.([]any)
			bs := b.([]any)
			res, ok := Compare(as, bs)
			if ok {
				return res, ok
			}
		} else if aType == reflect.Float64 && bType == reflect.Slice { // float, slice
			aa := []any{a}
			bb := b.([]any)
			res, ok := Compare(aa, bb)
			if ok {
				return res, ok
			}
		} else { // slice, floa
			aa := a.([]any)
			bb := []any{b}
			res, ok := Compare(aa, bb)
			if ok {
				return res, ok
			}
		}
	}

	return true, false
}

func parse(filename string) []Pair {
	input, _ := util.ReadInput(filename)
	pair := Pair{nil, nil}
	var packets []Pair
	for _, line := range input {
		if line == "" {
			packets = append(packets, pair)
			pair = Pair{nil, nil}
			continue
		}
		if pair.one == nil {
			json.Unmarshal([]byte(line), &pair.one)
		} else {
			json.Unmarshal([]byte(line), &pair.two)
		}
	}
	packets = append(packets, pair)
	return packets
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	pairs := parse("input.txt")
	sum := 0
	for i, p := range pairs {
		if c, _ := Compare(p.one, p.two); c {
			sum += i + 1
		}
	}
	fmt.Printf("Part 1: %v\n", sum)

	var packets [][]any
	for _, p := range pairs {
		packets = append(packets, p.one)
		packets = append(packets, p.two)
	}
	var div1, div2 []any
	json.Unmarshal([]byte("[[2]]"), &div1)
	json.Unmarshal([]byte("[[6]]"), &div2)
	packets = append(packets, div1)
	packets = append(packets, div2)
	sort.Slice(packets, func(i, j int) bool {
		a, _ := Compare(packets[i], packets[j])
		return a
	})

	product := 1
	t1, _ := json.Marshal(div1)
	t2, _ := json.Marshal(div2)
	s1 := string(t1)
	s2 := string(t2)
	for i, packet := range packets {
		m, _ := json.Marshal(packet)
		if s := string(m); s == s1 || s == s2 {
			product *= i + 1
		}
	}
	fmt.Printf("Part 2: %v\n", product)
}
