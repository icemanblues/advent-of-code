package main

import (
	"testing"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const TestFile string = "test1.txt"

var Input []string
var Tiles []Tile

func refTile() Tile {
	return Tile{20,
		[][]rune{
			{'1', '2', '3'},
			{'4', '5', '6'},
			{'7', '8', '9'},
		},
	}
}

func setup() {
	var err error
	Input, err = util.ReadInput(TestFile)
	if err != nil {
		panic("Unable to read input from file: " + TestFile)
	}
	Tiles = ParseTile(Input)
}

func TestMain(m *testing.M) {
	setup()
	m.Run()
}

func TestParseTiles(t *testing.T) {
	expectedTileIds := []int{
		2311, 1951, 1171, 1427, 1489, 2473, 2971, 2729, 3079,
	}
	expectedLen := 9

	tiles := ParseTile(Input)
	if l := len(tiles); l != expectedLen {
		t.Errorf("Should have parsed %v tiles. actual %v", expectedLen, l)
	}
	for i, expID := range expectedTileIds {
		if expID != tiles[i].ID {
			t.Errorf("Tile %v expecting %v. Actual %v", i, expID, tiles[i].ID)
		}
	}
}

func TestNewGrid(t *testing.T) {
	tests := []struct {
		name string
		n    int
	}{
		{"one", 1},
		{"three", 3},
		{"nine", 9},
		{"zero", 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			grid := newGrid(test.n)
			if l := len(grid); l != test.n {
				t.Errorf("Number of columns should be n. Expected %v Actual %v", test.n, l)
			}
			for i := range grid {
				row := grid[i]
				if l := len(row); l != test.n {
					t.Errorf("row %v be n. Expected %v Actual %v", i, test.n, l)
				}
			}
		})
	}
}

func equalGrid(a, b [][]rune) bool {
	for j := range a {
		for i := range a[j] {
			if a[j][i] != b[j][i] {
				return false
			}
		}
	}
	return true
}

func TestFlip(t *testing.T) {
	expected := [][]rune{
		{'3', '2', '1'},
		{'6', '5', '4'},
		{'9', '8', '7'},
	}
	f := flip(refTile())
	if !equalGrid(f.Grid, expected) {
		t.Errorf("Expected:\n%v\nActual:\n%v", expected, f.Grid)
	}
}

func TestRotate90(t *testing.T) {
	expected := [][]rune{
		{'7', '4', '1'},
		{'8', '5', '2'},
		{'9', '6', '3'},
	}
	f := rotate90(refTile())
	if !equalGrid(f.Grid, expected) {
		t.Errorf("Expected:\n%v\nActual:\n%v", expected, f.Grid)
	}
}

func TestAllRotations(t *testing.T) {
	expected := [][][]rune{
		{
			{'1', '2', '3'},
			{'4', '5', '6'},
			{'7', '8', '9'},
		},
		{
			{'7', '4', '1'},
			{'8', '5', '2'},
			{'9', '6', '3'},
		},
		{
			{'9', '8', '7'},
			{'6', '5', '4'},
			{'3', '2', '1'},
		},
		{
			{'3', '6', '9'},
			{'2', '5', '8'},
			{'1', '4', '7'},
		},
	}
	rotates := allRotations(refTile())
	for i, rot := range rotates {
		if !equalGrid(rot.Grid, expected[i]) {
			t.Errorf("rotation %v: Expected:\n%v\nActual:\n%v", i, expected[i], rot.Grid)
		}
	}
}

func TestRemoveBorder(t *testing.T) {
	expected := [][]rune{{'5'}}
	noBorder := removeBorder(refTile())
	if !equalGrid(noBorder.Grid, expected) {
		t.Errorf("incorrect border removed: Expected:\n%v\nActual:\n%v", expected, noBorder.Grid)
	}
}

func TestStitchImage(t *testing.T) {
	expected := [][]rune{
		[]rune("12"),
		[]rune("34"),
	}
	one := Tile{1, [][]rune{{'1'}}}
	two := Tile{2, [][]rune{{'2'}}}
	thr := Tile{3, [][]rune{{'3'}}}
	fou := Tile{4, [][]rune{{'4'}}}
	tiles := []Tile{one, two, thr, fou}
	stitch := stitchImage(tiles, 2)
	if !equalGrid(stitch.Grid, expected) {
		t.Errorf("stich not right: Expected:\n%v\nActual:\n%v", expected, stitch.Grid)
	}
}

// TODO: Make this a table test
func TestStitchImageBigger(t *testing.T) {
	expected := [][]rune{
		[]rune("1234"),
		[]rune("5678"),
		[]rune("abcd"),
		[]rune("efgh"),
	}

	one := Tile{1, [][]rune{
		[]rune("12"),
		[]rune("56"),
	}}
	two := Tile{2, [][]rune{
		[]rune("34"),
		[]rune("78"),
	}}
	thr := Tile{3, [][]rune{
		[]rune("ab"),
		[]rune("ef"),
	}}
	fou := Tile{4, [][]rune{
		[]rune("cd"),
		[]rune("gh"),
	}}

	tiles := []Tile{one, two, thr, fou}
	stitch := stitchImage(tiles, 2)

	if !equalGrid(stitch.Grid, expected) {
		t.Errorf("stich not right: Expected:\n%v\nActual:\n%v", expected, stitch.Grid)
	}
}

func TestCheckTopToBot(t *testing.T) {
	topTile := Tile{1, [][]rune{
		[]rune("###"),
		[]rune("###"),
		[]rune("..#"),
	}}
	botTile := Tile{2, [][]rune{
		[]rune("..#"),
		[]rune("#.."),
		[]rune("#.#"),
	}}
	noTile := Tile{-1, [][]rune{
		[]rune("..."),
		[]rune("..."),
		[]rune("..."),
	}}

	if ok := checkTopToBot(botTile, topTile); !ok {
		t.Errorf("This solution should work!")
	}
	if ok := checkTopToBot(noTile, topTile); ok {
		t.Errorf("This solution should NOT work!")
	}
}

func TestCheckLeftToRight(t *testing.T) {
	leftTile := Tile{1, [][]rune{
		[]rune("##."),
		[]rune("##."),
		[]rune("###"),
	}}
	rightTile := Tile{2, [][]rune{
		[]rune(".##"),
		[]rune(".##"),
		[]rune("###"),
	}}
	noTile := Tile{-1, [][]rune{
		[]rune("..."),
		[]rune("..."),
		[]rune("..."),
	}}

	if ok := checkLeftToRight(rightTile, leftTile); !ok {
		t.Errorf("This solution should work!")
	}
	if ok := checkLeftToRight(noTile, leftTile); ok {
		t.Errorf("This solution should NOT work!")
	}
}

func TestCheckBorder(t *testing.T) {
}

func TestIsMonster(t *testing.T) {
}

func TestFindConnections(t *testing.T) {
	corners := map[int]struct{}{
		1951: {},
		3079: {},
		2971: {},
		1171: {},
	}
	borders := map[int]struct{}{
		2311: {},
		2729: {},
		2473: {},
		1489: {},
	}
	middles := map[int]struct{}{
		1427: {},
	}

	check := func(t *testing.T, tiles []Tile, tileIds map[int]struct{}) {
		if len(tiles) != len(tileIds) {
			t.Errorf("Incorrect number of tiles: %v %v", len(tiles), len(tileIds))
		}
		for _, tile := range tiles {
			if _, ok := tileIds[tile.ID]; !ok {
				t.Errorf("Incorrect Tile ID found: %v in set %v", tile.ID, tileIds)
			}
		}
	}

	numSidesToTiles := FindConnections(Tiles)
	for sides, tiles := range numSidesToTiles {
		switch sides {
		case 2: // corners
			check(t, tiles, corners)
		case 3: // borders
			check(t, tiles, borders)
		case 4: // middle
			check(t, tiles, middles)
		default:
			t.Errorf("What!? %v %v", sides, tiles)
		}
	}
}

func TestSearch(t *testing.T) {
}
