package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "19"
	dayTitle = "Not Enough Minerals"
)

const (
	OreIdx = 0
	ClyIdx = 1
	ObsIdx = 2
	GdeIdx = 3
)

type Cost struct {
	Ore, Clay, Obsidian int
}

type BluePrint struct {
	ID                                         int
	OreCost, ClayCost, ObsidianCost, GeodeCost Cost
	MaxOre, MaxClay, MaxObsidian               int
}

type State struct {
	OreResource, ClayResource, ObsidianResource, GeodeResource int
	OreRobot, ClayRobot, ObsidianRobot, GeodeRobot             int
	Minutes                                                    int
	Build                                                      [4]bool
}

func (s State) CanBuild(c Cost) (State, bool) {
	if s.OreResource >= c.Ore && s.ClayResource >= c.Clay && s.ObsidianResource >= c.Obsidian {
		return State{
			s.OreResource - c.Ore,
			s.ClayResource - c.Clay,
			s.ObsidianResource - c.Obsidian,
			s.GeodeResource,
			s.OreRobot, s.ClayRobot, s.ObsidianRobot, s.GeodeRobot, s.Minutes, s.Build,
		}, true
	}
	return s, false
}

func (s State) Produce() State {
	return State{
		s.OreResource + s.OreRobot,
		s.ClayResource + s.ClayRobot,
		s.ObsidianResource + s.ObsidianRobot,
		s.GeodeResource + s.GeodeRobot,
		s.OreRobot, s.ClayRobot, s.ObsidianRobot, s.GeodeRobot,
		s.Minutes - 1, s.Build,
	}
}

func (s State) AddRobot(ore, clay, obsidian, geode int) State {
	return State{
		s.OreResource, s.ClayResource, s.ObsidianResource, s.GeodeResource,
		s.OreRobot + ore,
		s.ClayRobot + clay,
		s.ObsidianRobot + obsidian,
		s.GeodeRobot + geode,
		s.Minutes, s.Build,
	}
}

func (s State) SaveNoBuild(ore, clay, obsidian, geode bool) State {
	return State{
		s.OreResource, s.ClayResource, s.ObsidianResource, s.GeodeResource,
		s.OreRobot, s.ClayRobot, s.ObsidianRobot, s.GeodeRobot,
		s.Minutes,
		[4]bool{ore, clay, obsidian, geode},
	}
}

func NewInitialState(bp BluePrint, min int) State {
	return State{
		0, 0, 0, 0, // resources
		1, 0, 0, 0, // robots
		min, // minutes
		[4]bool{true, true, true, true},
	}
}

func ParseBluePrint(line string) BluePrint {
	fields := strings.Fields(line)

	fid := fields[1]
	fid = fid[:len(fid)-1]
	id, _ := strconv.Atoi(fid)

	maxOre := 0

	// Ore cost
	oreOre, _ := strconv.Atoi(fields[6])
	ore := Cost{oreOre, 0, 0}
	maxOre = oreOre

	// Clay cost
	clayOre, _ := strconv.Atoi(fields[12])
	clay := Cost{clayOre, 0, 0}
	if clayOre > maxOre {
		maxOre = clayOre
	}

	// Obsidian cost
	obsidianOre, _ := strconv.Atoi(fields[18])
	obsidianClay, _ := strconv.Atoi(fields[21])
	obsidian := Cost{obsidianOre, obsidianClay, 0}
	if obsidianOre > maxOre {
		maxOre = obsidianOre
	}

	// Geode cost
	geodeOre, _ := strconv.Atoi(fields[27])
	geodeObsidian, _ := strconv.Atoi(fields[30])
	geode := Cost{geodeOre, 0, geodeObsidian}
	if geodeOre > maxOre {
		maxOre = geodeOre
	}

	return BluePrint{
		id,
		ore, clay, obsidian, geode,
		maxOre, obsidianClay, geodeObsidian,
	}
}

func Parse(filename string) []BluePrint {
	input, _ := util.ReadInput(filename)
	blueprints := make([]BluePrint, 0, len(input))
	for _, line := range input {
		blueprints = append(blueprints, ParseBluePrint(line))
	}
	return blueprints
}

func BuildOrders(bp BluePrint, state State) State {
	if state.Minutes <= 0 {
		return state
	}

	// build geode robot - this is always the best move!! ??
	if s, ok := state.CanBuild(bp.GeodeCost); ok {
		s = s.Produce().AddRobot(0, 0, 0, 1).SaveNoBuild(true, true, true, true)
		ss := BuildOrders(bp, s)
		return ss
	}

	results := make([]State, 0, 4)
	best := 0

	// build ore robot
	oreState, oreOk := state.CanBuild(bp.OreCost)
	if oreOk && state.OreRobot < bp.MaxOre && state.Build[OreIdx] {
		oreState = oreState.Produce().AddRobot(1, 0, 0, 0).SaveNoBuild(true, true, true, true)
		ore := BuildOrders(bp, oreState)
		results = append(results, ore)
		best = util.Max(best, ore.GeodeResource)
	}
	// build clay robot
	clayState, clayOk := state.CanBuild(bp.ClayCost)
	if clayOk && state.ClayRobot < bp.MaxClay && state.Build[ClyIdx] {
		clayState = clayState.Produce().AddRobot(0, 1, 0, 0).SaveNoBuild(true, true, true, true)
		clay := BuildOrders(bp, clayState)
		results = append(results, clay)
		best = util.Max(best, clay.GeodeResource)
	}
	// build obsidian robot
	obsState, obsOk := state.CanBuild(bp.ObsidianCost)
	if obsOk && state.ObsidianRobot < bp.MaxObsidian && state.Build[ObsIdx] {
		obsState = obsState.Produce().AddRobot(0, 0, 1, 0).SaveNoBuild(true, true, true, true)
		obs := BuildOrders(bp, obsState)
		results = append(results, obs)
		best = util.Max(best, obs.GeodeResource)
	}
	// save up, don't build what we could have built here
	saveState := state.Produce().SaveNoBuild(!oreOk, !clayOk, !obsOk, true)
	save := BuildOrders(bp, saveState)
	results = append(results, save)
	best = util.Max(best, save.GeodeResource)

	// pick best result to return // best is highest geode count
	bestState := results[0]
	for i := 1; i < len(results); i++ {
		if results[i].GeodeResource > bestState.GeodeResource {
			bestState = results[i]
		}
	}
	return bestState
}

func part1() {
	blueprints := Parse("input.txt")
	qualitySum := 0
	for _, bp := range blueprints {
		s := BuildOrders(bp, NewInitialState(bp, 24))
		quality := bp.ID * s.GeodeResource
		qualitySum += quality
		fmt.Printf("Blueprint %v: Geode: %v Quality: %v Sum: %v\n", bp.ID, s.GeodeResource, quality, qualitySum)
	}
	fmt.Printf("Part 1: %v\n", qualitySum)
}

func part2() {
	blueprints := Parse("input.txt")
	bps := blueprints[:util.Min(3, len(blueprints))]
	geodeProd := 1
	for _, bp := range bps {
		s := BuildOrders(bp, NewInitialState(bp, 32))
		geodeProd *= s.GeodeResource
		fmt.Printf("Blueprint %v: Geode: %v Product: %v\n", bp.ID, s.GeodeResource, geodeProd)
	}
	fmt.Printf("Part 1: %v\n", geodeProd)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
