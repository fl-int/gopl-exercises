package main

import (
	"fmt"
	"os"
	"strings"
)

func echo0() {
	fmt.Println("echo0:")
	fmt.Println(strings.Join(os.Args[1:], " "))
	println()
}

// 1.1
func echo1() {
	fmt.Println("echo1:")
	fmt.Println(strings.Join(os.Args[:], " "))
	println()
}

// 1.2
func echo2() {
	fmt.Println("echo2:")
	for i, v := range os.Args {
		fmt.Printf("%d:%s\n", i, v)
	}
	println()
}

// 1.3
func echo3() {
	// todo
}
