package main

import (
	. "aoc/util"
	"fmt"
)

type Node struct {
	Int    int
	X      *Node
	Y      *Node
	Parent *Node
}

func (node Node) Magnitude() int {
	if node.Int != -1 {
		return node.Int
	}
	return 3*node.X.Magnitude() + 2*node.Y.Magnitude()

}

func (node Node) String() string {
	if node.Int != -1 {
		return fmt.Sprintf("%d", node.Int)
	}
	return fmt.Sprintf("[%s,%s]", node.X.String(), node.Y.String())
}

func main() {
	part1()
	part2()
}

func part1() {
	lines := ReadLines("2021/18/input")
	nodes := make([]*Node, len(lines))
	for i := 0; i < len(nodes); i++ {
		nodes[i] = parse(lines[i])
	}

	sum := nodes[0]

	for i := 1; i < len(nodes); i++ {
		sum = Add(sum, nodes[i])
	}
	fmt.Println(sum.Magnitude())
}

func part2() {
	magnitude := 0
	lines := ReadLines("2021/18/input")
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines); j++ {
			a := parse(lines[i])
			b := parse(lines[j])
			magnitude = Max(Add(a, b).Magnitude(), magnitude)
		}
	}

	fmt.Println(magnitude)
}

func Add(a *Node, b *Node) *Node {
	root := &Node{
		Int: -1,
		X:   a,
		Y:   b,
	}
	a.Parent = root
	b.Parent = root

	Reduce(root)

	return root
}

func Reduce(node *Node) {
	for Explode(node, 4) || Split(node) {
	}
}

func Explode(node *Node, depth int) bool {
	if node.Int != -1 {
		return false
	}

	if depth == 0 {
		previous := node.X
		current := node
		for current != nil && (current.X == previous || current.X == nil) {
			previous = current
			current = current.Parent
		}

		if current != nil {
			left := current.X

			for left.Int == -1 {
				if left.Y != nil {
					left = left.Y
				} else {
					left = left.X
				}
			}

			left.Int += node.X.Int
		}

		previous = node.Y
		current = node
		for current != nil && (current.Y == previous || current.Y == nil) {
			previous = current
			current = current.Parent
		}

		if current != nil {
			right := current.Y

			for right.Int == -1 {
				if right.X != nil {
					right = right.X
				} else {
					right = right.Y
				}
			}

			right.Int += node.Y.Int
		}

		node.Int = 0

		return true
	}

	return Explode(node.X, depth-1) || Explode(node.Y, depth-1)
}

func Split(node *Node) bool {
	if node.Int != -1 {
		if node.Int >= 10 {
			node.X = &Node{Int: node.Int / 2, Parent: node}
			node.Y = &Node{Int: int(float64(node.Int)/2.0 + 0.5), Parent: node}
			node.Int = -1

			return true
		}

		return false
	}

	return Split(node.X) || Split(node.Y)
}

func Walk(node *Node) bool {
	var walk func(int, *Node) bool
	walk = func(depth int, node *Node) bool {
		if node.Int == -1 {
			if depth == 4 {
				// Explode

				node.Int = 0

				// Find first left and right.
				node.X = nil
				node.Y = nil

				return true
			}

			return walk(depth+1, node.X) || walk(depth+1, node.Y)
		}

		return false
	}

	return walk(0, node)
}

func parse(line string) *Node {
	index := 0

	var walk func(*Node) *Node
	walk = func(parent *Node) *Node {
		if line[index] == '[' {
			pair := &Node{
				Int:    -1,
				Parent: parent,
			}

			index += 1 // Consume [
			pair.X = walk(pair)
			index += 1 // Consume ,
			pair.Y = walk(pair)
			index += 1 // Consume ]

			return pair
		} else {
			val := int(line[index] - '0')
			index += 1 // Consume integer

			return &Node{
				Int:    val,
				Parent: parent,
			}
		}
	}

	return walk(nil)
}
