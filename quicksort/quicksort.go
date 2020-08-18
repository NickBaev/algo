package quicksort

import (
	"math/rand"
)

func Sort(arr[] int) int {
	count := make(chan int, 1000)
	sort(arr, 0, len(arr) - 1, count)
	comp := 0
	close(count)
	for val := range count {
		comp += val
	}
	return comp
}

func sort (arr []int, left int, right int, count chan int) {
	if left >= right {
		return
	}

	i := pivotRand(left, right) // random
	//i := left // first
	//i := right // last

	arr[left], arr[i] = arr[i], arr[left]
	j := partition(arr, left, right, count)
	sort(arr, left, j - 1, count)
	sort(arr, j + 1, right, count)
}

func RSelect (arr[] int, index int) (int, int) {
	count := make(chan int, 1000)
	res := rselect(arr, 0, len(arr) - 1, index, count)
	comp := 0
	close(count)
	for val := range count {
		comp += val
	}
	return res, comp
}

func rselect (arr []int, left int, right int, index int, count chan int) int {
	if left >= right {
		return arr[right]
	}
	i := pivotRand(left, right) // random
	arr[left], arr[i] = arr[i], arr[left]
	j := partition(arr, left, right, count)
	if j == index {
		return arr[j]
	}
	if j > index {
		return rselect(arr, left, j - 1, index, count)
	} else {
		return rselect(arr, j + 1, right, index, count)
	}
}

func partition(arr []int, left int, right int, count chan int) int {
	p := arr[left]
	i := left + 1
	for j := left + 1; j <= right; j++ {
		if arr[j] < p {
			arr[j], arr[i] = arr[i], arr[j]
			i++
			count <- 1
		}
	}
	i--
	arr[left], arr[i] = arr[i], arr[left]
	return i
}

func pivotRand (left int, right int) int {
	n := right - left
	return rand.Intn(n) + left
}