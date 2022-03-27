package main

import (
	"fmt"
	"strings"
)

/*
   Author  :   Liuzhen
   Ques of     [剑指 Offer II 005]aseY1I
   Create  :   2022-03-07 12:53:26	获取当前时间
*/
func main() {

	ans := maxProduct([]string{"abcw", "baz", "foo", "bar", "fxyz", "abcdef"})
	fmt.Println(ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func maxProduct(words []string) int {
	result := 0
	var lenthledger []int
	for _, word := range words {
		lenthledger = append(lenthledger, len(word))
	}
	for i, word := range words {
		for j, word2 := range words[i+1:] {
			if !strings.ContainsAny(word, word2) {
				tmp := lenthledger[i] * lenthledger[j+i+1]
				if result < tmp {
					result = tmp
				}
			}
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
