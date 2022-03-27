package main

import (
	"fmt"
)

/*
   Author  :   Liuzhen
   Ques of     [590]n-ary-tree-postorder-traversal
   Create  :   2022-03-12 18:56:26	获取当前时间
*/
func main() {

	ans := postorder()
	fmt.Println(ans)
}

type Node struct {
	Val      int
	Children []*Node
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var res []int
	postorderSub(root, &res)
	return res
}

func postorderSub(root *Node, res *[]int) {

	for _, child := range root.Children {
		postorderSub(child, res)
	}
	*res = append(*res, root.Val)

}

func postorderWork() {

}

//leetcode submit region end(Prohibit modification and deletion)
