package main

func romanToInt(s string) int {
	var mapping = make(map[string]int, 0)
	mapping["I"] = 1
	mapping["V"] = 5
	mapping["X"] = 10
	mapping["L"] = 50
	mapping["C"] = 100
	mapping["D"] = 500
	mapping["M"] = 1000

	total := 0
	count := 0
	for _, c := range s {
		num := mapping[string(c)]
		if c == 'I' {
			if count < len(s)-1 {
				if s[count+1] == 'V' || s[count+1] == 'X' {
					num = -1
				}
			}
		}

		if c == 'X' {
			if count < len(s)-1 {
				if s[count+1] == 'L' || s[count+1] == 'C' {
					num = -10
				}
			}
		}

		if c == 'C' {
			if count < len(s)-1 {
				if s[count+1] == 'D' || s[count+1] == 'M' {
					num = -100
				}
			}
		}

		total += num
		count++
	}

	return total
}

func main() {
	println(romanToInt("MCMXCIV"))
}
