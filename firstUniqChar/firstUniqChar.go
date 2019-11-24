package main

import "fmt"

func firstUniqChar(s string) int {
	res := -1
	for i, str := range s {
		b := false
		for j, strT := range s {
			if i == j {
				continue
			}
			if str == strT {
				b = true
				break
			}
		}
		if b == false {
			res = i
			return res
		}
	}
	return res
}

func main() {
	fmt.Println(firstUniqChar("leetcode"))
}
