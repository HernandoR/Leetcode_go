package main

import "fmt"

/*
   Author  :   Liuzhen
   Ques of     [328]odd-even-linked-list
   Create  :   2022-03-07 13:58:52	获取当前时间
*/
func main() {

	var endNode, startNode *ListNode

	for i, val := range []int{2} {

		if i == 0 {

			tmpNode := ListNode{
				Val:  val,
				Next: nil,
			}
			endNode = &tmpNode
			startNode = endNode
		} else {
			tmpNode := ListNode{
				Val:  val,
				Next: nil,
			}
			endNode.Next = &tmpNode
			endNode = endNode.Next
		}
	}

	ans := oddEvenList(startNode)
	//fmt.Println(ans)
	for ans != nil {

		fmt.Printf("%d ", ans.Val)
		ans = ans.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

func oddEvenList(head *ListNode) *ListNode {
	if (head == nil) || (head.Next == nil) || (head.Next.Next == nil) {
		return head
	}
	var (
		oddList     *ListNode = head
		evenList    *ListNode = head.Next
		currentNode           = evenList.Next
		isodd                 = true
	)
	firstEvenAdd := head.Next
	for currentNode != nil {
		switch isodd {
		case false:
			evenList.Next = currentNode
			evenList = evenList.Next
		case true:
			oddList.Next = currentNode
			oddList = oddList.Next
		}

		isodd = !isodd
		currentNode = currentNode.Next
	}
	oddList.Next = firstEvenAdd
	evenList.Next = nil
	return head
}

//leetcode submit region end(Prohibit modification and deletion)
