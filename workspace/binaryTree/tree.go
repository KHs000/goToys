package main

import "fmt"

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

func treeWalk(tree *node) {
	fmt.Println(" --- Tree Walking --- ")
	fmt.Println("Tree root ->")
	fmt.Printf("Value: %d Left: %v Right: %v \n", tree.value, tree.left.value, tree.right.value)
	fmt.Printf("1 - Go left\n2 - Go Right\n3 - Go up\n0 - Go home\n")
	fmt.Print("Input: ")
	var input int
	fmt.Scanf("%d", &input)

	for input != 0 {
		switch input {
		case 1:
			tree = tree.left
			var lValue, rValue, pValue int
			lValue, rValue, pValue = extractSurrodingValue(tree)

			fmt.Println("Current node ->")
			fmt.Printf("Value: %d Left: %v Right: %v Parent: %v \n", tree.value, lValue, rValue, pValue)
			fmt.Print("Input: ")
			fmt.Scanf("%d", &input)
		case 2:
			tree = tree.right
			var lValue, rValue, pValue int
			lValue, rValue, pValue = extractSurrodingValue(tree)

			fmt.Println("Current node ->")
			fmt.Printf("Value: %d Left: %v Right: %v Parent: %v \n", tree.value, lValue, rValue, pValue)
			fmt.Print("Input: ")
			fmt.Scanf("%d", &input)
		case 3:
			tree = tree.parent
			var lValue, rValue, pValue int
			lValue, rValue, pValue = extractSurrodingValue(tree)

			fmt.Println("Current node ->")
			fmt.Printf("Value: %d Left: %v Right: %v Parent: %v \n", tree.value, lValue, rValue, pValue)
			fmt.Print("Input: ")
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
