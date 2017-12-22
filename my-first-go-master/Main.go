package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/oowuyue/hello/my-first-go-master/algo"
	"github.com/oowuyue/hello/my-first-go-master/ds"
)

/*
6 6 0 1 0 2 1 3 3 4 3 5 2 5 2

6 6 a b a c b d d e d f c f 2
*/
func main() {

	id := 1
	args := os.Args

	numVertices := toInt(args[id])
	id++
	numEdges := toInt(args[id])
	id++

	g := ds.CreateGraph(numVertices)
	for i := 0; i < numEdges; i++ {
		u := (args[id])
		id++
		v := (args[id])
		id++
		g.AddEdge(u, v)
	}
	g.Print()

	operation := toInt(args[id])
	switch operation {
	case 1:
		algo.PrintBfs(g)
	case 2:
		algo.PrintDfs(g)
	default:
		fmt.Println("incorrect operation")
	}
}

func toInt(data string) int {
	convData, err := strconv.Atoi(data)
	if err != nil {
		panic(err)
	} else {
		return convData
	}
}
