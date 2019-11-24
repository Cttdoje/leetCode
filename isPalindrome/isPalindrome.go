package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.Replace(s, " ", "", -1)
	if len(s) == 0 {
		return true
	}
	s = strings.ToLower(s)
	x := 0
	y := len(s) - 1
	for x < y {
		if !(s[x] >= 97 && s[x] <= 122) && !(s[x] >= 48 && s[x] <= 57) {
			x++
			continue
		}
		if !(s[y] >= 97 && s[y] <= 122) && !(s[y] >= 48 && s[y] <= 57) {
			y--
			continue
		}
		if s[x] != s[y] {
			return false
		} else {
			x++
			y--
		}
	}
	return true
}

func main() {
	str := "+-19azAZ"
	ss := []rune(str)
	for i, r := range ss {
		fmt.Println(i, r)
		//r -= '0'
		//fmt.Println(r)

	}
	//utf8.DecodeLastRuneInString(str)
	//fmt.Println(isPalindrome("0p"))
	r, _ := strconv.Atoi("123a")
	fmt.Println(r)
}
