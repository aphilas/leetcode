func nextGreatestLetter(letters []byte, target byte) byte {
    var lo, hi int

	for lo, hi = 0, len(letters); lo < hi; {
		mid := (hi + lo) / 2

		if letters[mid] > target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return letters[lo%len(letters)]
}