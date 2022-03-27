package main

import (
	"fmt"
	"strings"
)

/*
   Author  :   Liuzhen
   Ques of     [2055]plates-between-candles
   Create  :   2022-03-08 22:00:50	获取当前时间
*/
func main() {

	ans := platesBetweenCandles("***|**|*****|**||**|*", [][]int{{1, 17}, {4, 5}, {14, 17}, {5, 11}, {15, 16}})
	fmt.Println(ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func platesBetweenCandles(s string, queries [][]int) []int {
	var ans []int
	for _, query := range queries {
		subString := s[query[0] : query[1]+1]
		ans = append(ans, countPlates(subString))
	}
	return ans
}
func countPlates(s string) int {
	//inPlatesFlag := false
	tmpPlates := strings.Count(s, "*")
	if tmpPlates > len(s)-2 {
		return 0
	}
	suffixNum := len(s) - strings.LastIndex(s, "|") - 1
	prefixNum := strings.Index(s, "|")
	plates := tmpPlates - suffixNum - prefixNum
	//for _, char := range s {
	//				if inPlatesFlag {
	//		} else {
	//			inPlatesFlag = !inPlatesFlag
	//		}
	//	} else {
	//		if inPlatesFlag {
	//			tmpPlates++
	//		}
	//	}
	//	//if inPlatesFlag {
	//	//	if char == '*' {
	//	//		plates++
	//	//	} else {
	//	//		inPlatesFlag = !inPlatesFlag
	//	//	}
	//	//} else {
	//	//	if char == '|' {
	//	//		inPlatesFlag = !inPlatesFlag
	//	//	}
	//	//}
	//}
	return plates
}

//leetcode submit region end(Prohibit modification and deletion)
