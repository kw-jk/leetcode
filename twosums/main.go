package main

import "fmt"

func twoSum(nums []int, target int) []int {
	set := make(map[int]int)
	var ans []int
	for idx, v := range nums {
		if _, ok := set[target-v]; ok {
			ans = append(ans, idx)
			ans = append(ans, set[target-v])
		}
		set[v] = idx
	}

	return ans
}

func main() {
	nums := []int{2, 7, 11, 15}
	test := twoSum(nums, 9)
	fmt.Printf("%v", test)
}
