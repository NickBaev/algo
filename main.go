package main

import (
	"algo/karatsuba"
	"fmt"
)

func main() {
	a := 1234
	b := 2456
	fmt.Print(a * b)
	fmt.Print("\n")
	fmt.Print(karatsuba.Mult(a, b))
}