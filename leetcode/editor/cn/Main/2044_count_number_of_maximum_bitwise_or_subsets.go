package main

import "fmt"

/*
   Author  :   Liuzhen
   Ques of     [2044]count-number-of-maximum-bitwise-or-subsets
   Create  :   2022-03-15 11:06:48	获取当前时间
*/
func main() {
	nums := []int{3, 1}
	ans := countMaxOrSubsets(nums)
	fmt.Println(ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func countMaxOrSubsets(nums []int) (ans int) {
	maxOr := 0
	var dfs func(int, int)
	dfs = func(pos, or int) {
		if pos == len(nums) {
			if or > maxOr {
				maxOr = or
				ans = 1
			} else if or == maxOr {
				ans++
			}
			return
		}
		dfs(pos+1, or|nums[pos]) // take
		dfs(pos+1, or)           // dont
	}
	dfs(0, 0)
	return

}

//leetcode submit region end(Prohibit modification and deletion)
