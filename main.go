package main

import (
	"algo/quicksort"
	"math/rand"
	"time"
)

func main() {
	// Pivot element choice
	//last n=100 comp=180
	//first n=100 comp=180
	//rand n=100 comp=180
	n := 100
	comp := 0
	for i := 0; i < n; i++ {
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < n; i++ {
			a := rand.Perm(64)
			comp += quicksort.Sort(a)
		}
	}
	println(comp/(n*n))
}