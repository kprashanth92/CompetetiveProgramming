package problems

func FindKinArray(arr []int, k int) bool {
	a := make(map[int]int)
	for _, x := range arr {
		rem := k - x
		if a[rem] != 0 {
			return true
		} else {
			a[x] = rem
		}
	}
	return false
}
