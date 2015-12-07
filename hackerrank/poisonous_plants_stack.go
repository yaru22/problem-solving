// https://www.hackerrank.com/challenges/poisonous-plants

package main

import "fmt"

func main() {
	var n, sol int
	fmt.Scanf("%d\n", &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	tdeath := make([]int, n)
	stack := make([]int, 1)
	for i := 1; i < n; i++ {
		l := len(stack)
		if arr[i] > arr[stack[l-1]] {
			tdeath[i] = 1
		}
		for l > 0 && arr[i] <= arr[stack[l-1]] {
			if tdeath[i] < tdeath[stack[l-1]]+1 {
				tdeath[i] = tdeath[stack[l-1]] + 1
			}
			stack = stack[:l-1]
			l--
		}

		if len(stack) == 0 {
			tdeath[i] = 0
		}
		stack = append(stack, i)

		if sol < tdeath[i] {
			sol = tdeath[i]
		}
	}
	fmt.Println(sol)
}
