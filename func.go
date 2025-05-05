package main

import "fmt"

func add(a int, b int) int{
	return a+b;
}
func multiply(a int,b int) int{
	return a*b
}
func calculate(x int, y int , f func(int , int)int)int {
	res:=f(x,y)
	return res;
}

func main(){
	add:=calculate(3,5,add)
	fmt.Println(add)

	mul:=calculate(3,5,multiply)
	fmt.Println(mul)
	
}