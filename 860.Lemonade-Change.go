func lemonadeChange(bills []int) bool {
	b5, b10, b20 := 0, 0, 0
	for _, i := range bills {
		if i == 5 {
			b5++
		} else if i == 10 {
			b10++
			if b5 > 0 {
				b5--
			} else {
				return false
			}
		} else {
			b20++
			if b5 > 0 && b10 > 0 {
				b5--
				b10--
			} else if b5 > 3 {
				b5 -= 3
			} else {
				return false
			}
		}
	}
	return true
}
