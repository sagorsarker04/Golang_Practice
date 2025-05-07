package main

import "fmt"

func outer() func() int{
	counter:=0
	return func()int{
		counter+=1
		return counter
	}
}
func main(){
	//called the outer function once. it returns a function which manupulated a variable which is outside the scope of oouter function. this is closures
	func1:=outer()
	fmt.Println(func1())
	fmt.Println(func1())
}