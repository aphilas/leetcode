func peakIndexInMountainArray(a []int) int {
    var hi, lo, mid int

	for lo, hi = 0, len(a)-1; lo < hi; {
		mid = (hi + lo) / 2

		if a[mid] < a[mid+1] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return lo
}