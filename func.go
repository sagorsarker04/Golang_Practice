package main

import "fmt"

func add(a int, b int) int{
	return a*b;
}

func process(x int, y int , f func(int , int)int)int {
	res:=f(x,y)
	return res;
}

func main(){
	res:=process(3,5,add)
	fmt.Println(res)
}