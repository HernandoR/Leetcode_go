package main

/*
   Author  :   Liuzhen
   Ques of     [2043]simple-bank-system
   Create  :   2022-03-18 14:42:49	获取当前时间
*/

func main() {

	//ans := simpleBankSystem()
	//fmt.Println(ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
type Bank struct {
	balanceList []int64
}

func Constructor(balance []int64) Bank {
	return Bank{balanceList: balance}
}

func (b Bank) isValidAccount(account int) bool {
	return account <= len(b.balanceList)

}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if this.isValidAccount(account1) && this.isValidAccount(account2) {

		account1--
		account2--

		if this.balanceList[account1] >= money {
			this.balanceList[account1] -= money
			this.balanceList[account2] += money
			return true
		}
	}
	return false
}

func (this *Bank) Deposit(account int, money int64) bool {
	if this.isValidAccount(account) {

		account--
		this.balanceList[account] += money
		return true
	}
	return false
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if this.isValidAccount(account) {
		account--
		if this.balanceList[account] >= money {
			this.balanceList[account] -= money
			return true
		}
	}
	return false
}

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */
//leetcode submit region end(Prohibit modification and deletion)
