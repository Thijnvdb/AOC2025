package day1

import "github.com/alecthomas/participle/v2/lexer"

var (
	day1Lexer = lexer.MustSimple([]lexer.SimpleRule{
		{"Prefix", `(L|R)`},
		{"Number", `[0-9]+`},
		{"newLine", `\n`},
	})
)

type Lines struct {
	Lines []*Line `parser:"@@*"`
}

type Line struct {
	Prefix string `parser:"@Prefix"`
	Number int    `parser:"@Number"`
}
