package sparse_field

import (
	"fmt"

	"github.com/emirpasic/gods/sets/hashset"
)

func abs(v int) int {
	if v < 0 {
		return v * -1
	}
	return v
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func min2(a, b Int2) Int2 {
	return Int2{min(a.x, b.x), min(a.y, b.y)}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func max2(a, b Int2) Int2 {
	return Int2{max(a.x, b.x), max(a.y, b.y)}
}

type Int2 struct {
	x, y int
}

func MakeInt2(x, y int) Int2 {
	return Int2{x, y}
}

func (v Int2) FoldUp(yaxis int) Int2 {
	y := (yaxis - v.y) + yaxis
	return Int2{v.x, y}
}

func (v Int2) FoldLeft(xaxis int) Int2 {
	x := (xaxis - v.x) + xaxis
	return Int2{x, v.y}
}

func (p Int2) AsString() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

type Field struct {
	dots []Int2
}

func MakeField(dots []Int2) *Field {
	f := &Field{}
	for _, v := range dots {
		f.AddPoint(v)
	}
	return f
}

func (f *Field) Fold(fold Int2) {
	unchanged_dots := make([]Int2, 0, len(f.dots))
	set := hashset.New()
	new_dots := make([]Int2, 0, len(f.dots))
	if fold.x != 0 {
		for _, v := range f.dots {
			if v.x < fold.x {
				unchanged_dots = append(unchanged_dots, v)
				set.Add(v.AsString())
			} else {
				new_dots = append(new_dots, v.FoldLeft(fold.x))
			}
		}
	} else if fold.y != 0 {
		for _, v := range f.dots {
			if v.y < fold.y {
				unchanged_dots = append(unchanged_dots, v)
				set.Add(v.AsString())
			} else {
				new_dots = append(new_dots, v.FoldUp(fold.y))
			}
		}
	}
	for _, v := range new_dots {
		s := v.AsString()
		if set.Contains(s) {
			continue
		}
		set.Add(s)
		unchanged_dots = append(unchanged_dots, v)
	}
	(*f).dots = unchanged_dots
}

func (f *Field) AddPoint(p Int2) {
	f.dots = append(f.dots, p)
}

func (f *Field) NumPoints() int {
	return len(f.dots)
}

func (f *Field) Print() {
	values := f.dots
	var aabb_min, aabb_max Int2
	aabb_min = values[0]
	aabb_max = values[0]
	for i := 1; i < len(values); i++ {
		aabb_min = min2(aabb_min, values[i])
		aabb_max = max2(aabb_max, values[i])
	}

	width := aabb_max.x
	height := aabb_max.y

	field := make([][]bool, width+1)
	for i := 0; i <= width; i++ {
		field[i] = make([]bool, height+1)
	}
	for _, value := range values {
		field[value.x][value.y] = true
	}

	result := make([]byte, 0, (width+1)*(height))

	for y := 0; y <= height; y++ {
		for x := 0; x <= width; x++ {
			if field[x][y] {
				result = append(result, '#')
			} else {
				result = append(result, '.')
			}
		}
		result = append(result, '\n')
	}

	fmt.Printf("Field: \n%s\n", result)
}
