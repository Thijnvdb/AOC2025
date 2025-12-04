package day1

import (
	"fmt"
)

func part2(input Lines) error {
	value := 50

	result := 0

	for _, v := range input.Lines {
		newValue := value
		num := v.Number

		wholeRotations := num / 100
		partialRotationSteps := num % 100

		if v.Prefix == "L" {
			partialRotationSteps *= -1
		}

		newValue += partialRotationSteps

		if newValue < 0 {
			newValue = 100 + newValue
		}

		newValue %= 100

		if value != 0 && (newValue == 0 || (v.Prefix == "L" && newValue > value) || (v.Prefix == "R" && newValue < value)) {
			wholeRotations++
		}

		fmt.Printf("moved from %d by %d%s to %d, %d rotations\n", value, num, v.Prefix, newValue, wholeRotations)

		result += wholeRotations

		value = newValue
	}

	fmt.Printf("Amount of times the dial passed 0: %d\n", result)

	return nil
}
