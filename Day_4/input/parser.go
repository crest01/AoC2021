package input

import (
	"advent_of_code/day4/bingo"
	"os"
	"strconv"
	"strings"
)

func ParseFile(filename string) ([]bingo.Board, []uint8) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var draw_values []uint8
	var boards []bingo.Board
	var current_board *bingo.Board
	drawn_values := true
	row_idx := uint8(0)
	for _, entry := range strings.Split(string(data), "\n") {
		if len(entry) == 0 {
			continue
		}
		if drawn_values {
			for _, numbers := range strings.Split(entry, ",") {
				number, err := strconv.ParseUint(numbers, 10, 8)
				if err != nil {
					panic(err)
				}
				draw_values = append(draw_values, uint8(number))
			}
			drawn_values = false
			continue
		}
		if len(boards) == 0 {
			boards = append(boards, bingo.Board{})
		}
		current_board = &boards[len(boards)-1]
		var row_values [5]uint8

		for idx, numbers := range strings.Fields(entry) {
			number, err := strconv.ParseUint(numbers, 10, 8)
			if err != nil {
				panic(err)
			}
			row_values[idx] = uint8(number)
		}
		current_board.AddRow(row_values, row_idx)
		row_idx++
		if row_idx >= 5 {
			boards = append(boards, bingo.Board{})
			row_idx = 0
		}
	}

	return boards, draw_values
}
