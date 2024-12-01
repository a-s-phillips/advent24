package helpers

import (
	"math"
	"strings"
)

func TransposeArray[T any](arr [][]T) [][]T {
	rows := len(arr)
	cols := len(arr[0])

	newArr := make([][]T, cols)
	for i := range newArr {
		newArr[i] = make([]T, rows)
	}

	for i := range arr {
		for j := range arr[i] {
			newArr[j][i] = arr[i][j]
		}
	}

	return newArr
}

func MathMod(x int, n int) int {
	xf := float64(x)
	nf := float64(n)
	return int(xf - (nf * math.Floor((xf / nf))))
}

func SplitAndTrimInput(input string) []string {
	lines := strings.Split(input, "\n")
	for len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return lines
}
