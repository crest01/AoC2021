package input

type Direction uint

const (
	Undefined Direction = iota
	Forward
	AimUp
	AimDown
)

type Instruction struct {
	Dir  Direction
	Size uint
}
