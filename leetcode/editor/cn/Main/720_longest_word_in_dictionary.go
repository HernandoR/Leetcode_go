package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
   Author  :   Liuzhen
   Ques of     [720]longest-word-in-dictionary
   Create  :   2022-03-17 22:56:19	获取当前时间
*/

func main() {

	ans := longestWord([]string{"rac", "rs", "ra", "on", "r", "otif", "o", "onpdu", "rsf", "rs", "ot", "oti", "racy", "onpd"})
	fmt.Println(ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func longestWord(words []string) string {
	sort.Strings(words)
	//println(strings.Join(words, ","))
	prefix := ""
	ans := ""
	for _, word := range words {
		if len(word) == 1 {

			if betterWord(prefix, ans) {
				ans = prefix
			}
			prefix = word
		}

		if (len(prefix) == (len(word) - 1)) && strings.HasPrefix(word, prefix) {
			prefix = word
		} else {
			if betterWord(prefix, ans) {
				ans = prefix
			}
		}
	}
	if betterWord(prefix, ans) {
		ans = prefix
	}
	return ans
}
func betterWord(new string, old string) bool {
	longer := len(new) > len(old)
	//prefer := new <= old
	//return longer || (prefer && (len(new) == len(old)))
	return longer
}

//leetcode submit region end(Prohibit modification and deletion)
