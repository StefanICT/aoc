package main

import (
	. "aoc/util"
	"fmt"
)

func main() {
	template, instructions := input()

	do(10, template, instructions)
	do(40, template, instructions)
}

func do(count int, template string, instructions map[string]byte) {
	counts := map[string]int{}

	for i := 1; i < len(template); i++ {
		counts[template[i-1:i+1]]++
	}

	for i := 0; i < count; i++ {
		step(&counts, instructions)
	}

	absolutes := map[byte]int{}
	for key, value := range counts {
		absolutes[key[0]] += value
	}
	// Include the last character of the template otherwise it will be
	// forgotten, and since we never add after the last character we can use
	// the template. Took me too long...
	absolutes[template[len(template)-1]]++

	min := MaxInt
	max := MinInt
	for _, count := range absolutes {
		min = Min(min, count)
		max = Max(max, count)
	}

	fmt.Println(max - min)
}

func step(counts *map[string]int, instructions map[string]byte) {
	copy := map[string]int{}
	for key, _ := range *counts {
		if match, ok := instructions[key]; ok {
			copy[string([]byte{key[0], match})] += (*counts)[key]
			copy[string([]byte{match, key[1]})] += (*counts)[key]
		}
	}
	*counts = copy
}

func input() (string, map[string]byte) {
	lines := ReadLines("2021/14/input")

	instructions := map[string]byte{}
	for i := 2; i < len(lines); i++ {
		instructions[lines[i][:2]] = lines[i][6]
	}

	return lines[0], instructions
}
