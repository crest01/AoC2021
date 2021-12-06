package main

import (
	"advent_of_code/day2/input"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type float3 struct {
	x float64
	y float64
	z float64
}

func (a *float3) Add(v float3) {
	a.x += v.x
	a.y += v.y
	a.z += v.z
}

type Submarine struct {
	pos float3
	aim float64
}

func (sub *Submarine) AimDown(v uint) {
	sub.aim -= float64(v)
	fmt.Printf("down = %d, aim = %f\n", v, sub.aim)
}

func (sub *Submarine) AimUp(v uint) {
	sub.aim += float64(v)
	fmt.Printf("up = %d, aim = %f\n", v, sub.aim)
}

func (sub *Submarine) Forward(v uint) {
	sub.pos.x += float64(v)
	sub.pos.z += sub.aim * float64(v)
	fmt.Printf("forward = %d, pos = (%f, %f, %f)\n", v, sub.pos.x, sub.pos.y, sub.pos.z)
}

func (sub *Submarine) Apply(cmd input.Instruction) {
	switch {
	case cmd.Dir == input.Forward:
		sub.Forward(cmd.Size)
	case cmd.Dir == input.AimDown:
		sub.AimDown(cmd.Size)
	case cmd.Dir == input.AimUp:
		sub.AimUp(cmd.Size)
	default:
		panic("undefined instruction")
	}
}

func main() {

	instructions := input.ParseFile("./input.txt")

	var sub Submarine
	for _, entry := range instructions {
		sub.Apply(entry)
	}

	fmt.Printf("final pos = (%f, %f, %f)\n", sub.pos.x, sub.pos.y, sub.pos.z)
	fmt.Printf("pos.x * (-pos.z) = %f\n", sub.pos.x*(-sub.pos.z))

	// series := make([]chart.Series, 0)
	// series = append(series, chart.ContinuousSeries{Name: "Raw", XValues: xvals, YValues: vals})
	// series = append(series, chart.ContinuousSeries{Name: "Filtered", XValues: xvals[:len(filtered)], YValues: filtered})
	// series = append(series, chart.ContinuousSeries{Name: "dy Raw", XValues: xvals, YValues: dvals})
	// series = append(series, chart.ContinuousSeries{Name: "dy Filtered", XValues: xvals[:len(filtered_dvals)], YValues: filtered_dvals})

	// graph := chart.Chart{
	// 	Title:      "Advent Of Code Day 1",
	// 	TitleStyle: chart.StyleShow(),
	// 	Width:      1600,
	// 	Height:     1000,
	// 	XAxis:      chart.XAxis{Style: chart.StyleShow()},
	// 	YAxis:      chart.YAxis{Style: chart.StyleShow()},
	// 	Series:     series}

	// graph.Elements = append(graph.Elements, chart.LegendThin(&graph))

	// f, err := os.Create("graph.svg")
	// check(err)
	// defer f.Close()

	// w := bufio.NewWriter(f)
	// err2 := graph.Render(chart.SVG, w)
	// check(err2)

}
