// https://www.hackerrank.com/challenges/diagonal-difference

package main

import (
	"fmt"
	"math"
)

func main() {
	var n, diag1, diag2 int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			var x int
			fmt.Scanf("%d", &x)
			if i == j {
				diag1 += x
			}
			if i+j == n-1 {
				diag2 += x
			}
		}
	}
	fmt.Println(int(math.Abs(float64(diag1 - diag2))))
}
