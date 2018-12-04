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

type TimeInfoSlice []*TimeInfo

func (tis TimeInfoSlice) Len() int {
	return len(tis)
}
func (tis TimeInfoSlice) Swap(i, j int) {
	tis[i], tis[j] = tis[j], tis[i]
}
func (tis TimeInfoSlice) Less(i, j int) bool {
	return tis[i].Time.Before(tis[j].Time)
}

var EventState = struct {
	Asleep string
	Awake  string
	Begin  string
}{
	Asleep: "falls asleep",
	Awake:  "wakes up",
	Begin:  "begins shift",
}

func part1() {
	fmt.Println("Day 04 Part 1")

	file, _ := os.Open("test.txt")
	defer file.Close()

	// year-month-day hour:minute
	// Mon Jan 2 15:04:05 -0700 MST 2006
	// RFC3339     =   "2006-01-02T15:04:05Z07:00"
	const timeFormat = "2006-01-02 15:04"

	var events TimeInfoSlice

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// [1518-11-01 00:00] Guard #10 begins shift
		// [1518-11-01 00:05] falls asleep
		// [1518-11-01 00:25] wakes up
		lineTime := line[1:17]
		lineInfo := line[19:]

		t, err := time.Parse(timeFormat, lineTime)
		if err != nil {
			fmt.Printf("Bad time found when time parsing. %v\n", lineTime)
			break
		}

		info := &TimeInfo{
			Time:  t,
			event: lineInfo,
		}
		if lineInfo != EventState.Asleep && lineInfo != EventState.Awake {
			infoParts := strings.Split(lineInfo, " ")
			idPart := infoParts[1][1:]
			info.ID, err = strconv.Atoi(idPart)
			if err != nil {
				fmt.Printf("Bad time when parsing Guard ID: %v\n", idPart)
				break
			}
		}

		events = append(events, info)
	}

	sort.Sort(events)

	// now need to copy the guard id to all of its awake and asleep events
	var guard int
	for _, e := range events {
		if e.event == EventState.Asleep {
			e.ID = guard
		} else if e.event == EventState.Awake {
			e.ID = guard
		} else {
			// new guard is begining
			guard = e.ID
		}
	}

	for _, e := range events {
		fmt.Println(e)
	}
}
