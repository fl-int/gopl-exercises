package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func echo() {
	println("echo0:")
	fmt.Println(strings.Join(os.Args[1:], " "))
	println()
}

// 1.1
func echo1() {
	println("echo1:")
	fmt.Println(strings.Join(os.Args[:], " "))
	println()
}

// 1.2
func echo2() {
	println("echo2:")
	for i, v := range os.Args {
		fmt.Printf("%d:%s\n", i, v)
	}
	println()
}

// 1.3
func echo3() {
	// todo
}

func dup() {
	println("dup:")
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup0: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// 1.4
func dup1() {
	println("dup1:")
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		var printedFilename bool
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup0: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			if !printedFilename {
				fmt.Printf("Found duplicate in file %s\n", filename)
				printedFilename = true
			}
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
