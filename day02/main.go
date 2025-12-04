package day2

import (
	. "thijnvdb/aoc2025/shared"
)

var Definition = Day{InputFile: "day02/inputs/input.txt", ExampleFile: "day02/inputs/example.txt", Run: run}

func run(inputFile string, part uint) error {
	return RunPuzzle(inputFile, part, day2Lexer, part1, part2)
}
