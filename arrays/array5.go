package main

import "fmt"

func main() {
	r := [32]byte{31:1}
	s := [32]byte{1, 2, 3, 4, 5}
	fmt.Println("Initial values:")
	fmt.Printf("r: %v\ns: %v\n", r, s)
	fmt.Println()

	zero(&r)
	efficientZero(&s)

	fmt.Println("Post-zero'd values:")
	fmt.Printf("r: %v\ns: %v\n", r, s)
}

// Inefficiently iterates through entire array to zero out every element
func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

// More efficiently zeros out the array in a single operation
func efficientZero(ptr *[32]byte) {
	*ptr = [32]byte{}
}