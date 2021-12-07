package main

import (
	"advent_of_code/day7/input"
	"testing"
)

func TestCalcBestLocation(t *testing.T) {
	test_string := []byte("16,1,2,0,4,2,7,1,2,14")
	test_crabs := input.Parse(&test_string)
	type args struct {
		crabs *[]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Calc Best Location",
			args: args{crabs: &test_crabs},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcBestLocation(tt.args.crabs); got != tt.want {
				t.Errorf("CalcBestLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalcFuelRequirements(t *testing.T) {
	test_string := []byte("16,1,2,0,4,2,7,1,2,14")
	test_crabs := input.Parse(&test_string)
	type args struct {
		crabs    *[]int
		position int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Calc Fuel",
			args: args{crabs: &test_crabs, position: 2},
			want: 37,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcFuelRequirements(tt.args.crabs, tt.args.position); got != tt.want {
				t.Errorf("CalcFuelRequirements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalcBestLocationTask2(t *testing.T) {
	test_string := []byte("16,1,2,0,4,2,7,1,2,14")
	test_crabs := input.Parse(&test_string)
	type args struct {
		crabs   *[]int
		epsilon float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Calc Best Location Task 2, rounded up",
			args: args{crabs: &test_crabs, epsilon: 0.5},
			want: 5,
		},
		{
			name: "Calc Best Location Task 2, rounded down",
			args: args{crabs: &test_crabs, epsilon: -0.5},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcBestLocationTask2(tt.args.crabs, tt.args.epsilon); got != tt.want {
				t.Errorf("CalcBestLocationTask2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalcFuelRequirementsTask2(t *testing.T) {
	test_string := []byte("16,1,2,0,4,2,7,1,2,14")
	test_crabs := input.Parse(&test_string)
	type args struct {
		crabs    *[]int
		position int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Calc Fuel Task 2",
			args: args{crabs: &test_crabs, position: 2},
			want: 206,
		},
		{
			name: "Calc Fuel Task 2",
			args: args{crabs: &test_crabs, position: 5},
			want: 168,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcFuelRequirementsTask2(tt.args.crabs, tt.args.position); got != tt.want {
				t.Errorf("CalcFuelRequirementsTask2() = %v, want %v", got, tt.want)
			}
		})
	}
}
