package input

import (
	"advent_of_code/day13/sparse_field"
	"fmt"
	"os"
	"strconv"
)

func asInt(data *[]byte, start int, end int) int {
	substr := string((*data)[start:end])
	val, err := strconv.Atoi(substr)
	if err != nil {
		panic(err)
	}
	return int(val)
}

func asNumber(data *[]byte, pos int) int {
	value := int((*data)[pos] - 48)
	return value
}

func asString(data *[]byte, start int, end int) string {
	return string((*data)[start:end])
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

func Parse(input *[]byte) ([]sparse_field.Int2, []sparse_field.Int2) {
	return lex(input)
}

func ParseFile(filename string) ([]sparse_field.Int2, []sparse_field.Int2) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return Parse(&data)
}
