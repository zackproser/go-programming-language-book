package main

import "fmt"

func main () {
	s := "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7]) // "104 119" ('h' and 'w')

	fmt.Println(s[0:5]) // hello

	fmt.Println(s[:5]) // hello
	fmt.Println(s[7:]) // world
	fmt.Println(s[:]) // hello, world

	fmt.Println("goodbye" + s[5:]) // "goodbye, world"
}