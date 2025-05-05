package main

import "fmt"

func add(a int, b int) int{
	return a+b;
}

func substact(a int, b int) int{
	return a-b;
}

func multiply(a int,b int) int{
	return a*b
}
func calculate(x int, y int , f func(int , int)int)int {
	res:=f(x,y)
	return res
}

func calculateFloat(x int, y int , f func(int,int)float32)float32{
	res:=f(x,y)
	return res
}

func division(a int, b int)float32{
	if(b!=0){
		var res float32
		if b!=0{
			res=float32(a)/float32(b)
			return res
		}
		
	}
	return -1
}

func input()(int,int){
	var val1 int
	var val2 int
	fmt.Println("Enter value 1")
	fmt.Scanln(&val1)
	fmt.Println("Enter value 2")
	fmt.Scanln(&val2)
	return val1,val2
}

func main(){
	fmt.Println("Press 1 for addition\nPress 2 for substraction \nPress 3 for multiplication\nPress 4 for devision")
	var op int
	fmt.Scanln(&op)
	val1,val2:=input()
	if(op==1){
		res:=calculate(val1,val2,add)
		fmt.Println(res)
	}

	if(op==2){
		res:=calculate(val1,val2,substact)
		fmt.Println(res)
	}

	if(op==3){
		res:=calculate(val1,val2,multiply)
		fmt.Println(res)
	}
	if(op==4){
		res:=calculateFloat(val1,val2,division)
		fmt.Println(res)
	}

	
	
}