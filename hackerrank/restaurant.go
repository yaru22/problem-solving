// https://www.hackerrank.com/challenges/restaurant

package main

import "fmt"

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	var n, l, b int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d", &l, &b)
		g := gcd(l, b)
		fmt.Println(l / g * b / g)
	}
}
