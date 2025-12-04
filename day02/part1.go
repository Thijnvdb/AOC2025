package day2

import (
	"fmt"
	"strconv"
)

func splitAsNumbers(str string) (int, int) {
	left, errLeft := strconv.Atoi(str[:len(str)/2])
	right, errRight := strconv.Atoi(str[len(str)/2:])

	if errLeft != nil || errRight != nil {
		return 0, 0
	}

	return left, right
}

func checkRange(rng *Range) int {
	result := 0

	for i := rng.From; i <= rng.To; i++ {
		str := strconv.Itoa(i)
		if len(str)%2 != 0 {
			continue
		}

		left, right := splitAsNumbers(str)
		if left == right && left != 0 {
			fmt.Printf("found invalid id: %d\n", i)
			result += i
		}
	}

	if result == 0 {
		fmt.Printf("no invalid id's found\n")
	}

	return result
}

func part1(input Input) error {
	result := 0
	for _, v := range input.Ranges {
		fmt.Printf("\nchecking %d-%d\n", v.From, v.To)
		result += checkRange(v)
	}

	fmt.Printf("\nMax value: %d\n", result)

	return nil
}
