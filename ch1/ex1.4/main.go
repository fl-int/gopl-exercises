package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	dup()
}

func dup() {
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
