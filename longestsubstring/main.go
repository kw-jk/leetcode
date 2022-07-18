package main

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func lengthOfLongestSubstring(s string) int {
	testmap := make(map[string]int)
	sum := 0
	start := 0
	for idx, c := range s {
		if val, ok := testmap[string(c)]; ok {
			if val >= start {
				start = val + 1
			}
		}
		sum = max(sum, idx-start+1)
		testmap[string(c)] = idx
	}
	return sum
}

func main() {
	count := lengthOfLongestSubstring("dxdxfzxf")
	println(count)
}
