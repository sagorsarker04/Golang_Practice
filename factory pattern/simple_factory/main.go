package main

import "fmt"

type Burger interface{
	prepare()
}

type SimpleBurger struct{

}

type PremiumBurger struct{

}

func (s *SimpleBurger) prepare(){
	fmt.Println("Creating a simple burger")
}

func (p PremiumBurger) prepare(){
	fmt.Println("Creating a premium burger")
}

type Factory struct{
}

func (f *Factory) creator(name string)Burger{
	if name=="basic"{
		return &SimpleBurger{}
	}
	if name=="premium"{
		return &PremiumBurger{}
	}
	return &SimpleBurger{}
	
}

func main(){
	objectCreator:=Factory{}

}