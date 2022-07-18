package main

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func min(x int, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func maxArea(height []int) int {
	maxArea := 0
	left := 0
	right := len(height) - 1

	for left < right {
		width := right - left
		maxArea = max(maxArea, min(height[left], height[right])*width)
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func main() {
	println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
