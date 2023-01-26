package main

import (
	"fmt"
	"strconv"
)

const (
	dayNum   = "18"
	dayTitle = "Snailfish"
)

type Number struct {
	Value       int
	Left, Right *Number
	Parent      *Number
}

// IsLeaf true if this Number "Node" is a regular number
func (n *Number) IsLeaf() bool {
	return n.Left == nil && n.Right == nil
}

func (n *Number) IsRoot() bool {
	return n.Parent == nil
}

// String converts a Snailfish number to a printable string
func (n *Number) String() string {
	if n.IsLeaf() {
		return strconv.Itoa(n.Value)
	}
	return fmt.Sprintf("[%v,%v]", n.Left.String(), n.Right.String())
}

func (n *Number) Magnitude() int {
	if n.IsLeaf() {
		return n.Value
	}
	return 3*n.Left.Magnitude() + 2*n.Right.Magnitude()
}

func NewNumber(a int) *Number {
	return &Number{a, nil, nil, nil}
}

func NewPair(a, b int) *Number {
	an, bn := NewNumber(a), NewNumber(b)
	n := &Number{0, an, bn, nil}
	an.Parent, bn.Parent = n, n
	return n
}

func Pair(a, b *Number) *Number {
	n := &Number{0, a, b, nil}
	a.Parent, b.Parent = n, n
	return n
}

func ParseNumber(s string) *Number {
	n, _ := ParseRunes([]rune(s), 0)
	return n
}

func ParseRunes(r []rune, i int) (*Number, int) {
	root := &Number{0, nil, nil, nil}

	for i < len(r) {
		switch r[i] {
		case '[':
			number, j := ParseRunes(r, i+1)
			root.Left = number
			number.Parent = root
			i = j
		case ']':
			return root, i + 1
		case ',':
			number, j := ParseRunes(r, i+1)
			root.Right = number
			number.Parent = root
			i = j
		default: // must be a number
			num, err := strconv.Atoi(string(r[i]))
			if err != nil {
				fmt.Printf("Not a number: %v at index: %v\n", r[i], i)
				return root, len(r)
			}
			return NewNumber(num), i + 1
		}
	}
	return root, len(r)
}

func AddNumbers(a, b *Number) *Number {
	n := &Number{0, a, b, nil}
	a.Parent, b.Parent = n, n
	return n
}

func (n *Number) Reduce() {
	// keep reducing until no action has occurred
	// explode
	// split
}

type NumberPredicate func(n *Number, depth int) bool

var IsExplode NumberPredicate = func(n *Number, depth int) bool {
	return !n.IsLeaf() && depth == 4
}

var IsSplit NumberPredicate = func(n *Number, depth int) bool {
	return n.IsLeaf() && n.Value >= 10
}

func OrderedSlice(root *Number) []*Number {
	var slice []*Number

	if root.IsLeaf() {
		slice = append(slice, root)
		return slice
	}

	slice = append(slice, OrderedSlice(root.Left)...)
	slice = append(slice, root)
	slice = append(slice, OrderedSlice(root.Right)...)
	return slice
}

func Find(root *Number, depth int, l bool, predicate NumberPredicate) *Number {
	if predicate(root, depth) {
		return root
	}

	first := root.Right
	second := root.Left
	if l {
		first = root.Left
		second = root.Right
	}

	n := Find(first, depth+1, l, predicate)
	if n != nil {
		return n
	}
	n = Find(second, depth+1, l, predicate)
	if n != nil {
		return n
	}

	// check the parent on the correct side?
	return nil
}

func part1() {
	fmt.Println("Part 1")
	//input, _ := util.ReadInput("input.txt")
}

func part2() {
	fmt.Println("Part 2")
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()

	root := ParseNumber("[[1,2],[[3,4],5]]")
	slice := OrderedSlice(root)
	for i, n := range slice {
		fmt.Printf("index: %v: %v\n", i, n)
	}
}
