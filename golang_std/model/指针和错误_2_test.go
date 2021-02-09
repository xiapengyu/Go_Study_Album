package main

import (
	"fmt"
	"testing"
)

type Bitcoin struct {
	balance int
}

func (w *Bitcoin) GetBalance() int {
	return w.balance
}

func (w *Bitcoin) Deposit(amount int){
	fmt.Println("address of balance in Deposit is", &w.balance)
	w.balance += amount
}

func TestBitcoin(t *testing.T) {
	wallet := Bitcoin{}
	//在 Go 中，当调用一个函数或方法时，参数会被复制，两次打印的对象地址不一致，
	//当我们修改balance的值的时候，修改的是钱包的一个副本，而不知我们创建的钱包,
	//我们可以用 指针 来解决这个问题。指针让我们 指向 某个值，然后修改它。
	//所以，我们不是拿钱包的副本，而是拿一个指向钱包的指针，这样我们就可以改变它
	wallet.Deposit(10)

	got := wallet.GetBalance()

	fmt.Println("address of balance in Deposit is", &wallet.balance)

	want := 10

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
