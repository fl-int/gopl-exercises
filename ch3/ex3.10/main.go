package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaNoRecursion(s string) string {
	if len(s) <= 3 {
		return s
	}

	var buf bytes.Buffer
	step := len(s) % 3
	if step == 0 {
		step = 3
	}
	for i := 0; i < len(s); i += step {
		if i > 0 && step != 3 {
			step = 3
		}
		buf.WriteString(s[i : i+step])
		if i < len(s)-3 {
			buf.WriteString(",")
		}
	}
	return buf.String()
}

func main() {
	i := "12345678901234567890"
	r := comma(i)
	r2 := commaNoRecursion(i)
	fmt.Println(r)
	fmt.Println(r2)
}
