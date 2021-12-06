package main

import (
	"advent_of_code/day6/input"
	"fmt"
)

type FishOp func(fishes uint64, output *[]uint64)

func SimulateDay(starter []uint64) []uint64 {

	result := make([]uint64, len(starter))
	for i := 8; i > 0; i-- {
		result[i-1] += starter[i]
	}
	result[6] += starter[0]
	result[8] += starter[0]
	return result
}

func main() {

	fmt.Print("Task 1: 80 days\n")
	{

		fishes_for_days := input.ParseFile("./input.txt")
		fmt.Print("Start Fishes: ", fishes_for_days, "\n")
		days := 80
		for i := 0; i < days; i++ {
			fishes_for_days = SimulateDay(fishes_for_days)
			fmt.Print("Day ", i, " Fishes: ", fishes_for_days, "\n")
			sum := uint64(0)
			for _, val := range fishes_for_days {
				sum += val
			}
			fmt.Print("Total fishes: ", sum, "\n")
		}
	}
	fmt.Print("Task 2: 256 days\n")
	{
		fishes_for_days := input.ParseFile("./input.txt")
		fmt.Print("Start Fishes: ", fishes_for_days, "\n")
		days := 256
		for i := 0; i < days; i++ {
			fishes_for_days = SimulateDay(fishes_for_days)
			fmt.Print("Day ", i, " Fishes: ", fishes_for_days, "\n")
			sum := uint64(0)
			for _, val := range fishes_for_days {
				sum += val
			}
			fmt.Print("Total fishes: ", sum, "\n")
		}

	}
}
