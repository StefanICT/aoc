package main

import (
    "fmt"
    . "aoc/util"
    "sort"
)

func main() {
    part1()
    part2()
}

type Chunk struct {
    Start int
    End int
    Character byte

    Parent *Chunk
    Children []Chunk
}

func part1() {
    pairs := map[byte]byte{
        '(': ')',
        '[': ']',
        '{': '}',
        '<': '>',
    }

    points := map[byte]int{
        ')': 3,
        ']': 57,
        '}': 1197,
        '>': 25137,
    }

    sum := 0

    for _, line := range ReadLines("2021/10/input") {
        previous := []byte{}

        for i := 0; i < len(line); i++ {
            if _, ok := pairs[line[i]]; ok {
                previous = append(previous, line[i])
            } else if pairs[previous[len(previous) - 1]] == line[i] {
                previous = previous[:len(previous) - 1]
            } else {
                sum += points[line[i]]
                break
            }
        }
    }

    fmt.Println(sum)
}

func part2() {
    pairs := map[byte]byte{
        '(': ')',
        '[': ']',
        '{': '}',
        '<': '>',
    }

    points := map[byte]int{
        ')': 1,
        ']': 2,
        '}': 3,
        '>': 4,
    }

    scores := []int{}

    lines:
    for _, line := range ReadLines("2021/10/input") {
        previous := []byte{}

        for i := 0; i < len(line); i++ {
            if _, ok := pairs[line[i]]; ok {
                previous = append(previous, line[i])
            } else if pairs[previous[len(previous) - 1]] == line[i] {
                previous = previous[:len(previous) - 1]
            } else {
                continue lines
            }
        }

        score := 0
        for i := len(previous) - 1; i >= 0; i-- {
            score = score * 5 + points[pairs[previous[i]]]
        }

        scores = append(scores, score)
    }
    sort.Ints(scores)
    fmt.Println(scores[len(scores) / 2])
}
