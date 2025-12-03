package shared

import (
	"errors"
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"os"
)

func RunPuzzle[T any](inputFile string, part uint, lexer *lexer.StatefulDefinition, part1 func(input T) error, part2 func(input T) error) error {
	input, err := Parse[T](inputFile, lexer)
	if err != nil {
		return err
	}

	switch part {
	case 1:
		return part1(*input)
	case 2:
		return part2(*input)
	}

	return errors.New("Part could not be found")
}

func Parse[T any](inputFile string, lexer *lexer.StatefulDefinition) (*T, error) {
	parser, err := participle.Build[T](participle.Lexer(lexer))
	if err != nil {
		return nil, err
	}

	r, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	ast, err := parser.Parse(inputFile, r)
	if err != nil {
		return nil, err
	}

	return ast, nil
}
