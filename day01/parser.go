package day1

import "github.com/alecthomas/participle/v2/lexer"

var (
	day1Lexer = lexer.MustSimple([]lexer.SimpleRule{
		{Name: "Prefix", Pattern: `(L|R)`},
		{Name: "Number", Pattern: `[0-9]+`},
		{Name: "newLine", Pattern: `\n`},
	})
)

type Lines struct {
	Lines []*Line `parser:"@@*"`
}

type Line struct {
	Prefix string `parser:"@Prefix"`
	Number int    `parser:"@Number"`
}
