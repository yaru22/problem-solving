// https://www.hackerrank.com/challenges/and-xor-or

package main

import (
	"bufio"
	"fmt"
	"os"
)

func calc(m1, m2 int) int {
	return ((m1 & m2) ^ (m1 | m2)) & (m1 ^ m2)
}

func main() {
	var n, max, numMin, val int
	input := bufio.NewReader(os.Stdin)

	fmt.Scanf("%d\n", &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(input, "%d", &arr[i])
		if i > 0 {
			numMin = arr[i-1]
		}
		for j := i - 1; j >= 0; j-- {
			val = calc(arr[i], numMin)
			if max < val {
				max = val
			}
			if i-j > 5 {
				break
			}
			if numMin <= arr[i] {
				break
			}
			if arr[j] < numMin {
				numMin = arr[j]
			}
			if arr[i]+numMin <= max {
				break
			}
		}
	}
	fmt.Println(max)
}
