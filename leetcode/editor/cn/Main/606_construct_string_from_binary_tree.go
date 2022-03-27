package main

import (
	"fmt"
	"strings"
)

/*
   Author  :   Liuzhen
   Ques of     [606]construct-string-from-binary-tree
   Create  :   2022-03-19 12:25:12	获取当前时间
*/

func main() {

	//ans := constructStringFromBinaryTree()
	//fmt.Println(ans)
}

//Definition for a binary tree node.
type TreeNode_606 struct {
	Val   int
	Left  *TreeNode_606
	Right *TreeNode_606
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func tree2str(root *TreeNode_606) string {
	ans := myTree2str(root, false)

	ans = strings.TrimSuffix(strings.TrimPrefix(ans, "("), ")")
	return ans
}

func myTree2str(root *TreeNode_606, left bool) string {
	var ans string
	if root != nil {
		ans = fmt.Sprintf("(%d%s%s)", root.Val, myTree2str(root.Left, root.Right != nil), myTree2str(root.Right, false))
	} else if left {
		ans = "()"
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
