package main

import "fmt"

/*
   Author  :   Liuzhen
   Ques of     [432]all-oone-data-structure
   Create  :   2022-03-16 17:52:42	获取当前时间
*/

func main() {

	ans := AllOne()
	fmt.Println(ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
type KeyPair struct {
	key   string
	value int
}

type AllOne struct {
	Value  map[string]int
	MinOne KeyPair
	MaxOne KeyPair
}

func Constructor() AllOne {
	return AllOne{Value: make(map[string]int), MinOne: KeyPair{"", 1}, MaxOne: KeyPair{"", 0}}
}

func (this *AllOne) Inc(key string) {
	if len(this.Value) == 0 {
		this.MinOne = KeyPair{"", 1}
		this.MaxOne = KeyPair{"", 0}
	} else {
		i, exist := this.Value[key]
		if !exist {
			this.Value[key] = 1
		} else {
			this.Value[key]++
			if this.MaxOne.value <= (i + 1) {
				this.MaxOne = KeyPair{key, i + 1}
			}
			if this.MinOne.key == key {
				this.MinOne.value++
			}
		}
	}

}
func findMinOne(set map[string]int) KeyPair {
	minOne := KeyPair{"", 0}
	for key, value := range set {
		if minOne.value > value || minOne.value == 0 {
			minOne.value = value
			minOne.key = key
		}
	}
	return minOne
}

func findMaxOne(set map[string]int) KeyPair {
	maxOne := KeyPair{"", 0}
	for key, value := range set {
		if maxOne.value < value || maxOne.value == 0 {
			maxOne.value = value
			maxOne.key = key
		}
	}
	return maxOne
}

func (this *AllOne) Dec(key string) {
	i, _ := this.Value[key]
	if i == 1 {
		delete(this.Value, key)
		if len(this.Value) == 0 {
			this.MinOne = KeyPair{"", 1}
			this.MaxOne = KeyPair{"", 0}
		} else {
			if this.MinOne.key == key {
				this.MinOne = findMinOne(this.Value)
			}
			if this.MaxOne.key == key {
				this.MaxOne = findMaxOne(this.Value)
			}
		}
	} else {
		this.Value[key]--
		if this.MaxOne.key == key {
			this.MaxOne.value--
		}
		if this.MinOne.key == key {
			this.MinOne.value--
		}
	}

}

func (this *AllOne) GetMaxKey() string {
	return this.MaxOne.key

}

func (this *AllOne) GetMinKey() string {
	return this.MinOne.key

}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
//leetcode submit region end(Prohibit modification and deletion)
