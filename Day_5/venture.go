package main

import (
	"advent_of_code/day5/input"
	"advent_of_code/day5/vents"
	"fmt"
)

func main() {

	vent_lines := input.ParseFile("./input.txt", false)

	_, max := vents.CalcAABB(&vent_lines)

	canvas := make([][]uint8, max.X+1)
	for i := range canvas {
		canvas[i] = make([]uint8, max.Y+1)
	}

	for idx := range vent_lines {
		line := &vent_lines[idx]
		(*line).Rasterize(&canvas)
	}
	count_larger := 0
	for x := 0; x < max.X; x++ {
		for y := 0; y < max.Y; y++ {
			if canvas[x][y] > 1 {
				count_larger++
			}
		}
	}
	fmt.Printf("Overlapping points: %d", count_larger)
}
