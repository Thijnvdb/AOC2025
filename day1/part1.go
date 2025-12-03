package day1

import "fmt"

func part1(input Lines) error {
	value := 50

	result := 0

	for _, v := range input.Lines {
		num := v.Number
		if v.Prefix == "L" {
			num *= -1
		}

		value += num

		if value < 0 {
			value = 100 + value
		}

		value %= 100

		if value == 0 {
			result++
		}
	}

	fmt.Printf("Amount of times the dial was left at 0: %d\n", result)

	return nil
}
