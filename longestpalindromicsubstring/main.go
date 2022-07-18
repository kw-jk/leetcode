package main

func longestPalindrome(s string) string {
	m := ""
	for idx, _ := range s {
		ans := helper(idx, idx, s, m)
		ans2 := helper(idx, idx+1, s, m)
		if len(ans2) > len(ans) {
			ans = ans2
		}
		if len(ans) > len(m) {
			m = ans
		}
	}
	return m
}

func helper(l, r int, s string, ans string) string {
	ansLen := len(ans)
	for l >= 0 && r < len(s) && string(s[l]) == string(s[r]) {
		if r-l+1 > ansLen {
			ans = s[l : r+1]
			ansLen = r - l + 1
		}
		l--
		r++
	}
	return ans
}

func main() {
	substring := longestPalindrome("aacabdkacaa")
	println(substring)
}
