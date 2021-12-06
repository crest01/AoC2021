package main

import (
	"advent_of_code/day3/input"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func setBit(val *uint32, idx uint) {
	mask := uint32(1)
	mask = mask << idx
	*val = *val | mask
}

func isBitSet(val *uint32, idx uint) bool {
	mask := uint32(1)
	mask = mask << idx
	bit_set := *val & mask
	return bit_set == mask
}

func find_gamma_epsilon(values []uint32, num_bits int, max uint32) (uint32, uint32) {
	var gamma uint32
	for i := 0; i < num_bits; i++ {
		num_set := 0
		num_unset := 0
		for _, entry := range values {
			if isBitSet(&entry, uint(i)) {
				num_set++
			} else {
				num_unset++
			}
		}
		if num_set >= num_unset {
			setBit(&gamma, uint(i))
		}
	}
	epsilon := max ^ gamma
	return gamma, epsilon
}

func filter_values(values []uint32, filter uint32, bit uint) []uint32 {
	var mask uint32
	setBit(&mask, bit)
	var result []uint32
	for _, entry := range values {
		if isBitSet(&entry, bit) == isBitSet(&filter, bit) {
			result = append(result, entry)
		} else {
		}
	}
	return result
}

func print_array_binary(values []uint32, bits int) {
	for _, entry := range values {
		fmt.Printf("%012b\n", entry)
	}
}

func main() {

	values, num_bits := input.ParseFile("./input.txt")
	var max uint32
	for i := 0; i < num_bits; i++ {
		setBit(&max, uint(i))
	}

	gamma, epsilon := find_gamma_epsilon(values, num_bits, max)

	fmt.Printf("gamma = %d\n", gamma)
	fmt.Printf("epsilon = %d\n", epsilon)
	fmt.Printf("gamma * epsilon = %d\n", gamma*epsilon)

	var oxygen_value uint32
	oxygen_values := values
	//print_array_binary(oxygen_values, num_bits)
	for found := false; found == false; {
		for i := num_bits - 1; i >= 0; i-- {
			gamma, epsilon = find_gamma_epsilon(oxygen_values, num_bits, max)
			oxygen_values = filter_values(oxygen_values, gamma, uint(i))
			//print_array_binary(oxygen_values, num_bits)
			if len(oxygen_values) == 1 {
				oxygen_value = oxygen_values[0]
				found = true
				break
			}
			// if len(oxygen_values) == 2 {
			// 	oxygen_value = gamma
			// 	found = true
			// 	break
			// }
		}
	}
	var scrubber_value uint32
	scrubber_values := values
	for found := false; found == false; {
		for i := num_bits - 1; i >= 0; i-- {
			gamma, epsilon = find_gamma_epsilon(scrubber_values, num_bits, max)
			scrubber_values = filter_values(scrubber_values, epsilon, uint(i))
			if len(scrubber_values) == 1 {
				scrubber_value = scrubber_values[0]
				found = true
				break
			}
		}
	}

	fmt.Printf("oxygen = %d\n", oxygen_value)
	fmt.Printf("scrubber = %d\n", scrubber_value)
	fmt.Printf("oxygen * scrubber = %d\n", oxygen_value*scrubber_value)
}
