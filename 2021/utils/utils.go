package utils

import (
	"strconv"
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

func AsInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
