package main

import (
	. "aoc/util"
	"fmt"
)

type Point struct {
	X int
	Y int
}

func main() {
	x1 := 288
	x2 := 330
	y1 := -96
	y2 := -50

	fmt.Println(part1(x1, x2, y1, y2))
	fmt.Println(part2(x1, x2, y1, y2))
}

func part1(x1 int, x2 int, y1 int, y2 int) int {
	highestY := 0
	for x := 0; x <= x2; x++ {
		for y := y1; Abs(y) <= Abs(y1); y++ {
			initialVelocity := Point{x, y}
			if points, hit := Trajection(initialVelocity, x1, x2, y1, y2); hit {
				highestY = Max(HighestY(points), highestY)
			}
		}
	}
	return highestY
}

func HighestY(points []Point) int {
	highestY := 0
	for _, point := range points {
		highestY = Max(highestY, point.Y)
	}

	return highestY
}

func part2(x1 int, x2 int, y1 int, y2 int) int {
	count := 0

	for x := 0; x <= Abs(x2); x++ {
		for y := y1; Abs(y) <= Abs(y1); y++ {
			initialVelocity := Point{x, y}
			if _, hit := Trajection(initialVelocity, x1, x2, y1, y2); hit {
				count++
			}
		}
	}

	return count
}

func Trajection(initialVelocity Point, x1 int, x2 int, y1 int, y2 int) ([]Point, bool) {
	points := []Point{}

	probe := Point{
		X: 0,
		Y: 0,
	}
	velocity := initialVelocity

	points = append(points, probe)

	for {
		probe.X += velocity.X
		probe.Y += velocity.Y

		points = append(points, probe)

		// Drag adjustment
		if velocity.X > 0 {
			velocity.X -= 1
		} else if velocity.X < 0 {
			velocity.X += 1
		}

		// Gravity
		velocity.Y -= 1

		if probe.X >= x1 && probe.X <= x2 &&
			probe.Y >= y1 && probe.Y <= y2 {
			return points, true
		}

		if probe.Y < y1 {
			break
		}
	}

	return points, false
}
