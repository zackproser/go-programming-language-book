package main

import "fmt"

func main() {
	var a [3]int // array of integers
	fmt.Println(a[0]) // print the first element
	fmt.Println(a[len(a)-1]) // print the last element, a[2]


	// Print the indices and elements
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// Print the elements only
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
}