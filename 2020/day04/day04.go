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
	dayTitle = "Passport Processing"
)

type Passport map[string]string

func Parse(filename string) []Passport {
	lines, _ := util.ReadInput(filename)
	var passports []Passport

	pp := make(Passport)
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
	// EOF, so add it
	passports = append(passports, pp)
	return passports
}

func HasRequiredFIelds(pp Passport) bool {
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

func validYear(pp Passport, field string, min, max int) bool {
	value, ok := pp[field]
	if !ok {
		return false
	}
	year, err := strconv.Atoi(value)
	if err != nil {
		return false
	}
	if year < min || year > max {
		return false
	}

	return true
}

func validHeight(pp Passport) bool {
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

	return true
}

func validHaircolor(pp Passport) bool {
	hcl, hcl_ok := pp["hcl"]
	if !hcl_ok {
		return false
	}
	if len(hcl) != 7 {
		return false
	}
	matched, _ := regexp.MatchString("#[0-9a-f]+", hcl)
	if !matched {
		return false
	}

	return true
}

var valid_ecl = map[string]struct{}{
	"amb": {},
	"blu": {},
	"brn": {},
	"gry": {},
	"grn": {},
	"hzl": {},
	"oth": {},
}

func validEyeColor(pp Passport) bool {
	ecl, ecl_ok := pp["ecl"]
	if !ecl_ok {
		return false
	}
	if _, ok := valid_ecl[ecl]; !ok {
		return false
	}

	return true
}

func validPassportID(pp Passport) bool {
	pid, pid_ok := pp["pid"]
	if !pid_ok {
		return false
	}
	if len(pid) != 9 {
		return false
	}
	_, err := strconv.Atoi(pid)
	if err != nil {
		return false
	}

	return true
}

func ValidPassport(pp Passport) bool {
	if !validYear(pp, "byr", 1920, 2002) {
		return false
	}

	if !validYear(pp, "iyr", 2010, 2020) {
		return false
	}

	if !validYear(pp, "eyr", 2020, 2030) {
		return false
	}

	if !validHeight(pp) {
		return false
	}

	if !validHaircolor(pp) {
		return false
	}

	if !validEyeColor(pp) {
		return false
	}

	if !validPassportID(pp) {
		return false
	}

	return true
}

func countValid(passports []Passport, validator func(Passport) bool) int {
	count := 0
	for _, pp := range passports {
		if validator(pp) {
			count++
		}
	}
	return count
}

func part1() {
	input := Parse("input.txt")
	fmt.Printf("Part 1: %v\n", countValid(input, HasRequiredFIelds))
}

func part2() {
	input := Parse("input.txt")
	fmt.Printf("Part 2: %v\n", countValid(input, ValidPassport))
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
