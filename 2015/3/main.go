package main

import (
	. "aoc/util"
	"fmt"
)

type Point struct {
	x int
	y int
}

func main() {
	part1()
	part2()
}

func part1() {
	current := Point{0, 0}
	data := map[Point]int{
		current: 1,
	}
	for _, line := range ReadLines("2015/3/input") {
		for i := 0; i < len(line); i++ {
			switch string(line[i]) {
			case "^":
				current.y -= 1
			case ">":
				current.x += 1
			case "v":
				current.y += 1
			case "<":
				current.x -= 1
			}
			data[current] = data[current] + 1
		}
	}
	fmt.Println("Day 3.1", len(data))
}

func part2() {
	santa := Point{0, 0}
	robot := Point{0, 0}
	current := &santa

	data := map[Point]int{
		*current: 1,
	}
	for _, line := range ReadLines("2015/3/input") {
		for i := 0; i < len(line); i++ {
			switch string(line[i]) {
			case "^":
				(*current).y -= 1
			case ">":
				(*current).x += 1
			case "v":
				(*current).y += 1
			case "<":
				(*current).x -= 1
			}
			data[*current] = data[*current] + 1
			if current == &santa {
				current = &robot
			} else {
				current = &santa
			}
		}
	}
	fmt.Println("Day 3.2", len(data))
}
