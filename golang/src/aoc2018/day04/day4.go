package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("--- Day 4: Repose Record ---")

	part1()
}

type TimeInfo struct {
	ID    int
	Time  time.Time
	event string
}

type TimeInfoSlice []TimeInfo

func (tis TimeInfoSlice) Len() int {
	return len(tis)
}
func (tis TimeInfoSlice) Swap(i, j int) {
	tis[i], tis[j] = tis[j], tis[i]
}
func (tis TimeInfoSlice) Less(i, j int) bool {
	return tis[i].Time.Before(tis[j].Time)
}

func part1() {
	fmt.Println("Day 04 Part 1")

	file, _ := os.Open("input04.txt")
	defer file.Close()

	const timeFormat = "2006-01-04 15:05"

	var events TimeInfoSlice

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		// [1518-11-01 00:00] Guard #10 begins shift
		// [1518-11-01 00:05] falls asleep
		// [1518-11-01 00:25] wakes up
		lineTime := string(runes[1:17])
		lineInfo := string(runes[19:])

		t, err := time.Parse(timeFormat, lineTime)
		if err != nil {
			fmt.Printf("Bad time found when time parsing. %v\n", lineTime)
			break
		}

		info := TimeInfo{
			Time:  t,
			event: lineInfo,
		}
		if lineInfo != "falls asleep" && lineInfo != "wakes up" {
			infoParts := strings.Split(lineInfo, " ")
			idPart := infoParts[1][1:]
			info.ID, err = strconv.Atoi(idPart)
			if err != nil {
				fmt.Printf("Bad time when parsing Guard ID: %v\n", idPart)
			}
		}

		events = append(events, info)
	}

	sort.Sort(events)

	// now need to copy the guard id to all of its awake and asleep events
	for _, e := range events[:15] {
		fmt.Println(e)
	}
}
