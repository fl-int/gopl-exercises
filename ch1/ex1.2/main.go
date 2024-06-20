package main

import (
	"fmt"
	"os"
)

func main() {
	echo()
}

func echo() {
	for i, v := range os.Args {
		fmt.Printf("%d:%s\n", i, v)
	}
}
