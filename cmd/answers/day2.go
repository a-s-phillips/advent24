package answers

import (
	"advent24/cmd/helpers"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
)

var day2Test = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

var edgeCases = `48 46 47 49 51 54 56
1 1 2 3 4 5
1 2 3 4 5 5
5 1 2 3 4 5
1 4 3 2 1
1 6 7 8 9
1 2 3 4 3
9 8 7 6 7
7 10 8 10 11
29 28 27 25 26 25 22 20`

// The levels are either all increasing or all decreasing.
// Any two adjacent levels differ by at least one and at most three.

func unique(slice []int) []int {
	// create a map with all the values as key
	uniqMap := make(map[int]struct{})
	for _, v := range slice {
		uniqMap[v] = struct{}{}
	}

	// turn the map keys into a slice
	uniqSlice := make([]int, 0, len(uniqMap))
	for v := range uniqMap {
		uniqSlice = append(uniqSlice, v)
	}
	return uniqSlice
}

func checkReport(report []int) []int {
	var increasing bool
	var faultIdxs []int
	for i, num := range report[0 : len(report)-1] {
		diff := num - report[i+1]
		if i == 0 && diff > 0 {
			increasing = false
		}
		if i == 0 && diff < 0 {
			increasing = true
		}

		absDiff := math.Abs(float64(diff))

		zeroDiff := diff == 0
		wrongDir := i != 0 && ((diff > 0 && increasing) || (diff < 0 && !increasing))
		notInRange := absDiff > 3 || absDiff < 1
		if zeroDiff || wrongDir || notInRange {
			faultIdxs = append(faultIdxs, i)
		}
	}

	return faultIdxs
}

func Day2(input string) Answer {
	var splitInput = helpers.SplitAndTrimInput(input)

	reports := [][]int{}
	for _, line := range splitInput {
		splitLine := strings.Split(line, " ")
		report := []int{}
		for _, str := range splitLine {
			parsedNum, err := strconv.Atoi(str)
			if err != nil {
				log.Panic("Could not parse number in input")
			}
			report = append(report, parsedNum)
		}
		reports = append(reports, report)
	}

	part1SafeReports := 0
	part2SafeReports := 0
	for _, report := range reports {
		faults := checkReport(report)

		if len(faults) == 0 {
			part1SafeReports += 1
			part2SafeReports += 1
		} else {

			var removals []int
			for _, fault := range faults {
				if fault == 1 {
					removals = append(removals, 0)
				}
				removals = append(removals, fault)
				removals = append(removals, fault+1)
				removals = unique(removals)
				slices.SortFunc(removals, helpers.SortAsc)
			}

			for _, removal := range removals {
				reportCopy := make([]int, len(report))
				copy(reportCopy, report)
				alteredReport := append(reportCopy[:removal], reportCopy[removal+1:]...)
				if len(checkReport(alteredReport)) == 0 {
					part2SafeReports += 1
					break
				}
			}
		}
	}

	part1 := strconv.Itoa(part1SafeReports)
	part2 := strconv.Itoa(part2SafeReports)

	return Answer{Part1: &part1, Part2: &part2}
}
