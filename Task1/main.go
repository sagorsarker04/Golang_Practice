package main

import "fmt"

type Employee struct{
	Name string
	ID int
	Salary float32
}

type Manager struct{
	Employee
	TeamSize int
}

func (M Manager) Show(){
	fmt.Println("The name is",M.Name)
	fmt.Println("The ID is ",M.ID)
	fmt.Println("The Salray is ", M.Salary)
	fmt.Println("The team size is ",M.TeamSize)
}

func (E Employee) calculate(){
	fmt.Println("The name is",E.Name)
	fmt.Println("The ID is ",E.ID)
	fmt.Println("The Salray is ", E.Salary)

	TotalSalary:=(E.Salary*12)
	fmt.Println("Total Salary is",TotalSalary)
}

func main(){
	e1:=Employee{
		Name:"Sagor",
		ID: 1,
		Salary: 25000,

	}
	e1.calculate()

	m1:=Manager{
		
		TeamSize: 8,
		Employee: Employee{
			Name: "Sakib vai",
			ID: 2,
			Salary: 100000,
		},
	}
	m1.calculate()
}