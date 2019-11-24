package main

import "fmt"

func isAnagram(s string, t string) bool {

	map1 := make(map[rune]int)
	for star, key := range s {
		if _, exist := map1[key]; exist != false {
			continue
		}
		count := 0
		for i, cur := range s {
			if i < star {
				continue
			}
			if cur == key {
				count++
			}
		}
		map1[key] = count
	}

	for _, tKey := range t {
		count, exist := map1[tKey]
		if exist == false {
			return false
		}
		count--
		map1[tKey] = count
	}

	for _, count := range map1 {
		if count != 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}
