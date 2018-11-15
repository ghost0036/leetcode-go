package leetcode_50

func myPow(x float64, n int) float64 {
	ret := 1.0
	if n < 0{
		x = 1/x
		n = n * -1
	}
	for n != 0 {
		if n%2 != 0 {
			ret = ret * x
		}
		n /= 2
		x = x * x
	}
	return ret
}
