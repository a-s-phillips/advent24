package main

import (
	"advent24/cmd/answers"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Command func(input string) answers.Answer

var commands = map[int]Command{
  1: answers.Day1,
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	if len(os.Args) < 2 {
		fmt.Println("Provide a numeric command between 1 and 25")
		os.Exit(1)
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Command must be a number")
		os.Exit(1)
	}

	if num < 1 || num > 25 {
		fmt.Println("Command cannot be less than 1 or greater than 25")
		os.Exit(1)
	}

	cmd := commands[num]
	inputDir := os.Getenv("INPUT_DIR")
	content, err := os.ReadFile(inputDir + "day" + os.Args[1])
	if err != nil {
		fmt.Println("Could not find input file for day " + os.Args[1])
		os.Exit(1)
	}

	string := string(content)
	answer := cmd(string)
	fmt.Println("Day " + os.Args[1])
	fmt.Print("Part 1:")
	if answer.Part1 == nil {
		fmt.Println(" nil")
	} else {
		fmt.Println(" " + *answer.Part1)
	}

	fmt.Print("Part 2:")
	if answer.Part2 == nil {
		fmt.Println(" nil")
	} else {
		fmt.Println(" " + *answer.Part2)
	}
}
