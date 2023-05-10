package main

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func EvaluateExpressionTree(tree *BinaryTree) int {
	switch {
	case tree.Value == -1:
		return EvaluateExpressionTree(tree.Left) + EvaluateExpressionTree(tree.Right)
	case tree.Value == -2:
		return EvaluateExpressionTree(tree.Left) - EvaluateExpressionTree(tree.Right)
	case tree.Value == -3:
		return EvaluateExpressionTree(tree.Left) / EvaluateExpressionTree(tree.Right)
	case tree.Value == -4:
		return EvaluateExpressionTree(tree.Left) * EvaluateExpressionTree(tree.Right)
	default:
		return tree.Value
	}
}
