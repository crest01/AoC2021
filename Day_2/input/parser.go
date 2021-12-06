package input

import (
	"os"
	"strconv"
	"strings"
)

func to_direction(token string) Direction {
	switch {
	case token == "forward":
		return Forward
	case token == "down":
		return AimDown
	case token == "up":
		return AimUp
	}
	return Undefined
}

func to_size(token string) uint {
	val, _ := strconv.Atoi(token)
	return uint(val)
}

func ParseFile(filename string) []Instruction {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var instructions []Instruction
	for _, entry := range strings.Split(string(data), "\n") {
		if len(entry) == 0 {
			continue
		}
		tokens := strings.Split(entry, " ")
		var value Instruction
		value.Dir = to_direction(tokens[0])
		value.Size = to_size(tokens[1])
		instructions = append(instructions, value)
	}
	return instructions
}
