package main

import "fmt"


func byValue(num int){
	//here the num has come by value. it will be copied to a new address
	num=10
	//so now in the new address, the value is 10 and the num variable in main it has other value.

	fmt.Println("From Byvalue",num)
}

func byRef(num *int){
	//now we are taking a referenced copy of num which holds the value 5 and modifying the num changes the origial num in main
	//we need to dereference the pointer before using
	*num+=+10
	fmt.Println("From ByRef",*num)
}


func main(){
	num:=5
	byValue(num)
	byRef(&num)

}