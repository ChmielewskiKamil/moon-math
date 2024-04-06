package main

// Exercise 7 from Moon Math manual
// Integer Long Division
func longDivision(dividend, divisor int) (quotient int, remainder uint) {
	if divisor == 0 {
		panic("Dividing by 0!")
	}

	if dividend == 0 {
		return 0, 0
	}

	neg := false

	if dividend < 0 {
		dividend = -dividend
		neg = !neg
	}
	if divisor < 0 {
		divisor = -divisor
		neg = !neg
	}

	for dividend >= divisor {
		dividend -= divisor
		quotient++
	}

	remainder = uint(dividend)

	if neg {
		quotient = -quotient - 1
		remainder = uint(divisor) - remainder
	}

	return quotient, remainder
}

func main() {
	println("Hello, World!")
}
