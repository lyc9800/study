//简单的银行系统，具有存款，查询，取钱的功能
//
//

package main

import "fmt"

type Account struct {
	Acconut string
	Pwd     string
	Balance float64
}
var n string
var p string


func Code() {
	fmt.Println("请输入您的账号:")
	fmt.Scanln(&n)
	fmt.Println("请输入您的密码:")
	fmt.Scanln(&p)
}
//存款
func (account *Account) Deposit(money float64, pwd string) {
	// fmt.Println("请输入您的账号:")
	// fmt.Scanln(&n)
	// fmt.Println("请输入您的密码:")
	// fmt.Scanln(&p)
	Code()
	if p != account.Pwd {
		fmt.Println("存款……您输入的密码错误!，请确收后再次输入……")
		return
	}
	account.Balance += money
	fmt.Println("存款成功……")
	fmt.Printf("存款……您的账号是:%v您的余额是:%v\n", account.Acconut, account.Balance)
}

//查询
func (account *Account) Query(money float64,pwd string) {
	// fmt.Println("请输入您的账号:")
	// fmt.Scanln(&n)
	// fmt.Println("请输入您的密码:")
	// fmt.Scanln(&p)
	Code()
	if p != account.Pwd {
		fmt.Println("查询……您输入的密码错误!，请确收后再次输入……")
		return
	}
	fmt.Println("查询成功……")
	fmt.Printf("查询……您的账号是:%v您的余额是:%v\n", account.Acconut, account.Balance)
}

//取款
func (account *Account) Withdrawals(money float64, pwd string) {
	// fmt.Println("请输入您的账号:")
	// fmt.Scanln(&n)
	// fmt.Println("请输入您的密码:")
	// fmt.Scanln(&p)
	Code()
	if p != account.Pwd || account.Balance < 0 {
		fmt.Println("取款……您输入的密码错误!，请确收后再次输入……")
		return
	}
	account.Balance -= money
	fmt.Println("取款成功……")
	fmt.Printf("您的账号是:%v您的余额是:%v\n", account.Acconut, account.Balance)
}

func main() {
	account := Account{
		Acconut: "js666",
		Pwd:     "123456",
		Balance: 100.00,
	}
	var num int
	var n float64
	var p string
	// account.Deposit(100.00,"123456")
	// account.Query("123456")
	// account.Withdrawals(50.00,"123456")
	fmt.Printf("1---存款\n2---查询\n3---取款\n请输入您要选择的功能:")
	fmt.Scanln(&num)
	switch num {
	case 1:
		account.Deposit(n, p)
	case 2:
		account.Query(n,p)
	case 3:
		account.Withdrawals(n,p)
	default:
		fmt.Println("输入有误……")
	}
}
