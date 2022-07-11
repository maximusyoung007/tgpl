package chapter2

func Ex2_5(x uint64) int {
    ans := 0
    for x != 0 {
        x = x & (x - 1)
        ans++
    }
    return ans
}
