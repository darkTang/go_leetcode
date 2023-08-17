package main

func buildTree2(preorder []int, inorder []int) (root *TreeNode) {
	if len(preorder) <= 0 {
		return
	}
	val := preorder[0]
	preorder = preorder[1:]
	root = &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
	index := findIndex2(inorder, val)
	root.Left = buildTree2(preorder[:index], inorder[:index])
	root.Right = buildTree2(preorder[index:], inorder[index+1:])
	return
}

func findIndex2(slice []int, val int) int {
	for i, v := range slice {
		if v == val {
			return i
		}
	}
	return -1
}
