package answers

import (
	//  "fmt"
	"advent24/cmd/helpers"
	"log"
	"slices"
	"strconv"
	"strings"
)

var testData = `3   4
4   3
2   5
1   3
3   9
3   3`

var similarities = make(map[int]int)

func sortAsc(a, b int) int {
	if a < b {
		return -1
	}
	return 1
}

func getSimilarity(list []int, num1 int) int {
	var filteredList []int
	for _, val := range list {
		if val == num1 {
			filteredList = append(filteredList, val)
		}
	}
	return num1 * len(filteredList)
}

func Day1(input string) Answer {
	var splitInput = helpers.SplitAndTrimInput(input)
	var col1 []int
	var col2 []int
	for _, line := range splitInput {
		splitLine := strings.Split(line, "   ")

		firstParsed, err := strconv.Atoi(splitLine[0])
		if err != nil {
			log.Panic("Failed to parse line in first input")
		}
		col1 = append(col1, firstParsed)

		secondParsed, err := strconv.Atoi(splitLine[1])
		if err != nil {
			log.Panic("Failed to parse line in second input")
		}
		col2 = append(col2, secondParsed)
	}

	sortedCol1 := col1
	sortedCol2 := col2
	slices.SortFunc(sortedCol1, sortAsc)
	slices.SortFunc(sortedCol2, sortAsc)

	sum := 0
	totalSimilarity := 0
	for i, num1 := range sortedCol1 {
		num2 := sortedCol2[i]
		var diff int
		if num1 < num2 {
			diff = num2 - num1
		} else {
			diff = num1 - num2
		}
		sum += diff

		value, exists := similarities[num1]
		if exists {
			totalSimilarity += value
		} else {
			similarity := getSimilarity(col2, num1)
			similarities[num1] = similarity
			totalSimilarity += similarity
		}
	}

	part1 := strconv.Itoa(sum)
	part2 := strconv.Itoa(totalSimilarity)
	return Answer{Part1: &part1, Part2: &part2}
}
