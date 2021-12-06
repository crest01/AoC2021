package input

import (
	"os"
	"strconv"
	"strings"
)

func ParseFile(filename string) ([]uint32, int) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var values []uint32
	num_bits := 0
	for _, entry := range strings.Split(string(data), "\n") {
		if len(entry) == 0 {
			continue
		}
		num_bits = len(entry)
		val, err := strconv.ParseUint(entry, 2, 32)
		if err != nil {
			panic(err)
		}
		values = append(values, uint32(val))
	}

	return values, num_bits
}
