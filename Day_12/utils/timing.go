package utils

import (
	"fmt"
	"time"
)

func Stopwatch(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}
