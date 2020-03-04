package main

import "fmt"

func reverseString(s []byte) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		temp := s[length-i-1]
		s[length-i-1] = s[i]
		s[i] = temp
	}
}

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Println(s)
	reverseString(s)
	fmt.Println(s)
}
