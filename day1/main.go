package day1

import (
	. "thijnvdb/aoc2025/shared"
)

var Day1 = Day{InputFile: "day1/inputs/input.txt", ExampleFile: "day1/inputs/example.txt", Run: run}

func run(inputFile string, part uint) error {
	return RunPuzzle(inputFile, part, day1Lexer, part1, part2)
}
