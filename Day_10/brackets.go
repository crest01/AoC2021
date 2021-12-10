package main

import (
	"advent_of_code/day10/input"
	"fmt"
)

func main() {

	canvas := input.ParseFile("./test_input.txt")
	fmt.Print("Task 1: Find corrupt Lines\n")
	{
		_, heights := findBasins(canvas)
		fmt.Print("  Found ", len(heights), " basins\n")
		sum := 0
		for _, height := range heights {
			sum = sum + int(height) + 1
		}
		fmt.Print("  Sum: ", sum, "\n")
	}
	fmt.Print("Task 2: ???\n")
	{
	}
}
