package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	part1()
	part2()
}

func part1() {
	file, _ := os.Open("3/report")
	defer file.Close()

	slice := make([]uint64, 12)
	lines := 0

	scanner := bufio.NewScanner(file)
	for ; scanner.Scan(); lines++ {
		n, _ := strconv.ParseUint(scanner.Text(), 2, 64)

		for i := 0; n != 0; n >>= 1 {
			slice[i] += n & 1
			i++
		}
	}

	breakpoint := uint64(lines / 2)

	gamma := 0
	epsilon := 0

	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] < breakpoint {
			epsilon |= 1 << i
		} else {
			gamma |= 1 << i
		}
	}

	fmt.Printf("Day 3.1: %d\n", gamma*epsilon)
}

func part2() {
	file, _ := os.Open("3/report")
	defer file.Close()

	length := 12
	numbers := make([]uint64, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, _ := strconv.ParseUint(scanner.Text(), 2, 64)
		numbers = append(numbers, n)
	}

	oxygens := numbers
	co2s := numbers

	for i := length - 1; i >= 0; i-- {
		if len(oxygens) > 1 {
			target := uint8(0)
			if float64(common(oxygens, i)) >= math.Round(float64(len(oxygens))/2.0) {
				target = 1
			}

			oxygens = filter(oxygens, i, target)
		}

		if len(co2s) > 1 {
			target := uint8(1)
			if float64(common(co2s, i)) >= math.Round(float64(len(co2s))/2.0) {
				target = 0
			}

			co2s = filter(co2s, i, target)
		}
	}

	fmt.Println("Day 3.2:", oxygens[0]*co2s[0])
}

func common(numbers []uint64, i int) uint64 {
	n := uint64(0)
	for _, number := range numbers {
		n += number >> i & 1
	}
	return n
}

func filter(numbers []uint64, i int, target uint8) []uint64 {
	slice := []uint64{}
	for _, n := range numbers {
		if uint8(n>>i&1) == target {
			slice = append(slice, n)
		}

	}

	return slice
}
