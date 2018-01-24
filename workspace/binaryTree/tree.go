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
		} else {
			tree.right = &node {nil, nil, i}
		}
	} else {
		if i < tree.value {
			insert(tree.left, i)
		} else {
			insert(tree.right, i)
		}
	}
}

func main () {
	var x = []int {
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}

	arrHead := x[0:1][0]
	tree := node {nil, nil, arrHead}

	for _, i := range x[1:] {
		insert(&tree, i)
	}

	printTree(&tree)
}