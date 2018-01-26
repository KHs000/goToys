package main

import (
	"fmt"
)

type node struct {
	left   *node
	right  *node
	parent *node
	value  int
}

func printTreeOrd(node *node) {
	if node.left != nil {
		printTreeOrd(node.left)
	}
	fmt.Println(node.value)
	if node.right != nil {
		printTreeOrd(node.right)
	}
}

func insert(tree *node, i int) {
	if (i < tree.value && tree.left == nil) || (i >= tree.value && tree.right == nil) {
		if i < tree.value {
			tree.left = &node{nil, nil, tree, i}
		} else {
			tree.right = &node{nil, nil, tree, i}
		}
	} else {
		if i < tree.value {
			insert(tree.left, i)
		} else {
			insert(tree.right, i)
		}
	}
}

func extractSurrodingValue(tree *node) (l, r, p int) {
	l = 0
	if tree.left != nil {
		l = tree.left.value
	}
	r = 0
	if tree.right != nil {
		r = tree.right.value
	}
	p = 0
	if tree.parent != nil {
		p = tree.parent.value
	}

	return
}

func takeInput(intput *int) {
	fmt.Print("Input: ")
	fmt.Scanf("%d", &input)
}

func printWalkStep(v, l, r, p int) {
	fmt.Print("Current node -> ")
	fmt.Printf("Value: %d Left: %v Right: %v Parent: %v \n", v, l, r, p)
	fmt.Print("Input: ")
}

func treeWalk(tree *node) {
	/* Menu interface */
	fmt.Println(" --- Tree Walking --- ")
	fmt.Print("Tree root -> ")
	fmt.Printf("Value: %d Left: %v Right: %v \n", tree.value, tree.left.value, tree.right.value)
	fmt.Printf("1 - Go left\n2 - Go Right\n3 - Go up\n0 - Go home\n")

	/* Initial input */
	var input *int

	for input != 0 {
		switch input {
		case 1:
			if tree.left == nil {
				panic("I cannon't move that way")
			}
			tree = tree.left
			lValue, rValue, pValue := extractSurrodingValue(tree)
			printWalkStep(tree.value, lValue, rValue, pValue)
			fmt.Scanf("%d", &input)
		case 2:
			tree = tree.right
			lValue, rValue, pValue := extractSurrodingValue(tree)
			printWalkStep(tree.value, lValue, rValue, pValue)
			fmt.Scanf("%d", &input)
		case 3:
			tree = tree.parent
			lValue, rValue, pValue := extractSurrodingValue(tree)
			printWalkStep(tree.value, lValue, rValue, pValue)
			fmt.Scanf("%d", &input)
		case 0:
			fmt.Println("Exiting...")
		}
	}
}

func main() {
	var x = []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	arrHead := x[0:1][0]
	tree := node{nil, nil, nil, arrHead}

	for _, i := range x[1:] {
		insert(&tree, i)
	}

	treeWalk(&tree)
	printTreeOrd(&tree)
}
