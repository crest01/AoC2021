package main

import (
	"advent_of_code/day6/input"
	"reflect"
	"testing"
)

func TestSimulateDay(t *testing.T) {

	day_0 := "3,4,3,1,2"
	day_1 := "2,3,2,0,1"
	day_2 := "1,2,1,6,0,8"
	day_3 := "0,1,0,5,6,7,8"
	day_4 := "6,0,6,4,5,6,7,8,8"
	day_5 := "5,6,5,3,4,5,6,7,7,8"

	type args struct {
		starter []uint64
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{
			name: "day 1",
			args: args{starter: input.Parse(day_0)},
			want: input.Parse(day_1),
		},
		{
			name: "day 2",
			args: args{starter: input.Parse(day_1)},
			want: input.Parse(day_2),
		},
		{
			name: "day 3",
			args: args{starter: input.Parse(day_2)},
			want: input.Parse(day_3),
		},
		{
			name: "day 4",
			args: args{starter: input.Parse(day_3)},
			want: input.Parse(day_4),
		},
		{
			name: "day 5",
			args: args{starter: input.Parse(day_4)},
			want: input.Parse(day_5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SimulateDay(tt.args.starter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimulateDay() = %v, want %v", got, tt.want)
			}
		})
	}
}
