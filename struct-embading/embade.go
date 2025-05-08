package main

import "fmt"

type customer struct{
	id int
	name string
	age int
}

type order struct{
	id int 
	name string
	status string
	customer
}

func main(){
	order1:=order{
		id : 1,
		name: "pizza",
		status: "pending",
		customer: customer{
			id:1,
			name:"sagor",
			age:25,
		},


	}
	fmt.Println(order1)
}