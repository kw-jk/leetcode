package main

import "math"

func reverse(x int) int {
	answer := 0
	for x != 0 {
		digit := x % 10
		answer = answer*10 + digit
		if answer >= math.MaxInt32 {
			return 0
		}
		if answer <= math.MinInt32 {
			return 0
		}
		x /= 10
	}
	return answer
}

func main() {
	println(reverse(123456))
}
