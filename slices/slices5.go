package main

import "fmt"

func main() {
	var runes []rune
	for _, r := range "Hello, ꟹ" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes) // ['H' 'e' 'l' 'l' 'o' ',' ' ' 'ꟹ']
	fmt.Printf("%U\n", runes) // [U+0048 U+0065 U+006C U+006C U+006F U+002C U+0020 U+A7F9]
	fmt.Printf("%d\n", runes) // [72 101 108 108 111 44 32 43001]
}