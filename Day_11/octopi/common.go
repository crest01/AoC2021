package octopi

type Int2 struct {
	x int
	y int
}

func make_int2(x int, y int) Int2 {
	return Int2{x, y}
}

func (a Int2) Add(b Int2) Int2 {
	return Int2{a.x + b.x, a.y + b.y}
}

type Octopus struct {
	pos     Int2
	energy  int
	virtual bool
}

type Stack struct {
	stack []Int2
}

func (s *Stack) Push(o Int2) {
	s.stack = append(s.stack, o)
}

func (s *Stack) Pop() Int2 {
	v := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return v
}

func (s *Stack) Empty() bool {
	return len(s.stack) == 0
}

type Field struct {
	area [][]Octopus
	size Int2
}

func MakeField(data [][]int) Field {
	var f Field
	size_x := len(data)
	size_y := len(data[0])
	f.Init(Int2{size_x, size_y})
	for x := 0; x < f.size.x; x++ {
		for y := 0; y < f.size.y; y++ {
			f.SetEnergy(Int2{x, y}, data[x][y])
		}
	}
	return f
}

func (f *Field) Init(size Int2) {
	f.area = make([][]Octopus, size.x)
	for i := 0; i < size.y; i++ {
		f.area[i] = make([]Octopus, size.y)
	}
	f.size = size
}

func (f *Field) Add(pos Int2, energy int) {
	f.area[pos.x][pos.y].energy += energy
}

func (f *Field) Set(pos Int2, energy int) {
	f.area[pos.x][pos.y].energy = energy
}

func (f *Field) Get(pos Int2) Octopus {
	out_of_bounds := false
	if pos.x < 0 || pos.x >= f.size.x {
		out_of_bounds = true
	}
	if pos.y < 0 || pos.y >= f.size.y {
		out_of_bounds = true
	}
	if out_of_bounds {
		return Octopus{pos: pos, energy: 0, virtual: true}
	}
	f.area[pos.x][pos.y].pos = pos
	f.area[pos.x][pos.y].virtual = false
	return f.area[pos.x][pos.y]
}

func (f *Field) At(pos Int2) *Octopus {
	return &f.area[pos.x][pos.y]
}

func (f *Field) GetEnergy(pos Int2) int {
	return f.Get(pos).energy
}

func (f *Field) SetEnergy(pos Int2, energy int) {
	(*(*f).At(pos)).energy = energy
}

func (f *Field) GetNeighbors(pos Int2) []Octopus {
	up := Int2{1, 0}
	down := Int2{-1, 0}
	left := Int2{0, -1}
	right := Int2{0, 1}

	result := make([]Octopus, 0)

	result = append(result, f.Get(pos.Add(up).Add(left)))
	result = append(result, f.Get(pos.Add(up)))
	result = append(result, f.Get(pos.Add(up).Add(right)))
	result = append(result, f.Get(pos.Add(left)))
	result = append(result, f.Get(pos.Add(right)))
	result = append(result, f.Get(pos.Add(down).Add(left)))
	result = append(result, f.Get(pos.Add(down)))
	result = append(result, f.Get(pos.Add(down).Add(right)))
	return result
}

func (f *Field) ToString() []byte {
	result := make([]byte, 0, (f.size.x+1)*f.size.y)
	for x := 0; x < f.size.x; x++ {
		for y := 0; y < f.size.y; y++ {
			o := f.At(Int2{x, y})
			if o.energy > 9 {
				result = append(result, 'X')
			} else if o.energy < 0 {
				result = append(result, 'E')
			} else {
				result = append(result, byte(o.energy+48))
			}
		}
		result = append(result, '\n')
	}
	return result
}

func SimulateStep(f Field) []Int2 {

	blinks := make([]Int2, 0)
	var stack Stack

	// Increase everyones energy
	for x := 0; x < f.size.x; x++ {
		for y := 0; y < f.size.y; y++ {
			pos := Int2{x, y}
			energy := f.GetEnergy(pos)
			if energy == 9 {
				stack.Push(pos)
			}
			f.Add(pos, 1)
		}
	}
	// propagate blinks
	for !stack.Empty() {
		pos := stack.Pop()
		neighbors := f.GetNeighbors(pos)
		for _, n := range neighbors {
			if n.virtual { // outside the bounds
				continue
			} else if n.energy > 9 { // already blinking
				continue
			} else if n.energy == 9 { // about to blink
				stack.Push(n.pos)
			}
			f.Add(n.pos, 1)
		}
	}

	// count blinks
	for x := 0; x < f.size.x; x++ {
		for y := 0; y < f.size.y; y++ {
			o := f.At(Int2{x, y})
			if o.energy > 9 {
				(*o).energy = 0
				blinks = append(blinks, Int2{x, y})
			}
		}
	}
	return blinks
}
