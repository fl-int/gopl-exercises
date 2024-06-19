package main

import (
	"fmt"
	"os"
	"strings"
)

func echo3() {
	fmt.Println("Printing command line arguments:")
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func echo4() {
	fmt.Println("Printing command line arguments:")
	fmt.Println(strings.Join(os.Args[:], " "))
}
