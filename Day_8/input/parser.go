package input

import (
	"advent_of_code/day8/signal"
	"fmt"
	"os"
)

func asSignal(data *[]byte, start int, end int) signal.Signal {
	substr := (*data)[start:end]
	var b byte
	for _, entry := range substr {
		switch entry {
		case 'a':
			b |= signal.A
			break
		case 'b':
			b |= signal.B
			break
		case 'c':
			b |= signal.C
			break
		case 'd':
			b |= signal.D
			break
		case 'e':
			b |= signal.E
			break
		case 'f':
			b |= signal.F
			break
		case 'g':
			b |= signal.G
			break
		default:
			panic("Unknown Signal!")
		}
	}
	return signal.Signal{V: b, S: string(substr)}
}

// Pad string with YYMAXFILL zeroes at the end.
func pad(data *[]byte, fill int) *[]byte {
	padding := make([]byte, fill)
	*data = append(*data, padding...)
	return data
}

func printError(str *[]byte, start int, end int) {
	substr := string((*str)[start:end])
	fmt.Printf("Error: unexpected input '%s' at position %d", substr, start)
	panic("Unexpected Input")
}

func Parse(input *[]byte) []signal.Pattern {
	return lex(input)
}

func ParseFile(filename string) []signal.Pattern {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return Parse(&data)
}
