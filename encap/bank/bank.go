package bank

import "fmt"


type Bank struct{
	Name string
	Balance int
}

func(B *Bank) GetBalance() int{
return B.Balance
}

func (B *Bank) Deposit(Bal int){
	if(Bal>0){
		B.Balance+=Bal
	}
	fmt.Println("Balance is deposited current balance is ", B.Balance)
}

func (B *Bank) WithDraw(Bal int){
	if(B.Balance>Bal){
		B.Balance-=Bal
		fmt.Println("Withdrawl sucessfull Balance is",B.Balance)
	}else{
		fmt.Println("Not enough Balance")
	}
}