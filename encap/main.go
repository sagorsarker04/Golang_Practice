package main

import (
   
    "encap/bank"  // Importing the 'bank' package
)

func main(){
	sagor:=&bank.Bank{
		Name:"sagor",
		Balance:1000,
	}

	sagor.Deposit(1000)
	sagor.WithDraw(2000)
}