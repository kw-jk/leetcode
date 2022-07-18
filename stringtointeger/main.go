package main

import (
	"math"
)

func myAtoi(s string) int {
	var ans int64 = 0
	var sign int64 = 1
	start := false

	for _, c := range s {
		if c <= '9' && c >= '0' {
			if !start {
				start = true
			}
			ans = ans*10 + int64(c) - int64('0')
			if ans*sign > math.MaxInt32 {
				return math.MaxInt32
			} else if ans*sign < math.MinInt32 {
				return math.MinInt32
			}
		} else if !start && c == '+' {
			start = true
		} else if !start && c == '-' {
			start = true
			sign = -1
		} else if !start && c == ' ' {
			continue
		} else {
			break
		}
	}
	return int(ans * sign)
}

func main() {
	println(myAtoi("asdadasd 123"))
}
