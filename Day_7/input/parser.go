package input

import (
	"fmt"
	"os"
	"strconv"
)

// Pad string with YYMAXFILL zeroes at the end.
func pad(data *[]byte, fill int) *[]byte {
	padding := make([]byte, fill)
	*data = append(*data, padding...)
	return data
}

func asInt(data *[]byte, start int, end int) int {
	substr := string((*data)[start:end])
	val, err := strconv.Atoi(substr)
	if err != nil {
		panic(err)
	}
	return int(val)
}

func printError(str *[]byte, start int, end int) {
	substr := string((*str)[start:end])
	fmt.Printf("Error: unexpected input '%s' at position %d", substr, start)
	panic("Unexpected Input")
}

func Parse(input *[]byte) []int {
	return lex(input)
}

func ParseFile(filename string) []int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return Parse(&data)
}
