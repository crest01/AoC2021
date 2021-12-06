package main

import (
	"advent_of_code/day4/input"
	"fmt"
)

func main() {

	all_boards, drawn_values := input.ParseFile("./input.txt")

	var won_boards []int
	var next_boards []int
	var boards []int

	for idx := range all_boards {
		boards = append(boards, idx)
	}

	for n_idx, drawn_value := range drawn_values {
		fmt.Printf("Drawing %3d (%3d/%3d)\n", drawn_value, n_idx+1, len(drawn_values))
		var hit_boards []uint8
		var num_hits []uint8
		for _, idx := range boards {
			board := &all_boards[idx]
			if board.AddHit(drawn_value) {
				hit_boards = append(hit_boards, uint8(idx))
				num_hits = append(num_hits, board.NumHits())
			}
		}
		//fmt.Printf("We got %d hits on boards: ", len(hit_boards))

		// for i := 0; i < len(hit_boards); i++ {
		// 	fmt.Printf("%d (%d), ", hit_boards[i], num_hits[i])
		// }
		fmt.Printf("\n")
		next_boards = nil
		for _, idx := range boards {
			board := &all_boards[idx]
			if board.CheckWin() {
				fmt.Printf("Board %d Won\n", idx)
				// score := board.GetScore()
				// fmt.Printf("Board Score: %d\n", score)
				won_boards = append(won_boards, idx)
			} else {
				next_boards = append(next_boards, idx)
			}
		}
		boards = next_boards
		if len(boards) == 0 {
			last_idx := won_boards[len(won_boards)-1]
			board := &all_boards[last_idx]
			fmt.Printf("Board %d Won last.\n", last_idx)
			board.Print()
			score := board.GetScore()
			fmt.Printf("Board Score: %d\n", score)
			fmt.Printf("Score * current Number = %d\n", uint32(drawn_value)*score)
			break
		}
	}
}
