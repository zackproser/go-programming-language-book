package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	fmt.Printf("Added node: %d\n", value)
	return t
}

func (root *tree) PrintTree() {
	if root == nil {
		return
	}

	root.left.PrintTree()
	fmt.Printf("%d,", root.value)
	root.right.PrintTree()
}

func (t *tree) PrintNode() {
	fmt.Printf("PrintNode(): value: %v\n left: %v\n right: %v\n", t.value, t.left, t.right)
}

func main() {
	tp := add(nil, 5)
	tp2 := add(tp, 7)

	tp2.PrintNode()

	tp3 := add(tp2, 2)

	tp3.PrintNode()

	tp4 := add(tp3, 1)

	tp4.PrintNode()

	fmt.Println("Print tree in order:")
	tp.PrintTree()
}
