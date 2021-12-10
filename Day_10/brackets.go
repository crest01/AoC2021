package main

import (
	"advent_of_code/day10/input"
	"fmt"
	"sort"
)

func getCorruptScore(bracket byte) int {
	switch bracket {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	}
	return 0
}

func getCompleteScore(bracket byte) int {
	switch bracket {
	case ')':
		return 1
	case ']':
		return 2
	case '}':
		return 3
	case '>':
		return 4
	}
	return 0
}

func getMedian(n []int) int {
	sort.Ints(n) // sort the numbers
	mNumber := len(n) / 2
	return n[mNumber]
}

func main() {

	raw_lines, corrupt_chars, corrupt_line_num, raw_stacks := input.ParseFile("./input.txt")

	lines := make([][]byte, 0)
	stacks := make([][]byte, 0)

	corrupt_lines := make([][]byte, 0)

	j := 0
	for i := 0; i < len(raw_lines); i++ {
		corrupt := -1
		if j < len(corrupt_line_num) {
			corrupt = corrupt_line_num[j]
		}
		if i == corrupt {
			corrupt_lines = append(corrupt_lines, raw_lines[i])
			j++
		} else {
			lines = append(lines, raw_lines[i])
			stacks = append(stacks, raw_stacks[i])
		}
	}

	fmt.Print("Task 1: Find corrupt Lines\n")
	{
		score := 0
		for i := 0; i < len(corrupt_lines); i++ {
			fmt.Printf("%d: %s -> '%c' is wrong\n", corrupt_line_num[i], corrupt_lines[i], corrupt_chars[i])
			score += getCorruptScore(corrupt_chars[i])
		}
		fmt.Printf("Score: %d\n", score)

	}
	fmt.Print("Task 2: autocomplete closing brackets\n")
	{
		scores := make([]int, 0)
		for i := 0; i < len(lines); i++ {
			line := lines[i]
			stack := stacks[i]
			appended := make([]byte, 0)
			score := 0
			for j := len(stack) - 1; j >= 0; j-- {
				closing_bracket := input.GetClosingBracket(stack[j])
				appended = append(appended, closing_bracket)
				score *= 5
				score += getCompleteScore(closing_bracket)
			}
			scores = append(scores, score)
			fmt.Printf("'%s' + '%s': %d Points\n", line, appended, score)
		}
		fmt.Printf("Final Score: %d", getMedian(scores))
	}
}
