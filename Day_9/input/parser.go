package input

import (
	"fmt"
	"os"
)

func asInt(data *[]byte, pos int) int8 {
	value := int8((*data)[pos] - 48)
	return value
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

func Parse(input *[]byte) [][]int8 {
	return lex(input)
}

func ParseFile(filename string) [][]int8 {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return Parse(&data)
}
