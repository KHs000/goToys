package main

import "fmt"

type node struct {
	left *node
	right *node
	value int
}

func printTree (node *node) {
	if node.left != nil {printTree(node.left)}
	fmt.Println(node.value)
	if node.right != nil {printTree(node.right)}
}

func insert (tree *node, i int) {
	if (i < tree.value && tree.left == nil) || (i >= tree.value && tree.right == nil) {
		if i < tree.value {
			tree.left = &node {nil, nil, i}
		} else {tree.right = &node {nil, nil, i}}
	} else {
		if i < tree.value {
			insert(tree.left, i)
		} else {insert(tree.right, i)}
	}
}

func main () {
	tree := node {nil, nil, 4}

	insert(&tree, 3)
	insert(&tree, 5)
	insert(&tree, 2)

	printTree(&tree)
}