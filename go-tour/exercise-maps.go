package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	dict := make(map[string]int)
	for _, word := range strings.Split(s, " ") {
		dict[word]++
	}
	return dict
}

func main() {
	wc.Test(WordCount)
}
