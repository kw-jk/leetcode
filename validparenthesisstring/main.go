package main

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func checkValidString(s string) bool {
	if len(s) == 0 {
		return false
	}

	lo, hi := 0, 0

	for _, c := range s {
		if c == '(' {
			lo++
		} else {
			lo--
		}

		if c != ')' {
			hi++
		} else {
			hi--
		}

		if hi < 0 {
			break
		}
		lo = max(lo, 0)
	}

	return lo == 0
}

func main() {
	println(checkValidString("(*))"))
}
