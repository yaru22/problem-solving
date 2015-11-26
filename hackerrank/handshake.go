// https://www.hackerrank.com/challenges/handshake

package main

import "fmt"

func main() {
	var t, n int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		fmt.Scanf("%d", &n)
		fmt.Println(n * (n - 1) / 2)
	}
}
