package main

/*
   Author  :   Liuzhen
   Ques of     [653]two-sum-iv-input-is-a-bst
   Create  :   2022-03-21 19:55:04	获取当前时间
*/

func main() {

	//ans := twoSumIvInputIsABst()
	//fmt.Println(ans)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//type TreeNode TreeNode_653

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	} else {
		if findValInTree(root.Left, k-root.Val) || findValInTree(root.Right, k-root.Val) {
			return true
		} else {
			return findTarget(root.Left, k) || findTarget(root.Right, k)
		}
	}
}

func findValInTree(root *TreeNode, k int) bool {
	if root != nil {
		return root.Val == k || findValInTree(root.Left, k) || findValInTree(root.Right, k)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
