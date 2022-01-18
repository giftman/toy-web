package afterclass

type Tree interface {
	PreTraverse()
	InorderTraverse()
	PostTraverse()
}

// 二叉树
type binaryTree struct {
	Data string
	Tree
	Left  *binaryTree
	Right *binaryTree
}

// 多叉树
type mutliWayTree struct {
	Data string
	Tree
	Left  *binaryTree
	Right *binaryTree
}
