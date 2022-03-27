package main

import (
	"fmt"
)

/*
   Author  :   Liuzhen
   Ques of     [2100]find-good-days-to-rob-the-bank
   Create  :   2022-03-06 15:33:16	获取当前时间
*/

func main() {
	security := []int{30000, 29999, 29998, 29997, 29996, 29995, 29994, 29993, 29992, 29991, 29990, 29989, 29988, 29987, 29986, 29985, 29984, 29983, 29982, 29981, 29980, 29979, 29978, 29977, 29976, 29975, 29974, 29973, 29972, 29971, 29970, 29969, 29968, 29967, 29966, 29965, 29964, 29963, 29962, 29961, 29960, 29959, 29958, 29957, 29956, 29955, 29954, 29953, 29952, 29951, 29950, 29949, 29948, 29947, 29946, 29945, 29944, 29943, 29942, 29941, 29940, 29939, 29938, 29937, 29936, 29935, 29934, 29933, 29932, 29931, 29930, 29929, 29928, 29927, 29926, 29925, 29924, 29923, 29922, 29921, 29920, 29919, 29918, 29917, 29916, 29915, 29914, 29913, 29912, 29911, 29910, 29909, 29908, 29907, 29906, 29905, 29904, 29903, 29902, 29901, 29900, 29899, 29898, 29897, 29896, 29895, 29894, 29893, 29892, 29891, 29890, 29889, 29888, 29887, 29886, 29885, 29884, 29883, 29882, 29881, 29880, 29879, 29878, 29877, 29876, 29875, 29874, 29873, 29872, 29871, 29870, 29869, 29868, 29867, 29866, 29865, 29864, 29863, 29862, 29861, 29860, 29859, 29858, 29857, 29856, 29855, 29854, 29853, 29852, 29851, 29850, 29849, 29848, 29847, 29846, 29845, 29844, 29843, 29842, 29841, 29840}

	time := 1
	ans := goodDaysToRobBank(security, time)
	fmt.Println(ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func goodDaysToRobBank(security []int, time int) []int {
	var goodDays []int
	for i := 0; i < len(security)-2*time; i++ {
		Day := i + time
		j := Day + time

		if isGoodDay(security[i:j+1], time) {
			goodDays = append(goodDays, Day)
		}
	}

	return goodDays

}

func isGoodDay(security []int, time int) bool {
	_ = true
	tmpV := security[0]
	for _, i2 := range security[0 : time+1] {
		if i2 > tmpV {
			return false
		} else {
			tmpV = i2
		}
	}
	tmpV = security[time]
	for _, i2 := range security[time:] {
		if i2 < tmpV {
			return false
		} else {
			tmpV = i2
		}
	}
	return true

}

//leetcode submit region end(Prohibit modification and deletion)
