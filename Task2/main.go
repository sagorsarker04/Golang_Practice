package main

import (
	"fmt"
	"time"
)

type Vehicle struct{
	Make string
	Model string
	Year int
	Price float64
}

type Car struct{
	Vehicle
	SeatingCapacity int
}

type Bike struct{
	Vehicle
	HasCarrieer bool
}

func (C Car) ShowDetails(){
	fmt.Println("Make is", C.Make)
	fmt.Println("Model is ", C.Model)
	fmt.Println("Year is ",C.Year)
	fmt.Println("Price is ",C.Price)
	fmt.Println("Seating capacity is ",C.SeatingCapacity)
}

func (B Bike) ShowDetails(){
	fmt.Println("Make is", B.Make)
	fmt.Println("Model is ", B.Model)
	fmt.Println("Year is ",B.Year)
	fmt.Println("Price is ",B.Price)
	fmt.Println("It has career? ",B.HasCarrieer)
}

func (c Car) CalculateDepreciation() float64 {
	currentYear := time.Now().Year()
	carAge := currentYear - c.Year
	depreciationRate := 0.15

	currentValue := c.Price
	for i := 0; i < carAge; i++ {
		currentValue -= currentValue * depreciationRate
	}
	return currentValue
}

func (B Bike) CalculateDepreciationBike() float64 {
	currentYear := time.Now().Year()
	bikeAge := currentYear - B.Year
	depreciationRate := 0.15

	currentValue := B.Price
	for i := 0; i < bikeAge; i++ {
		currentValue -= currentValue * depreciationRate
	}
	return currentValue
}
func main(){
	car1:=Car{
		SeatingCapacity: 12,
		Vehicle:Vehicle{
			Make: "Honda",
			Year: 2012,
			Model: "120",
			Price: 120000,
		},

	}
	currentValue := car1.CalculateDepreciation()
	fmt.Println("Cureent value of the car is ",currentValue)

	bike1:=Bike{
		HasCarrieer: true,
			Vehicle:Vehicle{
				Make: "Phonix",
				Year: 2018,
				Model: "1203",
				Price: 13433,
		},
	}
	bikeValue:=bike1.CalculateDepreciationBike()
	fmt.Println("Current value is ", bikeValue)

	
}