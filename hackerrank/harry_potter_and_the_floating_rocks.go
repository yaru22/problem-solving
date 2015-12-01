// https://www.hackerrank.com/challenges/harry-potter-and-the-floating-rocks
// The same algorithm in C passes all the tests but this times out.

package main

import "fmt"

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	var T, x1, y1, x2, y2 int
	fmt.Scanf("%d\n", &T)
	for i := 0; i < T; i++ {
		fmt.Scanf("%d %d %d %d\n", &x1, &y1, &x2, &y2)
		g := gcd(y2-y1, x2-x1)
		if g > 0 {
			fmt.Println(g - 1)
		} else {
			fmt.Println(-g - 1)
		}
	}
}
