package main

func pathSum(root *TreeNode, targetSum int) (res [][]int) {
	var traverse func(node *TreeNode, targetSum int,temp *[]int, res *[][]int)
	traverse = func(node *TreeNode, targetSum int, temp *[]int, res *[][]int) {
		if node == nil {
			return
		}
		*temp = append(*temp, node.Val)
		targetSum-=node.Val
		if node.Left == nil && node.Right == nil  {
			if targetSum == 0 {
				pathCopy := make([]int, len(*temp))
				for i, element := range *temp {
					pathCopy[i] = element
				}
				*res = append(*res, pathCopy)
			}
		}
		traverse(node.Left, targetSum, temp, res)
		traverse(node.Right, targetSum,temp, res)
		*temp = (*temp)[:len(*temp)-1]
	}
	traverse(root, targetSum, new([]int), &res)
	return
}
