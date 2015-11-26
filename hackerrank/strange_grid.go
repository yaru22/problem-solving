// https://www.hackerrank.com/challenges/strange-grid

package main

import (
	"fmt"
)

func main() {
	var r, c int
	fmt.Scanf("%d %d", &r, &c)
	fmt.Println((r-1)/2*10 + (r+1)%2 + (c-1)*2)
}
