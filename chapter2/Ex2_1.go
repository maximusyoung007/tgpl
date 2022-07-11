//此处的package chapter2 等同于题目中的tempconv包
package chapter2

type Celsius float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	AbsoluteZeroK Kelvin  = 0
)

func CToK(c Celsius) Kelvin {
	return Kelvin(c - 273.15)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k + 273.15)
}
