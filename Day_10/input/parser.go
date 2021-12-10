package input

import (
	"fmt"
	"os"
)

func asInt(data *[]byte, pos int) int8 {
	value := int8((*data)[pos] - 48)
	return value
}

func GetClosingBracket(in byte) byte {
	switch in {
	case '<':
		return '>'
	case '(':
		return ')'
	case '{':
		return '}'
	case '[':
		return ']'
	default:
		panic("Unknown opening bracket " + string(in))
	}
}

// Pad string with YYMAXFILL zeroes at the end.
func pad(data *[]byte, fill int) *[]byte {
	padding := make([]byte, fill)
	*data = append(*data, padding...)
	return data
}

func printError(str *[]byte, start int, end int) {
	substr := string((*str)[start:end])
	fmt.Printf("Error: unexpected input '%s' at position %d", substr, start)
	panic("Unexpected Input")
}

func Parse(input *[]byte) ([][]byte, []byte, []int, [][]byte) {
	return lex(input)
}

func ParseFile(filename string) ([][]byte, []byte, []int, [][]byte) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return Parse(&data)
}
