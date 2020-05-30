package main

import "fmt"

func main() {
	x = append(x, 1)
	x = append(x, 2, 3)
	x = append(x, 4, 5, 6)
	x = append(x, x...) // append the slice x
	fmt.Println(x)      // "[1, 2 3 4 5 6 1 2 3 4 5 6]"
}
