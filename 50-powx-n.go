func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / myPow(x, -n)
	} else if n == 0 {
		return 1
	} else if n == 1 {
		return x
	}
	if n%2 == 1 {
		return myPow(x*x, n/2) * x
	}
	return myPow(x*x, n/2)
}

func myPow(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}
	pow := 1.0
	for n != 0 {
		if n&1 != 0 {
			pow = pow * x
		}
		x = x * x
		n >>= 1
	}
	return pow
}
