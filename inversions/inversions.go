package inversions

func Count(arr []int) ([]int, int){
	n := len(arr)
	if n == 0 || n == 1 {
		return arr, 0
	}
	a, b := split(arr)
	left, lcount := Count(a)
	right, rcount := Count (b)
	arr, crossCount := merge(left, right)
	return arr, lcount + rcount + crossCount
}

func merge(left []int, right []int) ([]int, int)  {
	rlen := len(left) + len(right)
	result := make([]int, rlen)
	i := 0
	j := 0
	inversions := 0
	for k := 0; k < rlen; k++ {
		if i >= len(left) {
			for j < len(right) {
				result[k] = right[j]
				j++
				k++
			}
			break
		}
		if j >= len(right) {
			for i < len(left) {
				result[k] = left[i]
				i++
				k++
			}
			break
		}
		if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			inversions += len(left) - i
			j++
		}
	}
	return result, inversions
}

func split (arr []int) ([]int, []int) {
	n := len(arr)/2
	return arr[:n], arr[n:]
}