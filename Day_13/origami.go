package main

import (
	"advent_of_code/day13/input"
	"advent_of_code/day13/sparse_field"
	"advent_of_code/day13/utils"
	"fmt"
	"time"
)

func parse_test(filename string) {
	defer utils.Stopwatch(time.Now(), "Parsing")
	dots, folds := input.ParseFile(filename)
	fmt.Printf("Got %d dots and %d folds from file %s\n", len(dots), len(folds), filename)
}

func task1(filename string) {
	dots, folds := input.ParseFile(filename)

	defer utils.Stopwatch(time.Now(), "Task 1")
	field := sparse_field.MakeField(dots)
	fmt.Print("Task 1: Find Number of dots after folding\n")
	field.Fold(folds[0])
	fmt.Printf("Found %d dots after the first fold\n", field.NumPoints())

}

func task2(filename string) {
	defer utils.Stopwatch(time.Now(), "Task 2")
	fmt.Print("Task 2: fully fold fields\n")
	dots, folds := input.ParseFile(filename)
	field := sparse_field.MakeField(dots)

	for i := 0; i < len(folds); i++ {
		field.Fold(folds[i])
		fmt.Printf("Folded with (%s)\n", folds[i].AsString())
	}
	fmt.Printf("Field after %d folds: \n", len(folds))
	field.Print()
}

func main() {
	filename := "./input.txt"

	parse_test(filename)

	task1(filename)

	task2(filename)

}
