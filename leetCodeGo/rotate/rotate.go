package main

import (
	"fmt"
	"os"
)

func rotate(matrix [][]int) {
	row := len(matrix)
	var endRow int
	if len(matrix)%2 == 0 {
		endRow = len(matrix) / 2
	} else {
		endRow = len(matrix)/2 + 1
	}
	for i := 0; i < endRow; i++ {
		for j := i; j < len(matrix)-i-1; j++ {
			nextI := j
			nextJ := row - i - 1
			current := matrix[i][j]
			var temp int
			for k := 0; k < 4; k++ {
				temp = matrix[nextI][nextJ]
				matrix[nextI][nextJ] = current
				current = temp
				nextI, nextJ = nextJ, row-nextI-1
			}
		}
	}
}

func printRotate(matrix [][]int) {
	for i := range matrix {
		for j := range matrix[i] {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}

func main() {
	file, err := os.Open("leetCode/image.in")
	if err != nil {
		panic(err)
	}
	var row int
	_, _ = fmt.Fscanf(file, "%d", &row)
	var matrix = make([][]int, row)
	for i := range matrix {
		matrix[i] = make([]int, row)
		for j := range matrix[i] {
			_, _ = fmt.Fscanf(file, "%d", &matrix[i][j])
		}
	}

	fmt.Println("旋转前:")
	printRotate(matrix)
	rotate(matrix)
	fmt.Println("旋转后:")
	printRotate(matrix)
}
