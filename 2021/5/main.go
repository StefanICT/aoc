package main

import (
	. "aoc/util"
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Line struct {
	start Point
	end   Point
}

func (line Line) isDiagonal() bool {
	return line.start.x != line.end.x && line.start.y != line.end.y
}

type Point struct {
	x int
	y int
}

func main() {
	part1()
	part2()
}

func part1() {
	lines := makeLines()
	grid := makeGrid(max(lines))

	for _, line := range lines {
		if line.isDiagonal() {
			continue
		}

		plot(line, grid)
	}

	fmt.Println("Day 5.1:", count(grid, 2))
}

func part2() {
	lines := makeLines()
	grid := makeGrid(max(lines))

	for _, line := range lines {
		plot(line, grid)
	}

	fmt.Println("Day 5.1:", count(grid, 2))
}

func makeLines() []Line {
	file, _ := os.Open("2021/5/input")
	defer file.Close()

	re := regexp.MustCompile(`(\d+),(\d+)\s+->\s+(\d+),(\d+)`)

	lines := []Line{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		match := re.FindStringSubmatch(scanner.Text())

		line := Line{
			start: Point{
				x: AsInt(match[1]),
				y: AsInt(match[2]),
			},
			end: Point{
				x: AsInt(match[3]),
				y: AsInt(match[4]),
			},
		}
		lines = append(lines, line)

	}

	return lines
}

func makeGrid(max Point) [][]int {
	grid := make([][]int, max.x+1)

	for i := 0; i < max.y+1; i++ {
		grid[i] = make([]int, max.x+1)
	}

	return grid
}

func max(lines []Line) Point {
	x := 0
	y := 0

	for _, line := range lines {
		x = Max(x, line.start.x)
		y = Max(y, line.start.y)

		x = Max(x, line.end.x)
		y = Max(y, line.end.y)
	}

	return Point{x, y}
}

func plot(line Line, grid [][]int) {
	// Use Bresenham's line algorithm
	// https://medium.com/geekculture/bresenhams-line-drawing-algorithm-2e0e953901b3

	dx := Abs(line.end.x - line.start.x)
	dy := Abs(line.end.y - line.start.y)

	var sx int
	if line.start.x < line.end.x {
		sx = 1
	} else {
		sx = -1
	}

	var sy int
	if line.start.y < line.end.y {
		sy = 1
	} else {
		sy = -1
	}

	d := dx - dy

	for {
		grid[line.start.y][line.start.x] += 1
		if line.start == line.end {
			return
		}

		d2 := 2 * d
		if d2 > (0 - dy) {
			d -= dy
			line.start.x += sx
		}
		if d2 < dx {
			d += dx
			line.start.y += sy
		}
	}
}

func count(grid [][]int, value int) int {
	sum := 0
	for _, row := range grid {
		for _, column := range row {
			if column >= value {
				sum += 1
			}
		}
	}
	return sum
}
