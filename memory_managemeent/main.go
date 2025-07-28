package main

import (
	"fmt"
	"unsafe"
)

func main() {
	type Inefficient struct {
		a bool
		b int64
		c bool
		d int32
	}

	type Efficient struct {
		b int64
		d int32
		a bool
		c bool
	}

	fmt.Printf("Size of Inefficient: %d bytes\n", unsafe.Sizeof(Inefficient{}))
	fmt.Printf("Size of Efficient: %d bytes\n", unsafe.Sizeof(Efficient{}))
}