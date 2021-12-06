package main

import (
	. "aoc/util"
	"bufio"
	"fmt"
	"os"
)

func main() {
	part1()
	part2()
}

func part1() {
	fmt.Printf("Day 6.1: %d\n", count(read(), 80))
}

func part2() {
	fmt.Printf("Day 6.2: %d\n", count(read(), 256))
}

func read() []int {
	file, _ := os.Open("2021/6/input")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return AsInts(scanner.Text(), ",")
}

func count(ints []int, days int) int64 {
	data := map[int]int64{}
	for _, i := range ints {
		data[i] = int64(data[i] + 1)
	}

	for i := 0; i < days; i++ {
		zero := data[0]
		one := data[1]
		two := data[2]
		three := data[3]
		four := data[4]
		five := data[5]
		six := data[6]
		seven := data[7]
		eight := data[8]

		data[0] = one
		data[1] = two
		data[2] = three
		data[3] = four
		data[4] = five
		data[5] = six
		data[6] = seven + zero
		data[7] = eight
		data[8] = zero
	}

	sum := int64(0)
	for _, value := range data {
		sum += value
	}

	return sum
}
