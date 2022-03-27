package main

import (
	"fmt"
	"strconv"
)

/*
   Author  :   Liuzhen
   Ques of     [504]base-7
   Create  :   2022-03-07 10:13:34	获取当前时间
*/
func main() {
	ans := convertToBase7(100)
	fmt.Println(ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func convertToBase7(num int) string {
	//newbit := ""
	str7 := ""
	//for num > 7 {
	//	newbit = fmt.Sprintf("%d", num%7)
	//	num = num / 7
	//	str7 = strings.Join([]string{str7, newbit}, "")
	//}
	str7 = strconv.FormatInt(int64(num), 7)
	return str7

}

//leetcode submit region end(Prohibit modification and deletion)
