package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
    part1()
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

	fmt.Printf("Day 3: %d\n", gamma*epsilon)
}
