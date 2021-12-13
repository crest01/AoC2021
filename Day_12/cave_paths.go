package main

import (
	"advent_of_code/day12/graph"
	"advent_of_code/day12/input"
	"advent_of_code/day12/utils"
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"

	"github.com/emirpasic/gods/sets/hashset"
)

func parse_test(filename string) {
	defer utils.Stopwatch(time.Now(), "Parsing")
	graph_edges := input.ParseFile(filename)
	g := graph.MakeGraph(graph_edges)

	f, e := os.Create("./graph.dot")
	if e != nil {
		panic(e)
	}
	defer f.Close()
	fmt.Fprintf(f, "%s", g.ToDotString())
}

func task1(filename string) {
	graph_edges := input.ParseFile(filename)
	g := graph.MakeGraph(graph_edges)

	defer utils.Stopwatch(time.Now(), "Task 1")
	fmt.Print("Task 1: Find unique paths\n")

	unique_paths_1 := g.FindAllPaths()
	fmt.Printf("Found %d unique paths\n", len(unique_paths_1))
	unique_paths_2 := g.FindAllPathsReverse()
	fmt.Printf("Found %d unique paths\n", len(unique_paths_2))

}

func task2(filename string) {
	graph_edges := input.ParseFile(filename)
	g := graph.MakeGraph(graph_edges)

	defer utils.Stopwatch(time.Now(), "Task 2")
	fmt.Print("Task 2: Find unique paths, visiting one small cave twice\n")
	graphs := g.CreateVisitPermutations()
	unique_paths := hashset.New()
	for idx, p_g := range graphs {
		paths := p_g.FindAllPaths()
		fmt.Printf("Permutation %d: Found %d unique paths", idx, len(paths))
		old_len := unique_paths.Size()
		for _, path := range paths {
			// key := graph.GetPathKey(path)
			// unique_paths.Add(key)
			unique_paths.Add(path)
		}
		fmt.Printf(" -> %d were new\n", unique_paths.Size()-old_len)
	}
	fmt.Printf("Found %d unique paths\n", unique_paths.Size())
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	filename := "./input.txt"

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			panic(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	parse_test(filename)

	task1(filename)

	task2(filename)

}
