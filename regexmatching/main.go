package main

import "fmt"

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	j := 2
	for j <= len(p) {
		dp[0][j] = p[j-1] == '*' && dp[0][j-2]
		println(dp[0][1])
		j++
	}

	j = 1
	for j <= len(p) {
		i := 1
		for i <= len(s) {
			if p[j-1] == s[i-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2] || ((s[i-1] == p[j-2]) || p[j-2] == '.') && dp[i-1][j]
			}
			i++
		}
		j++
	}
	return dp[len(s)][len(p)]
}

func main() {
	fmt.Printf("%v", isMatch("aab", "a.*"))
}
