// https://www.hackerrank.com/challenges/staircase

package main

import "fmt"

func main() {
	var n int = 5
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		for j := 1; j <= n; j++ {
			if j < n-i {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}
