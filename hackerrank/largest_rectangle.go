// https://www.hackerrank.com/challenges/largest-rectangle

package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	maxArea := 0
	for i := 0; i < n; i++ {
		min := math.MaxInt32
		for j := i; j < n; j++ {
			if arr[j] < min {
				min = arr[j]
			}
			// Early exit if we know the max area starting from
			// i-th building is going to be less than the known
			// maxArea.
			if min*(n-i+1) < maxArea {
				break
			}
			area := min * (j - i + 1)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	fmt.Println(maxArea)
}
