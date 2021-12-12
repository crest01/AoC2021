package main

import (
	"advent_of_code/day11/input"
	"advent_of_code/day11/octopi"
	"fmt"
)

func main() {

	lines := input.ParseFile("./input.txt")
	fmt.Print("Task 1: Find number of Flashes\n")
	{
		field := octopi.MakeField(lines)
		fmt.Printf("Initial field:\n%s\n", field.ToString())
		num_blinks := 0
		num_rounds := 100
		for i := 0; i < num_rounds; i++ {
			blinks := octopi.SimulateStep(field)
			num_blinks += len(blinks)
		}
		fmt.Printf("Sum Flashes after %d rounds: %d\n", num_rounds, num_blinks)
	}
	fmt.Print("Task 2: Find Simultaneous Flashes\n")
	{
		field := octopi.MakeField(lines)
		fmt.Printf("Initial field:\n%s\n", field.ToString())
		round := 1
		found := false
		for round < 1000 {
			blinks := octopi.SimulateStep(field)
			if len(blinks) == 100 {
				fmt.Printf("Simultaneous Flashes, Step %d:\n%s\n", round, field.ToString())
				found = true
				break
			}
			round++
		}
		if !found {
			fmt.Printf("No Simulatneous flashes in %d rounds\n", 1000)
		}
	}
}
