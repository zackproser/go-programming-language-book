package main

import "fmt"

func main() {
	// Define an array with 100 elements, all zero except for the last, which has value -1
	r := [...]int{99: -1}
	fmt.Println(r)

	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // true false false

	d := [3]int{1, 2}
	fmt.Println(a == d) // compile error: cannot compare [2]int == [3]int
}