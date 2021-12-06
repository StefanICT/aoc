package utils

import (
	"strconv"
	"strings"
)

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

func Max(x int, y int) int {
	if x < y {
		return y
	}
	return x
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
