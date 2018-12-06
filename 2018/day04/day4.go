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

	// part1()
	part2()
}

type TimeInfo struct {
	ID     int
	Minute int
	Time   time.Time
	event  string
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

	file, err := os.Open("input04.txt")
	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
		return
	}
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
			Time:   t,
			Minute: t.Minute(),
			event:  lineInfo,
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
	guardSleep := make(map[int]int)
	var start int
	var guard int
	for _, e := range events {
		if e.event == EventState.Asleep {
			e.ID = guard
			start = e.Minute
		} else if e.event == EventState.Awake {
			e.ID = guard

			sleep := e.Minute - start
			s, ok := guardSleep[guard]
			if !ok {
				guardSleep[guard] = sleep
			} else {
				guardSleep[guard] = s + sleep
			}
		} else {
			// new guard is begining
			guard = e.ID
		}
	}

	// find the max sleeping guard
	var max int
	var maxGuard int
	for g, s := range guardSleep {
		if s > max {
			maxGuard = g
			max = s
		}
	}

	type Interval struct {
		Start int
		End   int
	}

	var intervals []Interval
	var is int
	for _, e := range events {
		if e.ID == maxGuard {
			if e.event == EventState.Asleep {
				is = e.Minute
			} else if e.event == EventState.Awake {
				intervals = append(intervals, Interval{
					Start: is,
					End:   e.Minute,
				})
			}
		}
	}
	// fmt.Println(intervals)

	minuteMap := make(map[int]int)
	for i := 0; i < 60; i++ {
		for _, interval := range intervals {
			if interval.Start <= i && i < interval.End {
				c, ok := minuteMap[i]
				if !ok {
					minuteMap[i] = 1
				} else {
					minuteMap[i] = c + 1
				}
			}
		}
	}

	// fmt.Println(minuteMap)

	// find the largest minute here
	var maxMin, maxCount int
	for k, v := range minuteMap {
		if v > maxCount {
			maxCount = v
			maxMin = k
		}
	}
	fmt.Printf("guard %v slept the most at minute %v with count %v\n", maxGuard, maxMin, maxCount)
	fmt.Printf("guard x minute = %v\n", maxGuard*maxMin)
}

func part2() {
	fmt.Println("Day 04 Part 2")

	file, err := os.Open("input04.txt")
	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
		return
	}
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
			Time:   t,
			Minute: t.Minute(),
			event:  lineInfo,
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
	guardSleep := make(map[int]int)
	var start int
	var guard int
	for _, e := range events {
		if e.event == EventState.Asleep {
			e.ID = guard
			start = e.Minute
		} else if e.event == EventState.Awake {
			e.ID = guard

			sleep := e.Minute - start
			s, ok := guardSleep[guard]
			if !ok {
				guardSleep[guard] = sleep
			} else {
				guardSleep[guard] = s + sleep
			}
		} else {
			// new guard is begining
			guard = e.ID
		}
	}

	type Interval struct {
		ID    int
		Start int
		End   int
	}

	var intervals []Interval
	var is int
	for _, e := range events {
		if e.event == EventState.Asleep {
			is = e.Minute
		} else if e.event == EventState.Awake {
			intervals = append(intervals, Interval{
				ID:    e.ID,
				Start: is,
				End:   e.Minute,
			})
		}
	}
	// fmt.Println(intervals)

	guardIntervals := make(map[int][]Interval)
	for _, interval := range intervals {
		slice, ok := guardIntervals[interval.ID]
		if !ok {
			slice = []Interval{}
		}
		slice = append(slice, interval)
		guardIntervals[interval.ID] = slice
	}

	type GMC struct {
		Guard  int
		Minute int
		Count  int
	}

	var totals []GMC
	for g, ints := range guardIntervals {
		// for guard g
		minuteMap := make(map[int]int)
		for i := 0; i < 60; i++ {
			for _, interval := range ints {
				if interval.Start <= i && i < interval.End {
					c, ok := minuteMap[i]
					if !ok {
						minuteMap[i] = 1
					} else {
						minuteMap[i] = c + 1
					}
				}
			}
		}

		// find the largest minute here
		var maxMin, maxCount int
		for k, v := range minuteMap {
			if v > maxCount {
				maxCount = v
				maxMin = k
			}
		}

		totals = append(totals, GMC{
			Guard:  g,
			Minute: maxMin,
			Count:  maxCount,
		})
	}

	// find the highest count
	var solution GMC
	for _, gmc := range totals {
		if gmc.Count > solution.Count {
			solution = gmc
		}
	}

	// fmt.Println(minuteMap)

	fmt.Printf("guard %v slept the most at minute %v with count %v\n", solution.Guard, solution.Minute, solution.Count)
	fmt.Printf("guard x minute = %v\n", solution.Guard*solution.Minute)
}
