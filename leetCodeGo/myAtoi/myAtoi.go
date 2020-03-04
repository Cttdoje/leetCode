package main

import (
	"fmt"
	"math"
)

func myAtoi(str string) int {
	var res int64
	isFu := false
	if len(str) == 0 {
		return 0
	}
	strRune := []rune(str)
	for strRune[0] == 32 {
		if len(strRune) > 1 {
			strRune = strRune[1:]
		} else {
			return 0
		}
	}

	for i, s := range strRune {
		if (s == 43 || s == 45) && i != 0 && res == 0 {
			return 0
		}
		if s == 45 && res == 0 {
			isFu = true
			continue
		}

		if s == 43 && res == 0 {
			continue
		}

		if s < 48 || s > 57 {
			break
		}
		s -= '0'
		if isFu && (res > math.MaxInt32/10 || (res == math.MaxInt32/10 && s > 7)) {
			return math.MinInt32
		}
		if !isFu && (res > math.MaxInt32/10 || (res == math.MaxInt32/10 && s > 7)) {
			return math.MaxInt32
		}

		res = res*int64(10) + int64(s)
	}
	if isFu {
		res = -res
	}
	return int(res)
}

func main() {
	fmt.Println(myAtoi("-13+8"))
}
