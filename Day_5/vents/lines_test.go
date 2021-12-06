package vents

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMakeInt2(t *testing.T) {
	var input_stack []int
	input_stack = append(input_stack, 5, 4, 3, 2, 1)

	type args struct {
		numbers *[]int
	}
	tests := []struct {
		name string
		args args
		want Int2
	}{
		{
			name: "Create Int2 1",
			args: args{&input_stack},
			want: Int2{X: 1, Y: 2},
		},
		{
			name: "Create Int2 2",
			args: args{&input_stack},
			want: Int2{X: 3, Y: 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeInt2(tt.args.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeInt2() = %v, want %v", got, tt.want)
			}
		})
	}
	t.Run("Create Int2 Memory", func(t *testing.T) {
		if len(input_stack) != 1 {
			t.Errorf("MakeInt2() Memory error: len(input_stack)=%d, want %d", len(input_stack), 1)
		} else {
			if input_stack[0] != 5 {
				t.Errorf("MakeInt2() Memory error: input_stack[0]=%d, want %d", input_stack[0], 5)
			}
		}
	})
}

func TestMakeLine(t *testing.T) {
	var input_stack []Int2
	input_stack = append(input_stack, Int2{1, 0}, Int2{3, 2}, Int2{5, 4})
	type args struct {
		coords *[]Int2
	}
	tests := []struct {
		name string
		args args
		want Line
	}{
		{
			name: "Create Line",
			args: args{&input_stack},
			want: Line{P0: Int2{X: 5, Y: 4}, P1: Int2{X: 3, Y: 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeLine(tt.args.coords); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeLine() = %v, want %v", got, tt.want)
			}
		})
	}
	t.Run("Create Line Memory", func(t *testing.T) {
		if len(input_stack) != 1 {
			t.Errorf("MakeLine() Memory error: len(input_stack)=%d, want %d", len(input_stack), 1)
		} else {
			if !reflect.DeepEqual(input_stack[0], Int2{1, 0}) {
				t.Errorf("MakeLine() Memory error: input_stack[0]=(%d, %d), want (%d, %d)", input_stack[0].X, input_stack[0].Y, 1, 0)
			}
		}
	})
}

func print_canvas(canvas *[][]uint8, res_x int, res_y int) {
	fmt.Print("-------\n")
	for x := 0; x < res_x; x++ {
		fmt.Print("|")
		for y := 0; y < res_y; y++ {
			val := (*canvas)[x][y]
			if val > 0 {
				fmt.Printf("%d", val)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("|\n")
	}
	fmt.Print("-------\n")
}

func TestLine_Rasterize(t *testing.T) {
	canvas := make([][]uint8, 6)
	for i := range canvas {
		canvas[i] = make([]uint8, 6)
	}

	line_1 := Line{P0: Int2{X: 0, Y: 0}, P1: Int2{X: 5, Y: 0}}
	line_2 := Line{P0: Int2{X: 5, Y: 0}, P1: Int2{X: 5, Y: 5}}
	line_3 := Line{P0: Int2{X: 0, Y: 0}, P1: Int2{X: 5, Y: 5}}

	type args struct {
		canvas *[][]uint8
	}
	tests := []struct {
		name string
		l    *Line
		args args
	}{
		{
			name: "vertical line",
			l:    &line_1,
			args: args{canvas: &canvas},
		},
		{
			name: "horizontal line",
			l:    &line_2,
			args: args{canvas: &canvas},
		},
		{
			name: "diagonal line",
			l:    &line_3,
			args: args{canvas: &canvas},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.l.Rasterize(tt.args.canvas)
			print_canvas(tt.args.canvas, 6, 6)
		})
	}
}
