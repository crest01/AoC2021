package input

import (
	"os"
)

func Parse(input string) []uint64 {
	padded := pad(input)
	return lex(padded)
}

func ParseFile(filename string) []uint64 {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return Parse(string(data))
}
