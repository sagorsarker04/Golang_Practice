package main

import (
	"fmt"
	"time"
)
//thats how we declare a struct
type Order struct{
	id int
	name string
	status string
	createdAt time.Time
	
}

//constructor in go
//we return a object type by *object_name
func newOrder(id int, name string, status string) *Order {
	myOrder := Order{
		id:     id,
		name:   name,
		status: status,
	}
	return &myOrder
}
//this is how we link functions to structs as we did in classes. now o operator can access all the members of the struct.
//Convention is to use first letter of Struct name
// we are passing the reference of the struct so that we can modify the actual variables of struct
func (o *Order /*here is tell which struct is connected with this function*/) changeStatus(status string){
	o.status=status
}



func main(){
	order1:=newOrder(1,"pizza","complete")
	order1.changeStatus("Complete")
	//we can manually set or show struct variables
	order1.createdAt=time.Now()
	fmt.Println(order1)

	
	//we can shorthandly write struct
	language := struct {
		id   string
		name string
	}{
		id:   "1",
		name: "Bangla",
	}
	fmt.Println(language)
}



