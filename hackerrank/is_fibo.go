package main

import "fmt"

var fiboMap = make(map[int]bool)

func fibo() {
	a, b := 0, 1
	for {
		fiboMap[a] = true
		fiboMap[b] = true
		if b > 10000000000 {
			break
		}
		a, b = b, a+b
	}
}

func isFibo(n int) bool {
	return fiboMap[n]
}

func main() {
	fibo()

	var T, n int
	fmt.Scanf("%d", &T)
	for i := 0; i < T; i++ {
		fmt.Scanf("%d", &n)
		if isFibo(n) {
			fmt.Println("IsFibo")
		} else {
			fmt.Println("IsNotFibo")
		}
	}
}
