package main

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"log"
)

const prompt = `Please enter number of operation:
1.Create new account
2.Show detail of account
3.Deposit
4.Withdraw
5.Make transfer
6.List account by Id
7.List account by balance
8.Delete account by balance
9.Exit`

func main() {
	fmt.Println("Welcome bank of orm")
Exit:
	for {
		fmt.Println(prompt)
		var num int
		fmt.Scanf("%d\n", &num)

		switch num {
		case 1:
			fmt.Println("Please enter <name> <balance>")
			var name string
			var balance float64
			fmt.Scanf("%s %f", &name, &balance)
			if err := newAccount(name, balance); err != nil {
				fmt.Println(err)
			}
		case 2:
			fmt.Println("Please enter <id>")
			var id int64
			fmt.Scanf("%d\n", &id)
			a, err := getAccount(id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("%#v\n", a)
			}
		case 3:
			fmt.Println("Please enter <id> <deposit>")
			var id int64
			var deposit float64
			fmt.Scanf("%d %f\n",&id,&deposit)
			a,err := makeDeposit(id,deposit)
			if err != nil{
				fmt.Println(err)
			}else{
				fmt.Printf("%#v\n",a)
			}
		case 4:
			fmt.Println("Please enter <id> <withdraw>")
			var id int64
			var withdraw float64
			fmt.Scanf("%d %f\n",&id,&withdraw)
			a,err := makeWithdraw(id,withdraw)
			if err!=nil{
				fmt.Println(err)
			}else{
				fmt.Println("%#v\n",a)
			}

		case 5:
			fmt.Println("Please enter <id> <balance> <id>")
			var id1 int64
			var balance float64
			var id2 int64
			fmt.Scanf("%d %f %d\n",&id1,&balance,&id2)
			err := makeTransfer(id1,id2,balance)
			if err!=nil{
				fmt.Println(err)
			}
		case 6:

		case 9:
			break Exit

		}
	}
}

type Account struct {
	Id      int64
	Name    string `xorm:"unique"`
	Balance float64
	Version int `xorm:"version"`
}

var x *xorm.Engine

func init() {
	var err error
	x, err = xorm.NewEngine("sqlite3", "./bank.db")
	if err != nil {
		log.Fatalf("Fail to create engine %v", err)
	}

	if err = x.Sync(new(Account)); err != nil {
		log.Fatalf("Fail %v", err)
	}
}

// 创建新用户
func newAccount(name string, balance float64) error {
	_, err := x.Insert(&Account{Name: name, Balance: balance})
	return err
}

// 获取账户
func getAccount(id int64) (*Account, error) {
	a := &Account{}
	has, err := x.Id(id).Get(a)

	if err != nil {
		return nil, err
	} else if !has {
		return nil, errors.New("Account not found")
	}
	return a, nil
}

// 存钱
func makeDeposit(id int64,deposit float64)(*Account,error){
	a,err := getAccount(id)
	if err!=nil{
		return nil,err
	}
	a.Balance +=deposit
	_,err = x.Update(a)
	return a,err
}

// 取钱
func makeWithdraw(id int64,withdraw float64)(*Account,error){
	a,err := getAccount(id)
	if err != nil{
		return nil,err
	}
	if a.Balance <= withdraw{
		return nil,errors.New("Not enough balance")
	}

	a.Balance -= withdraw
	_,err = x.Update(a)
	return a,err
}

// 转账
func makeTransfer(id1,id2 int64,balance float64)error{
	_,err := makeWithdraw(id1,balance)
	if err != nil{
		return err
	}
	_,err = makeDeposit(id2,balance)
	if err!= nil{
		return err
	}
	return nil
}

// 根据id获取所有的账户




















