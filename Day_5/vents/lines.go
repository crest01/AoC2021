package vents

type Int2 struct {
	X int
	Y int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func min2(a, b Int2) Int2 {
	return Int2{X: min(a.X, b.X), Y: min(a.Y, b.Y)}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func max2(a, b Int2) Int2 {
	return Int2{X: max(a.X, b.X), Y: max(a.Y, b.Y)}
}

type Line struct {
	P0 Int2
	P1 Int2
}

func (l *Line) IsStraight() bool {
	if l.P0.X == l.P1.X || l.P0.Y == l.P1.Y {
		return true
	}
	return false
}

func (l *Line) Rasterize(canvas *[][]uint8) {
	// implemented straight from WP pseudocode
	x0 := l.P0.X
	x1 := l.P1.X
	y0 := l.P0.Y
	y1 := l.P1.Y

	dx := x1 - x0
	if dx < 0 {
		dx = -dx
	}
	dy := y1 - y0
	if dy < 0 {
		dy = -dy
	}
	var sx, sy int
	if x0 < x1 {
		sx = 1
	} else {
		sx = -1
	}
	if y0 < y1 {
		sy = 1
	} else {
		sy = -1
	}
	err := dx - dy

	for {
		(*canvas)[x0][y0]++
		if x0 == x1 && y0 == y1 {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
}

func CalcAABB(lines *[]Line) (Int2, Int2) {
	min := (*lines)[0].P0
	max := (*lines)[0].P0

	for idx := range *lines {
		entry := &(*lines)[idx]
		min = min2((*entry).P0, min)
		min = min2((*entry).P1, min)
		max = max2((*entry).P0, max)
		max = max2((*entry).P1, max)
	}
	return min, max
}

func MakeInt2(numbers *[]int) Int2 {
	n := len(*numbers) - 1
	x := (*numbers)[n]
	y := (*numbers)[n-1]
	*numbers = (*numbers)[:n-1]
	return Int2{X: x, Y: y}
}

func MakeLine(coords *[]Int2) Line {
	n := len(*coords) - 1
	p0 := (*coords)[n]
	p1 := (*coords)[n-1]
	*coords = (*coords)[:n-1]
	return Line{P0: p0, P1: p1}
}
