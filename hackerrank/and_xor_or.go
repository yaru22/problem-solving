// https://www.hackerrank.com/challenges/and-xor-or

package main

import "fmt"

func calc(m1, m2 int) int {
	return ((m1 & m2) ^ (m1 | m2)) & (m1 ^ m2)
}

func main() {
	var n, m1, m2, val int
	fmt.Scanf("%d\n", &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	max := 0
	for i := 0; i < n-1; i++ {
		if arr[i] < arr[i+1] {
			m1, m2 = arr[i], arr[i+1]
		} else {
			m1, m2 = arr[i+1], arr[i]
		}
		val = calc(m1, m2)
		if max < val {
			max = val
		}
		// NOTE: upper bound of calc(m1, m2) is m1 + m2
		if m1+m2 <= max {
			continue
		}
		for j := i + 2; j < n; j++ {
			if arr[j] >= m2 {
				continue
			} else if arr[j] < m1 {
				break
			} else {
				m1, m2 = m1, arr[j]
			}
			val = calc(m1, m2)
			if max < val {
				max = val
			}
			// NOTE: upper bound of calc(m1, m2) is m1 + m2
			if m1+m2 <= max {
				break
			}
		}
	}

	fmt.Println(max)
}
