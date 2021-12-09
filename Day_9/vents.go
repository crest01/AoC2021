package main

import (
	"advent_of_code/day9/input"
	"fmt"
	"sort"
)

type Int2 struct {
	x int
	y int
}

func (a Int2) add(b Int2) Int2 {
	return Int2{a.x + b.x, a.y + b.y}
}

func getNeighbors(data [][]int8, pos Int2) (Int2, Int2, Int2, Int2) {
	up := pos.add(Int2{1, 0})
	down := pos.add(Int2{-1, 0})
	left := pos.add(Int2{0, -1})
	right := pos.add(Int2{0, 1})

	if down.x < 0 {
		down.x = 1
	}
	if left.y < 0 {
		left.y = 1
	}

	size_x := len(data)
	size_y := len(data[0])

	if up.x >= size_x {
		up.x = size_x - 2
	}
	if right.y >= size_y {
		right.y = size_y - 2
	}
	return up, down, left, right
}

func getNeighborVals(data [][]int8, up, down, left, right Int2) (int8, int8, int8, int8) {
	v_up := data[up.x][up.y]
	v_down := data[down.x][down.y]
	v_left := data[left.x][left.y]
	v_right := data[right.x][right.y]
	return v_up, v_down, v_left, v_right
}

func findBasins(data [][]int8) ([]Int2, []int8) {
	size_x := len(data)
	size_y := len(data[0])

	basins := make([]Int2, 0)
	heights := make([]int8, 0)
	for x := 0; x < size_x; x++ {
		for y := 0; y < size_y; y++ {
			value := data[x][y]
			up_c, down_c, left_c, right_c := getNeighbors(data, Int2{x, y})
			up, down, left, right := getNeighborVals(data, up_c, down_c, left_c, right_c)
			if value < up && value < down && value < left && value < right {
				basins = append(basins, Int2{x, y})
				heights = append(heights, value)
			}
		}
	}
	return basins, heights
}

func findBasinSizes(data [][]int8, basins []Int2) []int {
	sizes := make([]int, len(basins))
	for idx := 0; idx < len(basins); idx++ {
		stack := make([]Int2, 0)
		stack = append(stack, basins[idx])
		size := 0
		for len(stack) > 0 {
			coord := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			value := data[coord.x][coord.y]
			if value < 9 && value >= 0 {
				up, down, left, right := getNeighbors(data, coord)
				stack = append(stack, up, down, left, right)
				data[coord.x][coord.y] = -1
				size++
			}
		}
		sizes[idx] = size
	}
	return sizes
}

func main() {

	canvas := input.ParseFile("./input.txt")
	fmt.Print("Task 1: Count Basins\n")
	{
		_, heights := findBasins(canvas)
		fmt.Print("  Found ", len(heights), " basins\n")
		sum := 0
		for _, height := range heights {
			sum = sum + int(height) + 1
		}
		fmt.Print("  Sum: ", sum, "\n")
	}
	fmt.Print("Task 2: Largest Basin Sizes\n")
	{
		basins, _ := findBasins(canvas)
		sizes := findBasinSizes(canvas, basins)
		sort.Ints(sizes)
		fmt.Print("  Found sizes ", sizes, "\n")
		fmt.Print("  Sum three largest: ", sizes[len(sizes)-1]*sizes[len(sizes)-2]*sizes[len(sizes)-3], "\n")

	}
}
