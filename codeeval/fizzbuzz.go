package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func fizzBuzz(x, y, n int) {
	for i := 1; i <= n; i++ {
		if i%x == 0 && i%y == 0 {
			fmt.Print("FB")
		} else if i%x == 0 {
			fmt.Print("F")
		} else if i%y == 0 {
			fmt.Print("B")
		} else {
			fmt.Print(i)
		}
		fmt.Print(" ")
	}
	fmt.Println()
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		x, _ := strconv.ParseInt(line[0], 10, 0)
		y, _ := strconv.ParseInt(line[1], 10, 0)
		n, _ := strconv.ParseInt(line[2], 10, 0)
		fizzBuzz(int(x), int(y), int(n))
	}
}
