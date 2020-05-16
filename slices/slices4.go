package main

import "fmt"

func main() {
	var s []int // len(s) == 0, s == nil
	s = nil // len(s) == 0, s == nil
	s = []int(nil) // len(s) == 0, s == nil
	s = []int{} // len(s) == 0, s != nil

	fmt.Println(s)
}

