package chapter2

func Ex2_3(x uint64) int {
	ans := 0
	for i := 0; i < 8; i++ {
		ans += int(pc[byte(x)])
		x >>= 8
	}
	return ans
}
