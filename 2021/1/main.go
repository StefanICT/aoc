package main

import (
	. "aoc/util"
	"fmt"
)

func main() {
	part1()
	part2()
}

func part1() {
	increased := 0
	previous := -1
	for _, value := range ReadInts("2021/1/input") {
		if previous != -1 && value > previous {
			increased++
		}

		previous = value
	}

	fmt.Printf("Day 1.1: %d\n", increased)
}

func part2() {
	increased := 0
	previous := -1
	ints := ReadInts("2021/1/input")
	for i, _ := range ints {
		if i+1 >= len(ints) || i+2 >= len(ints) {
			break
		}

		measurement := ints[i] + ints[i+1] + ints[i+2]
		if previous != -1 && measurement > previous {
			increased++
		}

		previous = measurement
	}

	fmt.Printf("Day 1.2: %d\n", increased)
}
