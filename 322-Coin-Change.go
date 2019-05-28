import (
	"fmt"
	"math"
)

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	m := make([]int, amount+1)
	for _, i := range coins {
		if i <= amount {
			m[i] = 1
		}
	}
	for i := 0; i < amount+1; i++ {
		t, t1 := math.MaxInt64, 0
		for _, j := range coins {
			if i-j > 0 {
				t1 = m[i-j]
			}
			if t > t1 && t1 != 0 {
				t = t1
			}
		}
		if m[i] == 0 && t != math.MaxInt64 {
			m[i] = t + 1
		} else if m[i] == 0 && t == math.MaxInt64 {
			m[i] = 0
		}
	}
	if m[amount] == 0 {
		return -1
	}
	return m[amount]
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	m := make([]int, amount+1)
	for i := range m {
		m[i] = amount + 1
	}
	m[0] = 0
	for i := 0; i < amount+1; i++ {
		for _, j := range coins {
			if j <= i {
				t := m[i-j] + 1
				if t < m[i] {
					m[i] = t
				}
			}
		}
	}
	fmt.Println(m)
	if m[amount] >= amount+1 {
		return -1
	}
	return m[amount]
}
