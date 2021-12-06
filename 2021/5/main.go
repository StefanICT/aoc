package main

import (
    . "aoc/utils"
    "regexp"
	"bufio"
	"fmt"
	"os"
)

type Line struct {
    start Point
    end Point
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
	file, _ := os.Open("5/test")
	defer file.Close()

    re := regexp.MustCompile(`(\d+),(\d)\s+->\s+(\d),(\d)`)

    lines := []Line{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
        match := re.FindStringSubmatch(scanner.Text())

        line := Line{
            start: Point {
                x: AsInt(match[1]),
                y: AsInt(match[2]),
            },
            end: Point {
                x: AsInt(match[3]),
                y: AsInt(match[4]),
            },
        }
        lines = append(lines, line)

	}

    fmt.Println(lines)
    fmt.Println(max(lines))
    debug(grid(max(lines)))

    fmt.Println("Day 5")
}

func part2() {
}

// Grid:
// >> X
// \/ Y
func grid(max Point) [][]int {
    grid := make([][]int, max.y + 1) 

    for i := 0; i < max.x; i++ {
        grid = append(grid, make([]int, max.x))
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

    return Point{ x, y }
}

func debug(grid [][]int) {
    fmt.Print("Grid:")
    for _, row := range grid {
        for _, column := range row {
            fmt.Print(column)
        }
        fmt.Print("\n")
    }
}
