package main

import (
	"fmt"
	"strconv"
	"strings"
)

func decodeString(s string) string {
	count := 0
	var decoded []string
	for idx, c := range s {
		if i, err := strconv.Atoi(string(c)); err == nil {
			count = count*10 + i
		} else if string(s[idx]) == "[" {
			decoded = append(decoded, strconv.Itoa(count))
			count = 0
		} else if string(s[idx]) == "]" {
			mystr := ""
			for {
				if _, err := strconv.Atoi(decoded[len(decoded)-1]); err != nil {
					mystr = decoded[len(decoded)-1] + mystr
					decoded = decoded[:len(decoded)-1]
				} else {
					break
				}
			}
			num, _ := strconv.Atoi(decoded[len(decoded)-1])
			decoded = decoded[:len(decoded)-1]
			newstr := strings.Repeat(mystr, num)
			decoded = append(decoded, newstr)
		} else {
			decoded = append(decoded, string(c))
		}
	}

	return strings.Join(decoded, "")
}

func main() {
	test := decodeString("100[leetcode]")
	fmt.Printf("decoded: %s", test)
}
