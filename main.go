package main

import (
	"algo/quicksort"
	"math/rand"
	"time"
)

// Sort 64 elements array. Pivot element choice
//last: n=100 comp=180
//first: n=100 comp=180
//rand: n=100 comp=180

// RSelect 64 element array comparisons
//index=5: n=100 comp=70
//index=60: n=100 comp=66
//index=32: n=100 comp=94
func main() {
	rselect()
}

func rselect() {
	n := 100
	comp := 0
	for i := 0; i < n; i++ {
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < n; i++ {
			a := rand.Perm(64)
			_, c := quicksort.RSelect(a, 32)
			comp += c
		}
	}
	println(comp/(n*n))
}

func sort()  {
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