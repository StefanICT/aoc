package main

import (
	. "aoc/util"
	"fmt"
	"strings"
)

type Route struct {
	From string
	To   string
}

type Node struct {
	Cave   string
	Nodes  []*Node
	Parent *Node
}

var routes = []Route{}

func main() {
	for _, line := range ReadLines("2021/12/input") {
		routes = append(routes, Route{
			From: strings.Split(line, "-")[0],
			To:   strings.Split(line, "-")[1],
		})
	}

	part1()
	part2()
}

func part1() {
	start := Node{
		Cave:   "start",
		Nodes:  []*Node{},
		Parent: nil,
	}
	children(&start, routes, 1)

	fmt.Println(count(start))
}

func part2() {
	start := Node{
		Cave:   "start",
		Nodes:  []*Node{},
		Parent: nil,
	}
	children(&start, routes, 2)

	fmt.Println(count(start))
}

func children(node *Node, routes []Route, max int) {
	for _, route := range routes {
		if route.From == node.Cave || route.To == node.Cave {
			cave := route.From
			if route.From == node.Cave {
				cave = route.To
			}

			if cave == "start" {
				continue
			}

			allowed := strings.ToUpper(cave) == cave

			if !allowed {
				counts := map[string]int{}
				counts[node.Cave] = 1
				for parent := node.Parent; parent != nil; parent = parent.Parent {
					counts[parent.Cave]++
				}

				allowed = true

				if _, ok := counts[cave]; ok {
					for cave, count := range counts {
						if strings.ToLower(cave) == cave {
							if count >= max {
								allowed = false
								break
							}
						}
					}
				}
			}

			if allowed {
				child := Node{
					Cave:   cave,
					Parent: node,
				}
				node.Nodes = append(node.Nodes, &child)

				if node.Cave != "end" {
					children(&child, routes, max)
				}
			}
		}
	}
}

func count(node Node) int {
	if node.Cave == "end" {
		return 1
	}

	sum := 0
	for _, node := range node.Nodes {
		sum += count(*node)
	}

	return sum
}
