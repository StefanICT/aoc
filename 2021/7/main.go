package main

import (
	. "aoc/util"
    "math"
	"fmt"
)

func main() {
	part1()
	part2()
}

func part1() {
    ints := AsInts(ReadLines("2021/7/input")[0], ",")

    min := math.MaxInt64
    for i := 0; i < len(ints); i++ {
        fuel := 0
        for _, j := range ints {
            fuel += Abs(j - i)
        }
        min = Min(fuel, min)
    }

	fmt.Printf("Day 6.1: %d\n", min)
}

func part2() {
    ints := AsInts(ReadLines("2021/7/input")[0], ",")

    min := math.MaxInt64
    for i := 0; i < len(ints); i++ {
        fuel := 0
        for _, j := range ints {
            steps := Abs(j - i)
            for z := 1; z <= steps; z++ {
                fuel += z
            }
        }
        min = Min(fuel, min)
    }

	fmt.Printf("Day 6.2: %d\n", min)
}
