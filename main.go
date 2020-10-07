package main

import (
	"algo/stree"
	"fmt"
)

func main() {
	tree := stree.Tree{}
	tree.Add(1,2)
	tree.Add(2,2)
	tree.Add(3,2)
	tree.Add(4,2)
	tree.Add(5,2)
	tree.Add(6,2)
	tree.Delete(1)
	fmt.Println(tree.Sort())
	for _, value := range tree.Sort() {
		fmt.Println(value.Key)
	}
}