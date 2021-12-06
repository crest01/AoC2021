package bingo

import (
	"fmt"
	"testing"
)

func TestRowWin(t *testing.T) {
	board := Board{}
	board.AddRow([5]uint8{1, 2, 3, 4, 5}, 0)
	board.AddRow([5]uint8{6, 7, 8, 9, 10}, 1)
	board.AddRow([5]uint8{11, 12, 13, 14, 15}, 2)
	board.AddRow([5]uint8{16, 17, 18, 19, 20}, 3)
	board.AddRow([5]uint8{21, 22, 23, 24, 25}, 4)

	board.AddHit(6)
	board.AddHit(7)
	board.AddHit(8)
	board.AddHit(9)
	board.AddHit(10)

	is := board.CheckWin()
	board.Print()
	fmt.Print("\n")
	want := true
	if is != want {
		t.Fatalf("Board didn't win with the second row")
	}
}

func TestColWin(t *testing.T) {
	board := Board{}
	board.AddRow([5]uint8{1, 2, 3, 4, 5}, 0)
	board.AddRow([5]uint8{6, 7, 8, 9, 10}, 1)
	board.AddRow([5]uint8{11, 12, 13, 14, 15}, 2)
	board.AddRow([5]uint8{16, 17, 18, 19, 20}, 3)
	board.AddRow([5]uint8{21, 22, 23, 24, 25}, 4)

	board.AddHit(2)
	board.AddHit(7)
	board.AddHit(12)
	board.AddHit(17)
	board.AddHit(22)

	is := board.CheckWin()
	fmt.Print("\n")
	board.Print()
	want := true
	if is != want {
		t.Fatalf("Board didn't win with the second column")
	}
}

func TestDiagWin(t *testing.T) {
	board := Board{}
	board.AddRow([5]uint8{1, 2, 3, 4, 5}, 0)
	board.AddRow([5]uint8{6, 7, 8, 9, 10}, 1)
	board.AddRow([5]uint8{11, 12, 13, 14, 15}, 2)
	board.AddRow([5]uint8{16, 17, 18, 19, 20}, 3)
	board.AddRow([5]uint8{21, 22, 23, 24, 25}, 4)

	board.AddHit(1)
	board.AddHit(7)
	board.AddHit(13)
	board.AddHit(19)
	board.AddHit(25)

	is := board.CheckWin()
	fmt.Print("\n")
	board.Print()
	want := false
	if is != want {
		t.Fatalf("Board didn't win with the diagonal")
	}
}

func TestScore(t *testing.T) {
	board := Board{}
	board.AddRow([5]uint8{1, 2, 3, 4, 5}, 0)
	board.AddRow([5]uint8{6, 7, 8, 9, 10}, 1)
	board.AddRow([5]uint8{11, 12, 13, 14, 15}, 2)
	board.AddRow([5]uint8{16, 17, 18, 19, 20}, 3)
	board.AddRow([5]uint8{21, 22, 23, 24, 25}, 4)

	board.AddHit(1)
	board.AddHit(7)
	board.AddHit(13)
	board.AddHit(19)
	board.AddHit(25)

	is := board.GetScore()
	fmt.Print("\n")
	board.Print()
	fmt.Printf("Score: %d", is)
	want := uint32(260)
	if is != want {
		t.Fatalf("Board got the wrong score: is=%d, want=%d", is, want)
	}
}
