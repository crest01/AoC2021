package main

import (
	"advent_of_code/day1/parser"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/wcharczuk/go-chart"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func derivate(vals []float64) []float64 {
	result := make([]float64, len(vals))
	result[0] = 0
	for i := 1; i < len(vals); i++ {
		result[i] = vals[i] - vals[i-1]
	}
	return result
}

func convolve(vals []float64, filt []float64) []float64 {
	result := make([]float64, 0, len(vals))
	for i := 0; i <= (len(vals) - len(filt)); i++ {
		result = append(result, 0)
		for j := 0; j < len(filt); j++ {
			result[i] += vals[i+j] * filt[j]
		}
	}
	return result
}

func main() {
	vals := parser.ParseFile("./input.txt")
	data, err := os.ReadFile("./input.txt")
	check(err)

	for _, entry := range strings.Split(string(data), "\n") {
		if len(entry) == 0 {
			continue
		}
		var value float64
		value, _ = strconv.ParseFloat(entry, 64)
		vals = append(vals, value)
	}

	dvals := derivate(vals)

	num_increased := 0
	for _, val := range dvals {
		if val > 0.0 {
			num_increased++
		}
	}
	fmt.Printf("num_increased raw: %d\n", num_increased)

	filtered := convolve(vals, []float64{1.0, 1.0, 1.0})
	filtered_dvals := derivate(filtered)

	num_increased_filtered := 0
	for _, val := range filtered_dvals {
		if val > 0.0 {
			num_increased_filtered++
		}
	}
	fmt.Printf("num_increased filtered: %d\n", num_increased_filtered)

	xvals := make([]float64, len(vals))
	for i := 0; i < len(vals); i++ {
		xvals[i] = float64(i)
	}

	series := make([]chart.Series, 0)
	series = append(series, chart.ContinuousSeries{Name: "Raw", XValues: xvals, YValues: vals})
	series = append(series, chart.ContinuousSeries{Name: "Filtered", XValues: xvals[:len(filtered)], YValues: filtered})
	series = append(series, chart.ContinuousSeries{Name: "dy Raw", XValues: xvals, YValues: dvals})
	series = append(series, chart.ContinuousSeries{Name: "dy Filtered", XValues: xvals[:len(filtered_dvals)], YValues: filtered_dvals})

	graph := chart.Chart{
		Title:      "Advent Of Code Day 1",
		TitleStyle: chart.StyleShow(),
		Width:      1600,
		Height:     1000,
		XAxis:      chart.XAxis{Style: chart.StyleShow()},
		YAxis:      chart.YAxis{Style: chart.StyleShow()},
		Series:     series}

	graph.Elements = append(graph.Elements, chart.LegendThin(&graph))

	f, err := os.Create("graph.svg")
	check(err)
	defer f.Close()

	w := bufio.NewWriter(f)
	err2 := graph.Render(chart.SVG, w)
	check(err2)

}
