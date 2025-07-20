package factorial

const (
	maxSliceSize = 5000
)

func Recursive(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * Recursive(n-1)
	}
}

func Dynamic(n int) int {
	memo := make([]int, n+1)
	memo[0] = 1
	for i := 1; i <= n; i++ {
		memo[i] = i * memo[i-1]
	}
	return memo[n]
}

func Calculate() [maxSliceSize]int {
	s := [maxSliceSize]int{}
	for i := 0; i < maxSliceSize; i++ {
		s[i] = Dynamic(i)
	}
	return s
}
