package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var res int64
	b := true
	if x < 0 {
		b = false
		x = 0 - x
	}
	for x > 0 {
		temp := x % 10
		x = x / 10
		res = res*10 + int64(temp)
	}
	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}
	if b == false {
		return int(0 - res)
	}
	return int(res)
}

func main() {
	fmt.Println(reverse(-123))
}
