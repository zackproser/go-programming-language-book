package main

import "fmt"

func main() {
	input := "123456789"
	fmt.Printf("Input: %s, comma output: %s\n", input, comma(input))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}