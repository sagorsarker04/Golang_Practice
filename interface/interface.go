package main

import "fmt"

type Speaker interface{
	speaks()
}
type Animal struct{
	name string
}
type Bird struct{
	Animal
	wing int
}
func (A *Animal) speaks(){
	fmt.Println("The animal is",A.name)
}
func (B *Bird) speaks(){
	fmt.Println("The animal is",B.name)
}

func AnimalType(s Speaker){
	s.speaks()
}

func main(){
	//this is a way to call interface functions
	//this the way to achieve polymorphism
	// var speaker Speaker
	// animal:=&Animal{name:"Lion"}
	// speaker=animal
	// speaker.speaks()
	animal:=&Bird{wing:30 , Animal: Animal{name: "Parrot"}}
	AnimalType(animal)
}