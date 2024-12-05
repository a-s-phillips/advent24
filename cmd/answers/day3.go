package answers

import (
	//	"advent24/cmd/helpers"
	"log"
	"regexp"
	"strconv"
)

var day3Test1 = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
var day3Test2 = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

func Day3(rawInput string) Answer {
	cmdRgx := regexp.MustCompile("mul\\([0-9]+,[0-9]+\\)|do\\(\\)|don't\\(\\)")
	mulCmds := cmdRgx.FindAllString(rawInput, -1)

	pairRgx := regexp.MustCompile("[0-9]+")
	var mulPairs [][2]int
	do := true
	for _, cmd := range mulCmds {
		if cmd == "do()" {
			do = true
		} else if cmd == "don't()" {
			do = false
		} else {
			if do == true {
				nums := pairRgx.FindAllString(cmd, 2)
				first, err := strconv.Atoi(nums[0])
				if err != nil {
					log.Panic("Could not convert item at index 0 in pair regex result to integer")
				}
				second, err := strconv.Atoi(nums[1])
				if err != nil {
					log.Panic("Could not convert item at index 1 in pair regex result to integer")
				}
				mulPairs = append(mulPairs, [2]int{first, second})
			}
		}
	}

	productSum := 0
	for _, pair := range mulPairs {
		productSum += pair[0] * pair[1]
	}

	part1 := strconv.Itoa(productSum)
	part2 := "part2"
	return Answer{Part1: &part1, Part2: &part2}
}
