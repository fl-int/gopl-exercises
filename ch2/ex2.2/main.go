package main

import (
	"fmt"
	tempconv "gopl-exercises/ch2/ex2.1"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Must have 2 args\n", os.Args[0])
		os.Exit(1)
	}

	f, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}

	switch os.Args[2] {
	case "c":
		temp := tempconv.Celsius(f)
		fmt.Printf("%s = %s\n", temp, tempconv.CToF(temp))
	case "f":
		temp := tempconv.Fahrenheit(f)
		fmt.Printf("%s = %s\n", temp, tempconv.FToC(temp))
	case "k":
		temp := tempconv.Kelvin(f)
		fmt.Printf("%s = %s\n", temp, tempconv.KToC(temp))
	default:
		fmt.Fprintf(os.Stderr, "Usage: %s <float> [c|f]\n", os.Args[0])
		os.Exit(1)
	}
}
