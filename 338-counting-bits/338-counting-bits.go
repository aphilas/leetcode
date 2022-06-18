func divmod(n, d int) (q, r int) {
	q = n / d
	r = n % d
	return
}

func countBits(n int) []int {
    ans := make([]int, n+1)
	curr := 0

	for i := 0; i <= n; i++ {
		q, r := divmod(i, 2)

		for q > 1 && r == 0 {
			curr -= 1
			q, r = divmod(q, 2)
		}

		if r > 0 {
			curr += 1
		}

		ans[i] = curr
	}

	return ans
}