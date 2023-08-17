package main

func buildTree(inorder []int, postorder []int) (root *TreeNode) {
	if len(postorder) <= 0 {
		return
	}
	val := postorder[len(postorder)-1]
	postorder = postorder[:len(postorder)-1]
	root = &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
	index := findIndex(inorder, val)
	root.Left = buildTree(inorder[:index], postorder[:index])
	root.Right = buildTree(inorder[index+1:], postorder[index:])
	return
}

func findIndex(slice []int, val int) int {
	for i, v := range slice {
		if v == val {
			return i
		}
	}
	return -1
}
