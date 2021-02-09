package main

import (
	"fmt"
	"testing"
)

type Wallet struct {
	balance int
}

func (w Wallet) GetBalance() int {
	return w.balance
}

func (w Wallet) Deposit(amount int){
	fmt.Println("address of balance in Deposit is", &w.balance)
	w.balance += amount
}

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	//在 Go 中，当调用一个函数或方法时，参数会被复制，两次打印的对象地址不一致，
	//当我们修改balance的值的时候，修改的是钱包的一个副本，而不知我们创建的钱包
	wallet.Deposit(10)

	got := wallet.GetBalance()

	fmt.Println("address of balance in Deposit is", &wallet.balance)

	want := 10

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
