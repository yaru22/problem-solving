// https://www.hackerrank.com/challenges/poisonous-plants

package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	var i, j, l, cnt int
	arrLen := n
	for cnt = 0; ; cnt++ {
		j = 0
		for i, l = 1, arrLen; i < l; i++ {
			if arr[i-1] >= arr[i] {
				j++
				arr[j] = arr[i]
			}
		}
		if i == j+1 {
			fmt.Println(cnt)
			break
		}
		arrLen = j + 1
	}

}
