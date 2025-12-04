package day2

import "fmt"

func part1(input Input) error {
	for _, v := range input.Ranges {
		fmt.Printf("%d - %d\n", v.From, v.To)
	}

	return nil
}
