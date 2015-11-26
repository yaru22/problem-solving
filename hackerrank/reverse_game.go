// https://www.hackerrank.com/challenges/reverse-game

package main

import "fmt"

// Original solution:
// func main() {
// 	var t, n, k int
// 	fmt.Scanf("%d", &t)
// testcase:
// 	for i := 0; i < t; i++ {
// 		fmt.Scanf("%d %d", &n, &k)
// 		var j, l int
// 		for j, l = 0, n; j < l; j++ {
// 			k = (n - 1) - k
// 			if k == 0 {
// 				fmt.Println(k + j)
// 				continue testcase
// 			}
// 			k--
// 			n--
// 		}
// 	}
// }

// Revised solution (much more readable)
// NOTE:
// [ 0 1 2 3 4 ]  ball 2's final position in [ 0 1 2 3 4 ]
// 4 [ 3 2 1 0 ]  = ball 2's final position in [ 3 2 1 0 ] + 1
// 4 0 [ 1 2 3 ]  = ball 2's final position in [ 1 2 3 ] + 2
// 4 0 3 [ 2 1 ]  = ball 2's final position in [ 2 1 ] + 3
// 4 0 3 1 [ 2 ]  = ball 2's final position in [ 2 ] + 4
func game(n, k int) int {
	if k < 0 {
		return k
	}
	// When you reverse the list, the position of the ball becomes
	// (n - 1) - k
	// the last -1 is for the shortened list
	// i.e. [ 0 1 2 3 4 ] -> [ 4 3 2 1 0 ] -> 4 [ 3 2 1 0 ]
	// (2 is now located at n - 1 - k - 1 in [ 3 2 1 0 ])
	return game(n-1, n-1-k-1) + 1
}

func main() {
	var t, n, k int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		fmt.Scanf("%d %d", &n, &k)
		fmt.Println(game(n, k))
	}
}
