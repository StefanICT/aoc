package main

import (
	. "aoc/util"
	"fmt"
	"sort"
)

func main() {
	part1()
	part2()
}

func part1() {
	grid := ReadLines("2021/9/input")
	risk := 0
	lowpoints(grid, func(i int, j int) {
		risk += int(grid[i][j]-'0') + 1
	})
	fmt.Println(risk)
}

func part2() {
	grid := ReadLines("2021/9/input")
	basins := []int{}
	lowpoints(grid, func(i int, j int) {
		points := map[string]int{}
		walk(i, j, grid, &points)
		count := len(points)

		if len(basins) < 3 {
			basins = append(basins, count)
		} else if basins[0] < count {
			basins[0] = count
		}
		sort.Ints(basins)
	})

	total := 1
	for _, basin := range basins {
		total *= basin
	}
	fmt.Println(total)
}

func lowpoints(grid []string, fn func(i int, j int)) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if (grid[i][j] > grid[i][Min(j+1, len(grid[i])-1)]) ||
				(grid[i][j] > grid[i][Max(0, j-1)]) ||
				(grid[i][j] > grid[Max(0, i-1)][j]) ||
				(grid[i][j] > grid[Min(i+1, len(grid)-1)][j]) {
				continue
			}

			fn(i, j)
		}
	}
}

func walk(i int, j int, grid []string, points *map[string]int) {
	id := fmt.Sprintf("%dx%d", i, j)
	if _, exists := (*points)[id]; exists || grid[i][j] == '9' {
		return
	}

	(*points)[id] = 0

	if grid[i][j] < grid[i][Min(j+1, len(grid[i])-1)] {
		walk(i, j+1, grid, points)
	}

	if grid[i][j] < grid[i][Max(0, j-1)] {
		walk(i, j-1, grid, points)
	}

	if grid[i][j] < grid[Max(0, i-1)][j] {
		walk(i-1, j, grid, points)
	}

	if grid[i][j] < grid[Min(i+1, len(grid)-1)][j] {
		walk(i+1, j, grid, points)
	}
}
