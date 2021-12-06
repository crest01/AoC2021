package parser

import (
	"os"
	"strconv"
	"strings"
)

func ParseFile(filename string) []float64 {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var vals []float64
	for _, entry := range strings.Split(string(data), "\n") {
		if len(entry) == 0 {
			continue
		}
		var value float64
		value, _ = strconv.ParseFloat(entry, 64)
		vals = append(vals, value)
	}
	return vals
}
