// https://www.hackerrank.com/challenges/connecting-towns

package main

import "fmt"

func main() {
	var t, c, n int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		fmt.Scanf("%d", &c)
		var w int = 1
		for j := 0; j < c-1; j++ {
			fmt.Scanf("%d", &n)
			w *= n
			w %= 1234567
		}
		fmt.Println(w)
	}
}
