package main

import (
	"advent_of_code/day8/input"
	"advent_of_code/day8/signal"
	"fmt"
)

func pow10(exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= 10
	}
	return result
}

func main() {
	//signal.PrintAllNumbers()

	signals := input.ParseFile("./input.txt")
	fmt.Print("Task 2: Count Simple Numbers\n")
	{
		total := signal.CountSimpleNumbers(signals)
		fmt.Print("  Sum simple digits: ", total, "\n")
	}
	fmt.Print("Task 2: Decode Values\n")
	{
		sum := 0
		for _, pattern := range signals {
			demangler := pattern.FindDemangler()
			value := 0
			for idx, signal := range pattern.O {
				value += signal.Demangle(demangler).Decode() * pow10(len(pattern.O)-idx-1)
			}
			sum += value
		}
		fmt.Printf("  Sum decoded values: %d\n", sum)
	}
}
