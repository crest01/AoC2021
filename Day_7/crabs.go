package main

import (
	"advent_of_code/day7/input"
	"fmt"
	"sort"
)

func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func CalcBestLocation(crabs *[]int) int {
	sorted_crabs := make([]int, len(*crabs))
	copy(sorted_crabs, *crabs)
	sort.Ints(sorted_crabs)
	index := len(sorted_crabs) / 2
	return (*crabs)[index]
}

func CalcBestLocationTask2(crabs *[]int, epsilon float64) int {
	var sum float64
	for i := range *crabs {
		sum += float64((*crabs)[i])
	}
	avg := sum / float64(len(*crabs))
	return int(avg + epsilon)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func CalcFuelRequirements(crabs *[]int, position int) int {
	fuel_cost := 0
	for i := range *crabs {
		fuel_cost += abs((*crabs)[i] - position)
	}
	return int(fuel_cost)
}

func CalcFuelRequirementsTask2(crabs *[]int, position int) int {
	fuel_cost := 0
	for i := range *crabs {
		n := abs((*crabs)[i] - position)
		fuel_cost += ((n * n) + n) / 2
	}
	return int(fuel_cost)
}

func main() {
	crabby_explosion := input.ParseFile("./input.txt")
	fmt.Print("Task 1: Fuel to best pos\n")
	{
		best_location := CalcBestLocation(&crabby_explosion)
		fmt.Print("  Best Pos: ", best_location, "\n")
		required_fuel := CalcFuelRequirements(&crabby_explosion, best_location)
		fmt.Print("  Required Fuel: ", required_fuel, "\n")
	}
	fmt.Print("Task 2: Fuel to best pos, increasing fuel costs\n")
	{
		best_location_1 := CalcBestLocationTask2(&crabby_explosion, 0.5)
		best_location_2 := CalcBestLocationTask2(&crabby_explosion, -0.5)

		fmt.Print("  Possible Best Pos for Task 2: ", best_location_1, " or ", best_location_2, "\n")
		required_fuel_1 := CalcFuelRequirementsTask2(&crabby_explosion, best_location_1)
		required_fuel_2 := CalcFuelRequirementsTask2(&crabby_explosion, best_location_2)

		fmt.Print("  Required Fuel: ", required_fuel_1, " or ", required_fuel_2, "\n")

		if required_fuel_1 < required_fuel_2 {
			fmt.Print("  Best Pos: ", best_location_1, " with ", required_fuel_1, " required Fuel ")
		} else {
			fmt.Print("  Best Pos: ", best_location_2, " with ", required_fuel_2, " required Fuel ")
		}
	}
}
