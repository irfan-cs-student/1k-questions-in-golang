package main

import "fmt"

// Create these methods:

// deposit(amount int)
// withdraw(amount int)
// showBalance()

// Rules:
// deposit() increases balance.
// withdraw() decreases balance only if sufficient money exists.
// Otherwise print "Insufficient balance".
// showBalance() prints the current balance

type BankAccount struct {
	owner   string
	balance int
}

func (owner *BankAccount) deposit(amount int) {

	owner.balance += amount
	fmt.Println("balace afer deposit:", owner.balance)

}
func (owner *BankAccount) withdraw(amount int) {

	if amount > owner.balance {
		fmt.Println("insufficient money")
		return
	}
	owner.balance -= amount
	fmt.Println("balace afer withdraw:", owner.balance)

}
func (owner BankAccount) checkBalance() {
	fmt.Println("total balance: ", owner.balance)
}

func main() {

	owner := BankAccount{"irfan", 770}
	fmt.Println("orignal :", owner.balance)

	//depositing balance
	owner.deposit(20)

	//withdrawing
	owner.withdraw(2000)

	//total balance
	owner.checkBalance()
}
