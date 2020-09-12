package main

import (
	"algo/kosaraju"
	"fmt"
)

func main() {
	v := []int{1,2,3,4,5,6,7,8,9,10,11,12}
	e := []kosaraju.SimpleEdge{
		{1,2},
		{2, 3},
		{2, 4},
		{2, 5},
		{3, 6},
		{4, 5},
		{4, 7},
		{5, 2},
		{5, 6},
		{5, 7},
		{6, 3},
		{6, 8},
		{7, 8},
		{7, 10},
		{8, 7},
		{9, 7},
		{10, 9},
		{10, 11},
		{11, 12},
		{12, 10},
	}
	graph := kosaraju.FindSCC(v, e)
	fmt.Println(kosaraju.GroupBySCC(graph))

	v = []int{1,2,3,4,5,6,7,8}
	e = []kosaraju.SimpleEdge{
		{1, 2},
		{2, 3},
		{3, 1},
		{3, 4},
		{5, 4},
		{6, 4},
		{8, 6},
		{6, 7},
		{7, 8},
		{4, 3},
		{4, 6},
	}
	graph = kosaraju.FindSCC(v, e)
	fmt.Println(kosaraju.GroupBySCC(graph))
}