package util

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

var MaxInt = math.MaxInt
var MinInt = math.MinInt

func ReadLines(path string) []string {
	file, _ := os.Open(path)
	defer file.Close()

	lines := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func ReadInts(path string) []int {
	file, _ := os.Open(path)
	defer file.Close()

	ints := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ints = append(ints, AsInt(scanner.Text()))
	}

	return ints
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func Memset(ints *[]int, value int) {
	for i := 0; i < len(*ints); i++ {
		(*ints)[i] = value
	}
}

func Min(numbers ...int) int {
	n := numbers[0]
	for i := range numbers {
		if numbers[i] < n {
			n = numbers[i]
		}
	}

	return n
}

func Max(numbers ...int) int {
	n := numbers[0]
	for i := range numbers {
		if numbers[i] > n {
			n = numbers[i]
		}
	}

	return n
}

func AsInts(s string, separator string) []int {
	components := strings.Split(s, separator)
	ints := make([]int, len(components))
	for i := 0; i < len(components); i++ {
		ints[i] = AsInt(components[i])
	}

	return ints
}

func AsInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
