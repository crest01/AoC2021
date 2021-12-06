package bingo

import (
	"fmt"
)

type Board struct {
	values   [5][5]uint8
	hit      [5][5]bool
	num_hits uint8
}

func (board *Board) GetScore() uint32 {
	var score uint32
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			value := board.values[row][col]
			if !board.hit[row][col] {
				score += uint32(value)
			}
		}
	}
	return score
}

func (board *Board) AddRow(values [5]uint8, row uint8) {
	for col := 0; col < 5; col++ {
		board.values[row][col] = values[col]
	}
}

func (board *Board) AddHit(value uint8) bool {
	found := false
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if board.values[row][col] == value {
				board.hit[row][col] = true
				found = true
				board.num_hits++
				break
			}
		}
	}
	return found
}

func (board *Board) NumHits() uint8 {
	return board.num_hits
}

func (board *Board) checkRowWin() (bool, int) {
	for row := 0; row < 5; row++ {
		win := true
		for col := 0; col < 5; col++ {
			win = win && board.hit[row][col]
		}
		if win {
			return true, row
		}
	}
	return false, 0
}

func (board *Board) checkColWin() (bool, int) {
	for col := 0; col < 5; col++ {
		win := true
		for row := 0; row < 5; row++ {
			win = win && board.hit[row][col]
		}
		if win {
			return true, col
		}
	}
	return false, 0
}

func (board *Board) checkDiagonalWin() bool {
	win := true
	row := 0
	col := 0
	for i := 0; i < 5; i++ {
		win = win && board.hit[row][col]
		row++
		col++
	}
	return win
}

func (board *Board) CheckWin() bool {
	col_win, _ := board.checkColWin()
	row_win, _ := board.checkRowWin()
	//diag_win := board.checkDiagonalWin()
	return col_win || row_win
}

func (board *Board) Print() {
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			value := board.values[row][col]
			hit := board.hit[row][col]
			if hit {
				fmt.Printf(" *%2d", value)
			} else {
				fmt.Printf("  %2d", value)
			}
		}
		fmt.Print("\n")
	}
}
