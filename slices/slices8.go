package main

import "fmt"

func main() {
	s := []string{
		"wikka",
		"wakka",
		"woo",
		"",
		"",
		"woo-woo-woo",
	}
	fmt.Printf("Initial slice: %v\n", s)

	ne := nonempty(s)

	fmt.Printf("Procssed slice: %v\n", ne)

}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
