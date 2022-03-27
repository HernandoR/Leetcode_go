package main

import (
	"fmt"
	"unicode"
)

/*
   Author  :   Liuzhen
   Ques of     [917]reverse-only-letters
   Create  :   2022-03-07 13:19:07	获取当前时间
*/
func main() {

	ans := reverseOnlyLetters("ab-cd")
	fmt.Println(ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func reverseOnlyLetters(s string) string {
	var letterStack []int32
	var puncStack []int32
	var puncMark []int
	for i, char := range s {
		if unicode.IsLetter(char) {
			letterStack = append(letterStack, char)
		} else {
			puncStack = append(puncStack, char)
			puncMark = append(puncMark, i)
		}
	}
	var (
		letterP   = len(letterStack) - 1
		puncP     = 0
		puncMarkP = 0
	)

	var str string

	for i := 0; i < len(s); i++ {
		if (puncMarkP < len(puncMark)) && (i == puncMark[puncMarkP]) {
			// pop punc
			str = str + string(puncStack[puncP])
			puncMarkP++
			puncP++
		} else {
			str = str + string(letterStack[letterP])
			letterP--
		}
	}
	return str
}

//leetcode submit region end(Prohibit modification and deletion)
