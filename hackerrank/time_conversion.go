// https://www.hackerrank.com/challenges/time-conversion

package main

import "fmt"

func main() {
	var hh, mm, ss int
	var ampm string
	fmt.Scanf("%d:%d:%d%s", &hh, &mm, &ss, &ampm)
	if ampm == "PM" && hh < 12 {
		fmt.Printf("%02d:%02d:%02d\n", hh+12, mm, ss)
	} else if ampm == "AM" && hh == 12 {
		fmt.Printf("%02d:%02d:%02d\n", hh-12, mm, ss)
	} else {
		fmt.Printf("%02d:%02d:%02d\n", hh, mm, ss)
	}

}
