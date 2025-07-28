package main

import (
	"fmt"
	"practice/model"
)

func main() {
	user := models.NewUser(1, "Sagor", "sagor@gmail.com")
	fmt.Println("User:", user)
}
