package main

import (
	. "aoc/util"
	"fmt"
	"regexp"
	"strings"
)

type Fold struct {
	Axis   string
	Amount int
}

func main() {
	// part1()
	part2()
}

func part1() {
	grid, folds := input()
	fold(&grid, folds[0])
	fmt.Println(count(grid))
}

func part2() {
	grid, folds := input()
	for _, instruction := range folds {
		fold(&grid, instruction)
	}
	debug(grid)
}

func fold(grid *[][]bool, fold Fold) {
	switch fold.Axis {
	case "y":
		newY := fold.Amount - 1
		for y := fold.Amount + 1; y < len(*grid); y++ {
			for x := 0; x < len((*grid)[y]); x++ {
				(*grid)[newY][x] = (*grid)[y][x] || (*grid)[newY][x]
			}
			newY--
		}
		*grid = (*grid)[:fold.Amount]
	case "x":
		for y := 0; y < len(*grid); y++ {
			newX := fold.Amount - 1
			for x := fold.Amount + 1; x < len((*grid)[y]); x++ {
				(*grid)[y][newX] = (*grid)[y][newX] || (*grid)[y][x]
				newX--
			}

			(*grid)[y] = (*grid)[y][:fold.Amount]
		}
	}
}

func count(grid [][]bool) int {
	sum := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] {
				sum++
			}
		}
	}
	return sum
}

func debug(grid [][]bool) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			value := " "
			if grid[y][x] {
				value = "█"
			}

			fmt.Print(value)
		}
		fmt.Println()
	}
}

type Point struct {
	X int // >>
	Y int // \/
}

func input() ([][]bool, []Fold) {
	re := regexp.MustCompile(`([x|y])=(\d+)`)

	lines := ReadLines("2021/13/input")
	points := []Point{}
	folds := []Fold{}
	max := Point{}
	for _, line := range lines {
		if line == "" {
			continue
		}

		if strings.Contains(line, "fold") {
			matches := re.FindStringSubmatch(line)
			axis := matches[1]
			amount := AsInt(matches[2])

			folds = append(folds, Fold{Axis: axis, Amount: amount})
		} else {
			components := strings.Split(line, ",")

			point := Point{
				X: AsInt(components[0]),
				Y: AsInt(components[1]),
			}
			points = append(points, point)
			max = Point{X: Max(point.X, max.X), Y: Max(point.Y, max.Y)}
		}
	}

	grid := make([][]bool, max.Y+1)
	for y := 0; y <= max.Y; y++ {
		grid[y] = make([]bool, max.X+1)
	}

	for _, point := range points {
		grid[point.Y][point.X] = true
	}

	return grid, folds
}
