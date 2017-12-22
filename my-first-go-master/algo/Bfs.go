package algo

import (
	"fmt"

	"github.com/oowuyue/hello/my-first-go-master/ds"
)

func PrintBfs(g ds.Graph) {
	fmt.Println("PrintBfs...")

	v := g.Vertices
	edges := g.Edges

	queue := ds.CreateQueue()
	visited := make([]bool, v)

	source := 0
	visited[source] = true
	queue.Enqueue(source)

	for queue.IsEmpty() == false {
		deqVal := queue.Dequeue().Value
		fmt.Println(g.VertSlice[deqVal])
		for i := range edges[deqVal] {
			vertex := edges[deqVal][i]
			if visited[vertex] == false {
				visited[vertex] = true
				queue.Enqueue(vertex)
			}
		}
	}
}
