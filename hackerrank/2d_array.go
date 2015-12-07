// https://www.hackerrank.com/challenges/2d-array

package main

import "fmt"
import "math"

func sumHourGlass(arr *[][]int, startI, startJ int) int {
	sum := 0
	for i := startI; i < startI+3; i++ {
		for j := startJ; j < startJ+3; j++ {
			if !(i == startI+1 && j == startJ) && !(i == startI+1 && j == startJ+2) {
				sum += (*arr)[i][j]
			}
		}
	}
	return sum
}

func main() {
	const n = 6
	// Read input data.
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &arr[i][j])
		}
	}

	// Check sum of hour glasses.
	max := math.MinInt32
	for i := 0; i <= n/2; i++ {
		for j := 0; j <= n/2; j++ {
			sum := sumHourGlass(&arr, i, j)
			if max < sum {
				max = sum
			}
		}
	}
	fmt.Println(max)
}
