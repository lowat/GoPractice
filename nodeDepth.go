package main

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func NodeDepths(root *BinaryTree) int {
	return NodeDepthsHelper(root, 0)
}

func NodeDepthsHelper(node *BinaryTree, depth int) int {
	if node == nil {
		return 0
	} else {
		return depth + NodeDepthsHelper(node.Left, depth+1) + NodeDepthsHelper(node.Right, depth+1)
	}
}
