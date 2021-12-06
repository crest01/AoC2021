package input

import (
	"advent_of_code/day5/vents"
	"os"
)

func ParseFile(filename string, only_straight bool) []vents.Line {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	padded := pad(string(data))
	numbers := lex(padded)
	var coords []vents.Int2
	for i := 0; len(numbers) >= 2; i++ {
		coords = append(coords, vents.MakeInt2(&numbers))
	}
	var lines []vents.Line

	for i := 0; len(coords) >= 2; i++ {
		line := vents.MakeLine(&coords)
		if only_straight {
			if line.IsStraight() {
				lines = append(lines, line)
			}
		} else {
			lines = append(lines, line)
		}
	}
	return lines
}
