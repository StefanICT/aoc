package main

import (
	. "aoc/util"
	"container/heap"
	"fmt"
)

type Distance struct {
	i        int
	j        int
	priority int
}

// Using the distances map to search for the new lowest value each time is too
// slow. Using a priority queue which keeps the lowest value at a O(n)
// operation away makes the program fast.
//
// Go kinda offers a priority queue using the heap package. Maybe in the future
// Go with generics can offer this implementation self.
type PriorityQueue []Distance

func (queue PriorityQueue) Len() int {
	return len(queue)
}

func (queue PriorityQueue) Less(index int, other int) bool {
	return queue[index].priority < queue[other].priority
}

func (queue PriorityQueue) Swap(index int, other int) {
	queue[index], queue[other] = queue[other], queue[index]
}

func (queue *PriorityQueue) Push(element interface{}) {
	*queue = append(*queue, element.(Distance))
}

func (queue *PriorityQueue) Pop() interface{} {
	element := (*queue)[len(*queue)-1]
	*queue = (*queue)[:len(*queue)-1]

	return element
}

var INF = MaxInt

func main() {
	lines := ReadLines("2021/15/input")
	grid := make([][]int, len(lines))

	for i, line := range lines {
		grid[i] = make([]int, len(line))

		for j := 0; j < len(line); j++ {
			grid[i][j] = int(line[j] - '0')
		}
	}

	fmt.Println(LowestTotalRisk(grid))
	fmt.Println(LowestTotalRisk(grow(grid, 5)))
}

func LowestTotalRisk(grid [][]int) int {
	distances := make([][]int, len(grid))
	visited := make([][]bool, len(grid))
	for z := 0; z < len(grid); z++ {
		distances[z] = make([]int, len(grid[z]))
		visited[z] = make([]bool, len(grid[z]))
		for b := 0; b < len(grid[z]); b++ {
			distances[z][b] = INF
		}
	}
	distances[0][0] = 0

	queue := &PriorityQueue{
		{0, 0, 0},
	}

	heap.Init(queue)

	for queue.Len() != 0 {
		distance := heap.Pop(queue).(Distance)
		i := distance.i
		j := distance.j

		visited[i][j] = true

		for _, point := range [][]int{{-1, 0}, {+1, 0}, {0, -1}, {0, +1}} {
			adjacentI := i + point[0]
			adjacentJ := j + point[1]

			if adjacentI < 0 || adjacentI == len(grid) || adjacentJ < 0 || adjacentJ == len(grid[adjacentI]) {
				continue
			}

			if visited[adjacentI][adjacentJ] {
				continue
			}

			if distances[i][j]+grid[adjacentI][adjacentJ] < distances[adjacentI][adjacentJ] {
				distances[adjacentI][adjacentJ] = distances[i][j] + grid[adjacentI][adjacentJ]
				heap.Push(queue, Distance{adjacentI, adjacentJ, distances[adjacentI][adjacentJ]})
			}
		}
	}

	return distances[len(distances)-1][len(distances[len(distances)-1])-1]
}

func grow(grid [][]int, iter int) [][]int {
	grown := make([][]int, len(grid))
	copy(grown, grid)

	for i := 0; i < len(grid); i++ {
		ptr := &grid[i]

		for j := 0; j < iter-1; j++ {
			row := increase(*ptr)
			grown[i] = append(grown[i], row...)
			ptr = &row
		}
	}

	for j := 0; j < iter-1; j++ {
		for i := 0; i < len(grid); i++ {
			grown = append(grown, increase(grown[len(grid)*j+i]))
		}
	}

	return grown
}

func increase(list []int) []int {
	numbers := make([]int, len(list))
	for i := 0; i < len(list); i++ {
		value := list[i] + 1
		if value > 9 {
			value = 1
		}

		numbers[i] = value
	}

	return numbers
}

func display(grid [][]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != INF {
				fmt.Printf("%d", grid[i][j])
			} else {
				fmt.Printf("∞")
			}
		}
		fmt.Println()
	}
}
