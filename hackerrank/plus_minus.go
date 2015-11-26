// https://www.hackerrank.com/challenges/plus-minus

package main

import "fmt"

func main() {
	var n, nneg, nzero, npos int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		var x int
		fmt.Scanf("%d", &x)
		if x < 0 {
			nneg++
		} else if x == 0 {
			nzero++
		} else {
			npos++
		}
	}
	fmt.Println(float64(npos) / float64(n))
	fmt.Println(float64(nneg) / float64(n))
	fmt.Println(float64(nzero) / float64(n))
}
