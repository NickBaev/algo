package secondb

func Secondb(arr []int) int  {
	min, _ := round(arr, len(arr))
	return min
}

func round(arr []int, n int) (int, int){
	if n == 2 {
		if arr[0] > arr[1] {
			return arr[1], arr[0]
		}
		return arr[0], arr[1]
	}
	n = n/2
	a, b := split(arr, n)
	amin, amax := round(a, n)
	bmin, bmax := round(b, n)
	if amax > bmax {
		if amin > bmax {
			return amin, amax
		}
		return bmax, amax
	}
	if bmin > amax {
		return bmin, bmax
	}
	return amax, bmax
}

func split (arr []int, n int) ([]int, []int) {
	return arr[:n], arr[n:]
}