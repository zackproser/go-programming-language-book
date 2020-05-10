package main

import "fmt"

func main() {
	s := "left foot"
	t := s
	s += ", root foot"

	fmt.Println(s)
	fmt.Println(t)
}