package main

import "fmt"

func main() {
	a := []string{"a", "b", "c", "d", "e", "f"}
	b := []string{"a", "b", "c", "d", "e", "f"}
	c := []string{"a", "b", "c"}
	d := []string{"a", "b", "c", "d", "e", "e"}

	fmt.Printf("%v and %v are equal: %t\n", a, b, equal(a, b))
	fmt.Printf("%v and %v are equal: %t\n", a, b, equal(a, c))
	fmt.Printf("%v and %v are equal: %t\n", a, b, equal(a, d))
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}