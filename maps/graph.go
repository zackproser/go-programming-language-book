package main

import "fmt"

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func main() {
	fmt.Printf("Graph: %v\n", graph)
	addEdge("wikka", "wakka")
	addEdge("wakka", "woo")
	fmt.Printf("Graph: %v\n", graph)

	fmt.Printf("hasEdge(\"wikka\", \"wakka\"): %v\n", hasEdge("wikka", "wakka"))
	fmt.Printf("hasEdge(\"wikka\", \"woo\"): %v\n", hasEdge("wikka", "woo"))
}
