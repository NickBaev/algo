package main

import (
	"algo/dijkstra"
	"fmt"
)

func main() {
	v := []int{1,2,3,4,5,6,7,8}
	e := []dijkstra.SimpleEdge{
		{1, 2, 1},
		{1, 8, 2},
		{2, 1, 1},
		{2, 3, 1},
		{3, 2, 1},
		{3, 4, 1},
		{4, 3, 1},
		{4, 5, 1},
		{5, 4, 1},
		{5, 6, 1},
		{6, 5, 1},
		{6, 7, 1},
		{7, 6, 1},
		{7, 8, 1},
		{8, 7, 1},
		{8, 1, 2},
	}

	graph := dijkstra.FindPath(v, e, 1)
	for _, value := range graph {
		fmt.Println(value.Value, " -> ",  value.Length)
	}
}