package main

import (
	"fmt"
	"sync"
)

type Config struct {
	Name string
	port string
}

var (
	Once     sync.Once
	instance *Config
)

func Getconfig() *Config {
	Once.Do(func() {
		instance = &Config{
			Name: "Golang",
			port: "8080",
		}
	})
	return instance
}

func main() {
	obj := Getconfig()
	fmt.Printf("Obj1 address is %p",obj)
	obj2:=Getconfig()
	fmt.Printf("Obj2 address is %p",obj2)
}
