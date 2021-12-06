package main

import (
	. "aoc/util"
	"fmt"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	horizontal := 0
	depth := 0

	for _, line := range ReadLines("2021/2/input") {
		components := strings.Split(line, " ")
		action := components[0]
		number := AsInt(components[1])

		switch action {
		case "forward":
			horizontal += number
		case "down":
			depth += number
		case "up":
			depth -= number
		}
	}

	fmt.Printf("Day 2.1: %d\n", horizontal*depth)
}

func part2() {
	horizontal := 0
	depth := 0
	aim := 0

	for _, line := range ReadLines("2021/2/input") {
		components := strings.Split(line, " ")
		action := components[0]
		number := AsInt(components[1])

		switch action {
		case "forward":
			horizontal += number
			depth += aim * number
		case "down":
			aim += number
		case "up":
			aim -= number
		}
	}

	fmt.Printf("Day 2.2: %d\n", horizontal*depth)
}
