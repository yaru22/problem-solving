package main

import (
	"fmt"
	"strconv"
)

func main() {
	var t, n int
	fmt.Scanf("%d\n", &t)
	for i := 0; i < t; i++ {
		fmt.Scanf("%d\n", &n)
		var queue = []string{"9"}
		for {
			var s string
			s, queue = queue[0], queue[1:]
			sNum, _ := strconv.Atoi(s)
			if sNum%n == 0 {
				fmt.Println(sNum)
				break
			}
			queue = append(queue, s+"0", s+"9")
		}
	}
}
