package ds

import "fmt"

// Node1 ..
type Node1 struct {
	Value interface{}
}

// Graph ..
type Graph struct {
	Vertices  int
	VertSlice []Node1
	Edges     [][]int
}

// CreateGraph ..
func CreateGraph(v int) Graph {
	return Graph{Vertices: v, Edges: make([][]int, v)}
}

// AddEdge ..
func (g *Graph) AddEdge(u interface{}, v interface{}) {
	uindex := g.AddVert(u)
	vindex := g.AddVert(v)
	g.Edges[uindex] = append(g.Edges[uindex], vindex)
	g.Edges[vindex] = append(g.Edges[vindex], uindex)
}

// AddVert ...
func (g *Graph) AddVert(v interface{}) int {
	for index, node := range g.VertSlice {
		if node.Value == v {
			return index
		}
	}
	g.VertSlice = append(g.VertSlice, Node1{Value: v})
	return len(g.VertSlice) - 1
}

// Print ...
func (g *Graph) Print() {
	fmt.Println("Graph")
	for k0 := range g.Edges {
		fmt.Print(g.VertSlice[k0], "--->")
		for _, v := range g.Edges[k0] {
			fmt.Print(g.VertSlice[v])
		}
		fmt.Println()
	}
}
