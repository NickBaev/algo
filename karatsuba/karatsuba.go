package karatsuba

import "math"

func Mult (v1, v2 int) int {
	n := count(v1)
	return mult(v1, v2, n)
}

func mult (v1, v2, n int) int {
	if n == 1 {
		return v1 * v2
	}
	a, b := split(v1, n)
	c, d := split(v2, n)
	p := a + b
	q := c + d
	n = n / 2
	ac := mult(a, c, n)
	bd := mult(b, d, n)
	pq := mult(p, q ,n)
	adbc := pq - ac - bd
	return int(math.Pow10(2 * n)) * ac + int(math.Pow10(n)) * adbc + bd
}

func split (a, n int) (int, int) {
	k := int(math.Pow10(n / 2))
	c := a % k
	return (a - c) / k, c
}

func count (a int) int {
	n := 0
	b := a
	for b > 0 {
		b = b / 10
		n++
	}
	return n
}
