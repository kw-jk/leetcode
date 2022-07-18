package main

import "strconv"

func Calculate(s string) int {
	if len(s) == 0 {
		return 0
	}

	count := 0
	sum := 0
	operator := "+"
	currentNum := 0
	var stack []int
	for count < len(s) {
		currentChar := string(s[count])
		isNum := false
		if s[count] <= '9' && s[count] >= '0' {
			isNum = true
			num, _ := strconv.Atoi(currentChar)
			currentNum = currentNum*10 + num
		}
		if !isNum && currentChar != " " || count == (len(s)-1) {
			if operator == "-" {
				stack = append(stack, -currentNum)
			}
			if operator == "+" {
				stack = append(stack, currentNum)
			}
			if operator == "*" {
				prevNum := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, prevNum*currentNum)
			}
			if operator == "/" {
				prevNum := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, prevNum/currentNum)
			}
			operator = currentChar
			currentNum = 0
		}
		count++
	}

	for _, nums := range stack {
		sum = sum + nums
	}

	return sum
}

func main() {
	println(Calculate("42"))
}
