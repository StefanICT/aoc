package main

import (
	"fmt"
)

func main() {
	part1()
	part2()
}

func part1() {
	grid := [10][10]int{
		[10]int{1, 5, 5, 3, 4, 2, 1, 2, 8, 8},
		[10]int{5, 2, 5, 5, 3, 8, 4, 8, 8, 2},
		[10]int{1, 2, 2, 4, 3, 1, 5, 7, 3, 2},
		[10]int{4, 2, 5, 8, 2, 4, 2, 2, 7, 4},
		[10]int{1, 6, 5, 8, 5, 6, 4, 2, 1, 6},
		[10]int{6, 8, 7, 2, 6, 5, 1, 1, 8, 2},
		[10]int{5, 7, 7, 5, 5, 5, 2, 2, 3, 8},
		[10]int{5, 6, 2, 2, 5, 4, 5, 1, 7, 2},
		[10]int{8, 7, 6, 6, 6, 7, 2, 3, 1, 8},
		[10]int{2, 1, 7, 8, 3, 7, 4, 8, 3, 5},
	}
	sum := 0
	for i := 0; i < 400; i++ {
		step(&grid, &sum)
	}

	fmt.Println(sum)
}

func part2() {
	grid := [10][10]int{
		[10]int{1, 5, 5, 3, 4, 2, 1, 2, 8, 8},
		[10]int{5, 2, 5, 5, 3, 8, 4, 8, 8, 2},
		[10]int{1, 2, 2, 4, 3, 1, 5, 7, 3, 2},
		[10]int{4, 2, 5, 8, 2, 4, 2, 2, 7, 4},
		[10]int{1, 6, 5, 8, 5, 6, 4, 2, 1, 6},
		[10]int{6, 8, 7, 2, 6, 5, 1, 1, 8, 2},
		[10]int{5, 7, 7, 5, 5, 5, 2, 2, 3, 8},
		[10]int{5, 6, 2, 2, 5, 4, 5, 1, 7, 2},
		[10]int{8, 7, 6, 6, 6, 7, 2, 3, 1, 8},
		[10]int{2, 1, 7, 8, 3, 7, 4, 8, 3, 5},
	}
	sum := 0
steps:
	for i := 0; true; i++ {
		step(&grid, &sum)
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				if grid[i][j] != 0 {
					continue steps
				}
			}
		}

		fmt.Println(i + 1)
		break
	}
}

func step(grid *[10][10]int, sum *int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j] += 1
		}
	}

	flashed := map[string]bool{}

	check(grid, &flashed, sum)
}

func check(grid *[10][10]int, flashed *map[string]bool, sum *int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > 9 {
				flash(grid, i, j, flashed, sum)
			}
		}
	}
}

func flash(grid *[10][10]int, i int, j int, flashed *map[string]bool, sum *int) {
	id := fmt.Sprintf("%dx%d", i, j)
	if (*flashed)[id] {
		return
	}
	(*flashed)[id] = true
	*sum += 1

	(*grid)[i][j] = 0

	points := []struct{ I, J int }{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for _, point := range points {
		if i+point.I < 0 || i+point.I == len(grid) || j+point.J < 0 || j+point.J == len(grid[i+point.I]) {
			continue
		}

		id := fmt.Sprintf("%dx%d", i+point.I, j+point.J)
		if !(*flashed)[id] {
			(*grid)[i+point.I][j+point.J] += 1
		}
	}

	check(grid, flashed, sum)
}
