package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	part1()
	part2()
}

func part1() {
	file, _ := os.Open("1/input")
	defer file.Close()

	floor := 0
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		switch scanner.Text() {
		case "(":
			floor++
		case ")":
			floor--
		}
	}

	fmt.Printf("Day 1.1: %d\n", floor)
}

func part2() {
	file, _ := os.Open("1/input")
	defer file.Close()

	floor := 0
    ch := 1
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		switch scanner.Text() {
		case "(":
			floor++
		case ")":
			floor--
		}

        if floor < 0 {
            break
        }
        ch++
	}

	fmt.Printf("Day 1.2: %d\n", ch)
}
