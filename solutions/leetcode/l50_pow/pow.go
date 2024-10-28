package l50pow

// https://leetcode.com/problems/powx-n/
func myPow(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}
	pow := 1.0
	for n != 0 {
		if (n & 1) != 0 {
			pow *= x
		}

		x *= x
		n >>= 1

	}
	return pow
}
