package main

import (
	"strconv"
)

func calculate(s string) int {
	var stack []int
	sum := 0
	symbol := ""
	isSubtract := false
	isDigit := false
	isNum := false
	newCal := true
	count := 0
	for _, c := range s {
		if symbol != "" {
			if c <= '9' && c >= '0' {
				num2, _ := strconv.Atoi(string(c))
				if count == (len(s) - 1) {
					if isNum {
						num1 := stack[len(stack)-1]
						num2 = num1*10 + num2
						stack = stack[:len(stack)-1]
					}
					num1 := stack[len(stack)-1]
					total := 0
					if symbol == "*" {
						total = num1 * num2
						println(total)
					} else if symbol == "/" {
						total = num1 / num2
					}
					stack = stack[:len(stack)-1]
					stack = append(stack, total)
					symbol = ""
					isNum = false
					newCal = true
					continue
				} else {
					if !isNum && newCal {
						stack = append(stack, num2)
						newCal = false
						isNum = true
					} else if isNum {
						num1 := stack[len(stack)-1]
						total := num1*10 + num2
						stack = stack[:len(stack)-1]
						stack = append(stack, total)
					}
				}
			} else {
				if newCal {
					continue
				}
				num2 := stack[len(stack)-1]
				num1 := stack[len(stack)-2]
				total := 0
				if symbol == "*" {
					total = num1 * num2
				} else if symbol == "/" {
					total = num1 / num2
				}
				stack = stack[:len(stack)-2]
				stack = append(stack, total)
				symbol = ""
				isNum = false
				newCal = true
			}
		}

		if c == '*' || c == '/' {
			symbol = string(c)
			isDigit = false
		} else if c <= '9' && c >= '0' && symbol == "" {
			num, _ := strconv.Atoi(string(c))
			if isSubtract {
				num = num * -1
				isSubtract = false
			}
			if isDigit {
				num1 := stack[len(stack)-1]
				total := num1*10 + num
				if num1 < 0 {
					total = num1*10 - num
				}
				stack = stack[:len(stack)-1]
				stack = append(stack, total)
			} else {
				stack = append(stack, num)
			}
			isDigit = true
		} else if c == '-' {
			isSubtract = true
			isDigit = false
		} else if c == '+' {
			isDigit = false
		}
		count++
	}

	for _, nums := range stack {
		sum = sum + nums
	}

	return sum
}

func main() {
	println(calculate("1*2-3/4+5*6-7*8+9/10"))
}
