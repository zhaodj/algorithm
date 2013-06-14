package algorithm

type AVLTree struct {
	root *AVLNode
}

type AVLNode struct {
	TreeNode
	height int
}
