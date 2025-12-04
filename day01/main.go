package day1

import (
	. "thijnvdb/aoc2025/shared"
)

var Definition = Day{InputFile: "day01/inputs/input.txt", ExampleFile: "day01/inputs/example.txt", Run: run}

func run(inputFile string, part uint) error {
	return RunPuzzle(inputFile, part, day1Lexer, part1, part2)
}
