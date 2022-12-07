package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "07"
	dayTitle = "No Space Left On Device"
)

type File struct {
	name     string
	size     int
	isFile   bool
	parent   *File
	children map[string]*File
}

func NewFile(name string, size int, parent *File) *File {
	return &File{name, size, true, parent, nil}
}

func NewDir(name string, parent *File) *File {
	return &File{name, 0, false, parent, make(map[string]*File)}
}

func (f *File) AddFile(add *File) {
	f.children[add.name] = add
}

type FileSystem struct {
	root *File
	curr *File
}

func NewFS() FileSystem {
	root := NewDir("/", nil)
	return FileSystem{root, root}
}

func parse(filename string) FileSystem {
	input, _ := util.ReadInput(filename)
	fs := NewFS()
	var cmd string
	for _, line := range input {
		fields := strings.Fields(line)
		if fields[0] == "$" {
			cmd = fields[1]

			if cmd == "cd" {
				// handle built-ins for cd
				if fields[2] == ".." {
					fs.curr = fs.curr.parent
				} else if fields[2] == "/" {
					fs.curr = fs.root
				} else {
					fs.curr = fs.curr.children[fields[2]]
				}
			}
			continue
		}

		// must be the output of ls
		if fields[0] == "dir" {
			dir := NewDir(fields[1], fs.curr)
			fs.curr.AddFile(dir)
			continue
		}
		s, _ := strconv.Atoi(fields[0])
		f := NewFile(fields[1], s, fs.curr)
		fs.curr.AddFile(f)
	}
	return fs
}

func sizeof(file *File, memo map[*File]int) map[*File]int {
	if _, ok := memo[file]; ok {
		return memo
	}

	if file.isFile {
		memo[file] = file.size
		return memo
	}

	s := 0
	for _, f := range file.children {
		memo = sizeof(f, memo)
		s += memo[f]
	}
	memo[file] = s
	return memo
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	fs := parse("input.txt")
	memo := make(map[*File]int)
	memo = sizeof(fs.root, memo)

	dirSize := 0
	for f, s := range memo {
		if !f.isFile && s <= 100000 {
			dirSize += s
		}
	}
	fmt.Printf("Part 1: %v\n", dirSize)

	unused := 70000000 - memo[fs.root]
	need := 30000000
	target := need - unused
	minSize := memo[fs.root]
	for f, s := range memo {
		if !f.isFile && s >= target && s < minSize {
			minSize = s
		}
	}
	fmt.Printf("Part 2: %v\n", minSize)
}
