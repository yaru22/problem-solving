// https://www.hackerrank.com/contests/projecteuler/challenges/euler001

package main

import "fmt"

func multiple_sum(m, n int) int {
	return (n - 1) / m * ((n-1)/m + 1) / 2 * m
}

func main() {
	var n, m int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &m)
		fmt.Println(multiple_sum(3, m) + multiple_sum(5, m) - multiple_sum(15, m))
	}
}
