package graph

import (
	"fmt"
	"hash/maphash"
	"strings"
	"unicode"

	"github.com/emirpasic/gods/sets/hashset"
	"github.com/emirpasic/gods/stacks/arraystack"

	"github.com/emicklei/dot"
)

type Graph struct {
	nodes           map[string]int
	r_nodes         map[int]string
	visits          []int
	small_caves     []bool
	edges           [][]int
	reachable_nodes map[int][]int
	start           int
	end             int
}

func MakeGraph(edges [][]string) *Graph {

	nodes := hashset.New()
	for _, entry := range edges {
		for _, node := range entry {
			nodes.Add(node)
		}
	}
	g := &Graph{}
	g.Init(nodes.Size())

	for _, entry := range edges {
		start := entry[0]
		end := entry[1]
		g.AddEdge(start, end)
		idx1 := g.GetNodeIdx(start)
		idx2 := g.GetNodeIdx(end)
		if start == "start" {
			g.start = idx1
		}
		if end == "end" {
			g.end = idx2
		}
		if unicode.IsLower(rune(start[0])) {
			g.small_caves[idx1] = true
		}
		if unicode.IsLower(rune(end[0])) {
			g.small_caves[idx2] = true
		}
	}
	g.PrecalcReachableNodes()
	return g
}

func (g *Graph) Duplicate() *Graph {
	result := &Graph{}
	*result = *g
	(*result).visits = make([]int, len(g.visits))
	for i := 0; i < len(g.nodes); i++ {
		(*result).visits[i] = (*g).visits[i]
	}
	return result
}

func (g *Graph) Init(max_num_nodes int) {

	g.nodes = make(map[string]int, max_num_nodes)
	g.r_nodes = make(map[int]string, max_num_nodes)
	g.visits = make([]int, max_num_nodes)
	g.small_caves = make([]bool, max_num_nodes)

	g.edges = make([][]int, max_num_nodes)
	for i := range g.edges {
		g.edges[i] = make([]int, max_num_nodes)
	}
}

func (g *Graph) GetNodeIdx(node string) int {
	idx, exists := g.nodes[node]
	if !exists {
		idx = len(g.nodes)
		g.nodes[node] = idx
		g.r_nodes[idx] = node
	}
	return idx
}

func (g *Graph) GetNode(idx int) string {
	node, exists := g.r_nodes[idx]
	if !exists {
		panic("Node idx: " + string(idx) + " doesn't exist!")
	}
	return node
}

func (g *Graph) AddEdge(start string, end string) {
	s_idx := g.GetNodeIdx(start)
	e_idx := g.GetNodeIdx(end)
	g.edges[s_idx][e_idx] = 1
	g.edges[e_idx][s_idx] = 1
}

func (g *Graph) isSmallCave(node int) bool {
	return g.small_caves[node]
}

func (g *Graph) isStartNode(node int) bool {
	return node == g.start
}

func (g *Graph) isEndNode(node int) bool {
	return node == g.end
}

func (g *Graph) ToDotString() string {
	dg := dot.NewGraph(dot.Directed)

	nodes := make([]dot.Node, len(g.nodes))
	for key, idx := range g.nodes {
		nodes[idx] = dg.Node(key)
		if key == "start" {
			nodes[idx].Attr("rank", "source")
		} else if key == "end" {
			nodes[idx].Attr("rank", "sink")
		}
	}
	for s_idx, edges := range g.edges {
		for e_idx, edge := range edges {
			if edge == 1 {
				dg.Edge(nodes[s_idx], nodes[e_idx])
			}
		}
	}
	return dg.String()
}

func contains(node int, path []int, allowed_visits int) bool {
	num_contains := 0
	for _, path_node := range path {
		if path_node == node {
			num_contains++
			if num_contains >= allowed_visits {
				return true
			}
		}
	}
	return false
}

func (g *Graph) GetNextPaths(prev_path []int) [][]int {
	node := prev_path[len(prev_path)-1]
	reachable := g.reachable_nodes[node]
	result := make([][]int, 0, len(reachable))
	result_idx := 0
	for _, next_node := range reachable {
		if g.isStartNode(next_node) || (g.isSmallCave(next_node) && g.path_contains(prev_path, next_node)) {
			continue
		}
		result = append(result, make([]int, 0, len(prev_path)+1))

		result[result_idx] = make([]int, 0, len(prev_path)+1)
		result[result_idx] = append(result[result_idx], prev_path...)
		result[result_idx] = append(result[result_idx], next_node)
		result_idx++
	}
	return result
}

func (g *Graph) PrecalcReachableNodes() {
	g.reachable_nodes = make(map[int][]int, len(g.nodes))
	for node := 0; node < len(g.nodes); node++ {
		next_nodes := make([]int, 0)
		for idx, edge := range g.edges[node] {
			if edge == 1 {
				next_nodes = append(next_nodes, idx)
			}
		}
		g.reachable_nodes[node] = next_nodes
	}
}

func (g *Graph) FindAllPaths() []string {

	var checked_nodes int64
	start_idx := g.nodes["start"]
	end_idx := g.nodes["end"]

	unique_paths := make([][]int, 0)

	stack := arraystack.New()

	for _, node := range g.reachable_nodes[start_idx] {
		path := make([]int, 1)
		path[0] = node
		stack.Push(path)
	}

	for !stack.Empty() {
		checked_nodes++
		it, _ := stack.Pop()
		path, _ := it.([]int)
		last_node := path[len(path)-1]
		if last_node == end_idx {
			unique_paths = append(unique_paths, path)
		} else {
			tmp := g.GetNextPaths(path)
			for _, n := range tmp {
				stack.Push(n)
			}
		}
	}

	result := make([]string, len(unique_paths))
	for i, path := range unique_paths {
		result[i] = ""
		for _, node := range path {
			result[i] += g.r_nodes[node]
		}
	}

	fmt.Printf("FindAllPaths checked nodes: %d\n", checked_nodes)

	return result
}

func (g *Graph) append_to_path(path string, node int) string {
	node_str := g.r_nodes[node]
	return path + node_str
}

func (g *Graph) str_contains(path string, node int) bool {
	num_visits := g.visits[node]
	node_str := g.r_nodes[node]
	num := strings.Count(path, node_str)
	return num > num_visits
}

func (g *Graph) path_contains(path []int, node int) bool {
	num_visits := g.visits[node]
	num := 0
	for _, entry := range path {
		if entry == node {
			num++
		}
	}
	return num > num_visits
}

var h maphash.Hash

func GetPathKey(nodes []int) uint64 {
	h.Reset()
	for _, c := range nodes {
		h.WriteByte(byte(c + 33))
	}
	return h.Sum64()
}

func check_or_add(paths *[][]int, keys *hashset.Set, path []int) bool {
	key := GetPathKey(path)
	if keys.Contains(key) {
		return false
	}
	(*paths) = append((*paths), path)
	(*keys).Add(key)
	return true
}

func (g *Graph) FindAllPathsReverse() [][]int {

	var checked_paths int64
	end_idx := g.nodes["end"]

	stack := arraystack.New()

	visited := make([]int, len(g.nodes))

	path_storage := make([][][]int, len(g.nodes))
	path_hash_storage := make([]*hashset.Set, len(g.nodes))

	for _, node := range g.nodes {
		path_storage[node] = make([][]int, 0)
		path_hash_storage[node] = hashset.New()
		visited[node] = 0
	}

	for _, node := range g.reachable_nodes[end_idx] {
		paths := &path_storage[node]
		paths_keys := path_hash_storage[node]
		path := make([]int, 1)
		path[0] = node
		check_or_add(paths, paths_keys, path)
		stack.Push(node)
	}

	for !stack.Empty() {
		it, _ := stack.Pop()
		current_node, _ := it.(int)
		if g.isStartNode(current_node) {
			continue
		}
		current_paths := &path_storage[current_node]
		for _, next_node := range g.reachable_nodes[current_node] {
			if g.isStartNode(next_node) || g.isEndNode(next_node) {
				continue
			}

			paths := &path_storage[next_node]
			path_keys := path_hash_storage[next_node]

			found_paths := 0
			for _, path := range *current_paths {
				checked_paths++
				if g.isSmallCave(next_node) && g.path_contains(path, next_node) {
					continue
				}
				new_path := make([]int, 0, len(path)+1)
				new_path = append(new_path, path...)
				new_path = append(new_path, next_node)
				if check_or_add(paths, path_keys, new_path) {
					found_paths++
					visited[next_node]++
				}
			}
			if found_paths > 0 {
				stack.Push(next_node)
			}
		}
	}
	unique_paths := make([][]int, 0)

	for _, next_node := range g.reachable_nodes[g.start] {
		paths := path_storage[next_node]
		for _, path := range paths {
			unique_path := make([]int, 0, len(path)+2)
			unique_path = append(unique_path, g.end)
			unique_path = append(unique_path, path...)
			unique_path = append(unique_path, g.start)
			unique_paths = append(unique_paths, unique_path)
		}
	}
	fmt.Printf("FindAllPathsReverse checked paths: %d\n", checked_paths)
	return unique_paths
}

func (g *Graph) CreateVisitPermutations() []*Graph {
	start_idx := g.nodes["start"]
	end_idx := g.nodes["end"]

	small_caves := make([]int, 0)

	for _, idx := range g.nodes {
		if idx == start_idx || idx == end_idx {
			continue
		} else if !g.isSmallCave(idx) {
			continue
		}
		small_caves = append(small_caves, idx)
	}

	result := make([]*Graph, 0)
	for i, small_cave_idx := range small_caves {
		result = append(result, g.Duplicate())
		result[i].visits[small_cave_idx] = 1
	}
	return result
}
