package main

import (
	"flag"
	"fmt"
	"thijnvdb/aoc2025/day01"
	"thijnvdb/aoc2025/day02"
	. "thijnvdb/aoc2025/shared"
)

var days []Day = []Day{day1.Definition, day2.Definition}

func main() {
	day := flag.Uint("day", 1, "the day")
	part := flag.Uint("part", 1, "the part")
	example := flag.Bool("example", false, "is example?")

	flag.Parse()

	if *part > 2 {
		fmt.Println("Error: part cannot be larger than 2")
		return
	}

	if *day > uint(len(days)) || *day < 1 {
		fmt.Println("Error: day cannot be larger than the amount of days, or smaller than 1")
		return
	}

	dayStruct := days[*day-1]

	file := dayStruct.InputFile
	if *example {
		fmt.Println("running in example mode!")
		file = dayStruct.ExampleFile
	}

	err := dayStruct.Run(file, *part)
	if err != nil {
		fmt.Println("Error occurred while running the puzzle: ", err.Error())
	}
}
