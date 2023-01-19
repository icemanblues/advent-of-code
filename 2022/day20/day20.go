package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "20"
	dayTitle = "Grove Positioning System"
)

type Node struct {
	num        int
	prev, next *Node
}

func Move(n *Node, e, l int) *Node {
	move := util.Abs(e) % (l - 1)
	for i := 0; i < move; i++ {
		if e > 0 {
			n = n.next
		} else {
			n = n.prev
		}
	}
	return n
}

type Circular struct {
	Original []int
	Length   int
	Nodes    []*Node
	Zero     *Node
}

func (c Circular) Print(idx, l int) {
	if idx < 0 || idx >= l {
		return
	}
	curr := c.Nodes[idx]
	for i := 0; i < l; i++ {
		fmt.Printf("%v ", curr.num)
		curr = curr.next
	}
	fmt.Println()
}

func Parse(filename string, key int) Circular {
	input, _ := util.ReadIntput(filename)
	l := len(input)
	orig := make([]int, l, l)
	orig[0] = input[0] * key
	nodes := make([]*Node, l, l)
	head := &Node{orig[0], nil, nil}
	nodes[0] = head
	zero := head
	curr := head
	for i := 1; i < l; i++ {
		orig[i] = input[i] * key
		n := &Node{orig[i], curr, nil}
		nodes[i] = n
		curr.next = n
		curr = n
		if input[i] == 0 {
			zero = n
		}
	}
	curr.next = head
	head.prev = curr
	return Circular{orig, l, nodes, zero}
}

func (c *Circular) Mix() {
	for i, e := range c.Original {
		if e == 0 {
			continue
		}
		node := c.Nodes[i]
		curr := Move(node, e, c.Length)

		// unlink node
		nprev := node.prev
		nnext := node.next
		nprev.next = nnext
		nnext.prev = nprev

		// curr is the target, where it wants to go
		if e > 0 { // go after it
			next := curr.next
			curr.next = node
			node.prev = curr
			node.next = next
			next.prev = node
		} else { // go before it
			prev := curr.prev
			curr.prev = node
			node.next = curr
			node.prev = prev
			prev.next = node
		}
	}
}

func (c *Circular) GPS() int {
	sum := 0
	for i := 1; i <= 3; i++ {
		sum += Move(c.Zero, i*1000, c.Length).num
	}
	return sum
}

func part1() {
	circ := Parse("input.txt", 1)
	circ.Mix()
	fmt.Printf("Part 1: %v\n", circ.GPS())
}

const decrypt = 811589153

// 3748730297707 too high
func part2() {
	circ := Parse("input.txt", decrypt)
	for i := 0; i < 10; i++ {
		circ.Mix()
	}
	fmt.Printf("Part 2: %v\n", circ.GPS())
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
