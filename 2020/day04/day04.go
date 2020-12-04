package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/icemanblues/advent-of-code/2020/util"
)

const (
	dayNum   = "04"
	dayTitle = "Title"
)

func parse(filename string) []map[string]string {
	lines, _ := util.ReadInput(filename)
	var passports []map[string]string

	pp := make(map[string]string)
	for _, line := range lines {
		// empty, so add it and clear it
		if len(strings.TrimSpace(line)) == 0 {
			passports = append(passports, pp)
			pp = make(map[string]string)
			continue
		}

		pairs := strings.Split(line, " ")
		for _, pair := range pairs {
			keyValue := strings.Split(pair, ":")
			pp[keyValue[0]] = keyValue[1]
		}
	}

	passports = append(passports, pp)
	return passports
}

func hasRequiredFIelds(pp map[string]string) bool {
	_, byr_ok := pp["byr"]
	_, iyr_ok := pp["iyr"]
	_, eyr_ok := pp["eyr"]
	_, hgt_ok := pp["hgt"]
	_, hcl_ok := pp["hcl"]
	_, ecl_ok := pp["ecl"]
	_, pid_ok := pp["pid"]
	//_, cid_ok := pp["cid"]
	return byr_ok && iyr_ok && eyr_ok && hgt_ok && hcl_ok && ecl_ok && pid_ok /*&& cid_ok*/
}

func validPassport(pp map[string]string) bool {
	byr, byr_ok := pp["byr"]
	if !byr_ok {
		return false
	}
	if len(byr) != 4 {
		return false
	}
	birthYear, err := strconv.Atoi(byr)
	if err != nil {
		return false
	}
	if birthYear < 1920 || birthYear > 2002 {
		return false
	}

	iyr, iyr_ok := pp["iyr"]
	if !iyr_ok {
		return false
	}
	issueYear, err := strconv.Atoi(iyr)
	if err != nil {
		return false
	}
	if issueYear < 2010 || issueYear > 2020 {
		return false
	}

	eyr, eyr_ok := pp["eyr"]
	if !eyr_ok {
		return false
	}
	expYear, err := strconv.Atoi(eyr)
	if err != nil {
		return false
	}
	if expYear < 2020 || expYear > 2030 {
		return false
	}

	hgt, hgt_ok := pp["hgt"]
	if !hgt_ok {
		return false
	}
	h := []rune(hgt)
	unit := string(h[len(h)-2:])
	height, _ := strconv.Atoi(string(h[:len(h)-2]))
	if unit != "in" && unit != "cm" {
		return false
	}
	if unit == "cm" {
		if height < 150 || height > 193 {
			return false
		}
	}
	if unit == "in" {
		if height < 59 || height > 76 {
			return false
		}
	}

	hcl, hcl_ok := pp["hcl"]
	if !hcl_ok {
		return false
	}
	if len(hcl) != 7 {
		return false
	}
	matched, err := regexp.MatchString("#[0-9a-f]+", hcl)
	if !matched {
		return false
	}

	ecl, ecl_ok := pp["ecl"]
	if !ecl_ok {
		return false
	}
	good := ecl == "amb" || ecl == "blu" || ecl == "brn" || ecl == "gry" || ecl == "grn" || ecl == "hzl" || ecl == "oth"
	if !good {
		return false
	}

	pid, pid_ok := pp["pid"]
	if !pid_ok {
		return false
	}
	if len(pid) != 9 {
		return false
	}
	_, err = strconv.Atoi(pid)
	if err != nil {
		return false
	}

	return true
}

func validCount(passports []map[string]string) int {
	count := 0
	for _, pp := range passports {
		if validPassport(pp) {
			count++
		}
	}
	return count
}

func part1() {
	input := parse("input.txt")
	fmt.Println(len(input))

	fmt.Printf("Part 1: %v\n", validCount(input))
}

func part2() {
	fmt.Println("Part 2")
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
