package signal

import "fmt"

const (
	A byte = 0b0000001
	B      = 0b0000010
	C      = 0b0000100
	D      = 0b0001000
	E      = 0b0010000
	F      = 0b0100000
	G      = 0b1000000
)

func AllSegments() []byte {
	return []byte{A, B, C, D, E, F, G}
}

func getKeys() []byte {
	keys := make([]byte, 0, 10)
	keys = append(keys, A|B|C|E|F|G)
	keys = append(keys, C|F)
	keys = append(keys, A|C|D|E|G)
	keys = append(keys, A|C|D|F|G)
	keys = append(keys, B|C|D|F)
	keys = append(keys, A|B|D|F|G)
	keys = append(keys, A|B|D|E|F|G)
	keys = append(keys, A|C|F)
	keys = append(keys, A|B|C|D|E|F|G)
	keys = append(keys, A|B|C|D|F|G)
	return keys
}

func decoderMap() map[byte]int {
	decoder := map[byte]int{}
	keys := getKeys()
	for value, key := range keys {
		decoder[key] = value
	}
	return decoder
}

type Signal struct {
	V byte
	S string
}

func (s Signal) Decode() int {
	m := decoderMap()
	return m[s.V]
}

func (s Signal) Demangle(mangle_map map[byte]byte) Signal {
	return Signal{V: mangle_map[s.V]}
}

func PrintAllNumbers() {
	keys := getKeys()
	for _, key := range keys {
		signal := Signal{V: key}
		signal.Print()
	}
}

func (s Signal) Print() {
	a := ' '
	if (s.V & A) == A {
		a = 'x'
	}
	b := ' '
	if (s.V & B) == B {
		b = 'x'
	}
	c := ' '
	if (s.V & C) == C {
		c = 'x'
	}
	d := ' '
	if (s.V & D) == D {
		d = 'x'
	}
	e := ' '
	if (s.V & E) == E {
		e = 'x'
	}
	f := ' '
	if (s.V & F) == F {
		f = 'x'
	}
	g := ' '
	if (s.V & G) == G {
		g = 'x'
	}
	fmt.Printf("Number %d: \n", s.Decode())
	fmt.Printf("__________\n")
	fmt.Printf("|  %c%c%c%c  |\n", a, a, a, a)
	fmt.Printf("| %c    %c |\n", b, c)
	fmt.Printf("| %c    %c |\n", b, c)
	fmt.Printf("| %c    %c |\n", b, c)
	fmt.Printf("| %c    %c |\n", b, c)
	fmt.Printf("|  %c%c%c%c  |\n", d, d, d, d)
	fmt.Printf("| %c    %c |\n", e, f)
	fmt.Printf("| %c    %c |\n", e, f)
	fmt.Printf("| %c    %c |\n", e, f)
	fmt.Printf("| %c    %c |\n", e, f)
	fmt.Printf("|  %c%c%c%c  |\n", g, g, g, g)
	fmt.Printf("__________\n")

}

func Combine(V []byte) Signal {
	var s Signal
	for _, v := range V {
		s.V |= v
	}
	return s
}

func (s Signal) Deconstruct() []byte {
	var result []byte
	if (s.V & A) == A {
		result = append(result, A)
	}
	if (s.V & B) == B {
		result = append(result, B)
	}
	if (s.V & C) == C {
		result = append(result, C)
	}
	if (s.V & D) == D {
		result = append(result, D)
	}
	if (s.V & E) == E {
		result = append(result, E)
	}
	if (s.V & F) == F {
		result = append(result, F)
	}
	if (s.V & G) == G {
		result = append(result, G)
	}
	return result
}

type Pattern struct {
	S   []Signal
	O   []Signal
	O_d int
}

func NumSetBits(n byte) int {
	count := 0
	for n != 0 {
		if n&1 == 1 {
			count++
		}
		n >>= 1
	}
	return count
}

func CountSimpleNumbers(patterns []Pattern) int {
	total := 0
	for _, pattern := range patterns {
		var count_2, count_3, count_4, count_7 int
		for _, signal := range pattern.O {
			if NumSetBits(signal.V) == 2 {
				count_2++
			}
			if NumSetBits(signal.V) == 3 {
				count_3++
			}
			if NumSetBits(signal.V) == 4 {
				count_4++
			}
			if NumSetBits(signal.V) == 7 {
				count_7++
			}
		}
		line_sum := count_2 + count_3 + count_4 + count_7
		total += line_sum
	}
	return total
}

func (p Pattern) Print() {
	for _, signal := range p.S {
		fmt.Printf("%s ", signal.S)
	}
	fmt.Printf("| ")
	for _, signal := range p.O {
		fmt.Printf("%s ", signal.S)
	}
	fmt.Printf("\n")
}

func (p Pattern) FindSimpleNumbers() (Signal, Signal, Signal, Signal) {
	var one, four, seven, eight Signal
	for _, signal := range p.S {
		if NumSetBits(signal.V) == 2 {
			one = signal
		}
		if NumSetBits(signal.V) == 4 {
			four = signal
		}
		if NumSetBits(signal.V) == 3 {
			seven = signal
		}
		if NumSetBits(signal.V) == 7 {
			eight = signal
		}
	}
	return one, four, seven, eight
}

func (p Pattern) Find(segment_len int) []Signal {
	var result []Signal
	for _, signal := range p.S {
		if NumSetBits(signal.V) == segment_len {
			result = append(result, signal)
		}
	}
	return result
}

func DeconstructBasicParts(one Signal, four Signal, seven Signal, eight Signal) (Signal, Signal, Signal, Signal) {
	var cf, bd, a, eg Signal
	cf = one
	bd = four.Difference(one)
	a = seven.Difference(one)
	eg = eight.Difference(four.Union(seven))
	return cf, bd, a, eg
}

func (p Pattern) FindDemangler() map[byte]byte {
	one, four, seven, eight := p.FindSimpleNumbers()
	cf, bd, _, eg := DeconstructBasicParts(one, four, seven, eight)

	var two, three, five Signal
	fiveSegments := p.Find(5)
	if len(fiveSegments) != 3 {
		panic("More than three Five-Len Segments!")
	}
	found := 0
	for _, entry := range fiveSegments {
		if entry.FullyContains(eg) {
			two = entry
			found++
		} else if entry.FullyContains(cf) {
			three = entry
			found++
		} else if entry.FullyContains(bd) {
			five = entry
			found++
		}
	}
	if found != 3 {
		panic("Not all five-len numbers found")
	}

	sixSegments := p.Find(6)
	if len(sixSegments) != 3 {
		panic("More than three Six-Len Segments!")
	}

	var six, nine, zero Signal
	found = 0
	for _, entry := range sixSegments {
		if entry.FullyContains(eg.Union(bd)) {
			six = entry
			found++
		} else if entry.FullyContains(cf.Union(bd)) {
			nine = entry
			found++
		} else if entry.FullyContains(cf.Union(eg)) {
			zero = entry
			found++
		}
	}
	if found != 3 {
		panic("Not all six-len numbers found")
	}

	keys := getKeys()
	demangler := make(map[byte]byte)

	demangler[zero.V] = keys[0]
	demangler[one.V] = keys[1]
	demangler[two.V] = keys[2]
	demangler[three.V] = keys[3]
	demangler[four.V] = keys[4]
	demangler[five.V] = keys[5]
	demangler[six.V] = keys[6]
	demangler[seven.V] = keys[7]
	demangler[eight.V] = keys[8]
	demangler[nine.V] = keys[9]
	return demangler
}

func (a Signal) Union(b Signal) Signal {
	var result Signal
	result.V = a.V | b.V
	return result
}

func (a Signal) Intersect(b Signal) Signal {
	var result Signal
	result.V = a.V & b.V
	return result
}

func (a Signal) Difference(b Signal) Signal {
	var result Signal
	result.V = a.V & ^b.V
	return result
}

func (a Signal) SymDiff(b Signal) Signal {
	var result Signal
	result = a.Union(b).Difference(a.Intersect(b))
	return result
}

func (a Signal) Equal(b Signal) bool {
	return a.V == b.V
}

func (a Signal) FullyContains(b Signal) bool {
	return a.Intersect(b).Equal(b)
}
