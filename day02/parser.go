package day2

import "github.com/alecthomas/participle/v2/lexer"

var (
	day2Lexer = lexer.MustSimple([]lexer.SimpleRule{
		{Name: "Number", Pattern: `[0-9]+`},
		{Name: "newLine", Pattern: `\n`},
		{Name: "seperator", Pattern: `[,-]`},
	})
)

type Input struct {
	Ranges []*Range `parser:"@@*"`
}

type Range struct {
	From int `parser:"@Number"`
	To   int `parser:"@Number"`
}
