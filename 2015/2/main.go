package main

import (
	. "aoc/Util"
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	part1()
	part2()
}

func part1() {
	file, _ := os.Open("2015/2/input")
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ints := AsInts(scanner.Text(), "x")
		l := ints[0]
		w := ints[1]
		h := ints[2]

		area := 2*l*w + 2*w*h + 2*h*l

		sum += area + Min(l*w, w*h, h*l)
	}

	fmt.Printf("Day 1.1: %d\n", sum)
}

func part2() {
	file, _ := os.Open("2015/2/input")
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ints := AsInts(scanner.Text(), "x")
		sort.Ints(ints)
		l := ints[0]
		w := ints[1]
		h := ints[2]

		sum += Max(2*(l+w), 4*l)
		sum += l * w * h
	}

	fmt.Printf("Day 1.2: %d\n", sum)
}
