package main

import (
	"fmt"
	"os"
	"strconv"
)

// todo:
// 1. compare performance
// 2. cover with tests
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: `go run . <int>`")
		os.Exit(1)
	}

	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Usage: `go run . <int>`", err)
		os.Exit(1)
	}

	var count int
	for i := 0; i < 64; i++ {
		count += 1 & (x >> i)
	}

	fmt.Println(count)
}
