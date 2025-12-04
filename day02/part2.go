package day2

import (
	"fmt"
	"github.com/dlclark/regexp2"
	"strconv"
)

var regex = regexp2.MustCompile(`^(\d+)\1+$`, regexp2.None)

func checkRangeWithRegex(rng *Range) int {
	result := 0
	for i := rng.From; i <= rng.To; i++ {
		str := strconv.Itoa(i)
		match, err := regex.MatchString(str)
		if err != nil {
			fmt.Printf("Error while checking regex: %s", err.Error())
			continue
		}

		if match {
			fmt.Printf("found invalid id: %d\n", i)
			result += i
		}
	}

	if result == 0 {
		fmt.Printf("no invalid id's found\n")
	}

	return result
}

func part2(input Input) error {
	result := 0
	for _, v := range input.Ranges {
		fmt.Printf("\nchecking %d-%d\n", v.From, v.To)
		result += checkRangeWithRegex(v)
	}

	fmt.Printf("\nMax value: %d\n", result)

	return nil
}
