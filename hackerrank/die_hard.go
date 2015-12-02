package main

import "fmt"

func canFill(a, b, c int) bool {
	visited := make(map[int]bool)
	if a == c || b == c {
		return true
	}
	visited[a] = true
	visited[b] = true
	n := a
	for {
		if n < b {
			n = a - (b - n)
		} else {
			n = n - b
		}

		if v, _ := visited[n]; v {
			return false
		}
		if c == n {
			return true
		}
		visited[n] = true
	}
}

func main() {
	var t, a, b, c int
	fmt.Scanf("%d\n", &t)
	for i := 0; i < t; i++ {
		fmt.Scanf("%d %d %d\n", &a, &b, &c)
		if a < b {
			b, a = a, b
		}
		if canFill(a, b, c) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
