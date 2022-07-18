package main

import "fmt"

func isMatched(s string, p string) bool {
	return helper(s, len(s)-1, p, len(p)-1)
}

func helper(s string, i1 int, p string, i2 int) bool {
	if i1 == -1 && i2 == -1 {
		return true
	} else if i2 == -1 {
		return false
	}

	if i1 >= 0 && (p[i2] == '.' || p[i2] == s[i1]) {
		// recursive call to find remaining matching patterns
		return helper(s, i1-1, p, i2-1)
	} else if p[i2] == '*' {
		ch := p[i2-1]
		isDot := ch == '.'
		j := i1
		for j >= 0 && (isDot || s[j] == ch) {
			if helper(s, j-1, p, i2-2) {
				return true
			}
			j--
		}
		return helper(s, i1, p, i2-2)
	} else {
		return false
	}
}

func main() {
	fmt.Printf("%v", isMatched("aaa", "aaaa"))
}
