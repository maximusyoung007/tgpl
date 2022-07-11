package chapter2

func Ex2_4(x uint64) int {
	res := 0
	for i := 0; i < 64; i++ {
		if x&1 == 1 {
			res++
		}
		x >>= 1
	}
	return res
}
