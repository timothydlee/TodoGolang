package main

import (
	"fmt"
)

func main() {
	// var myInt int
	// myInt = 42

	var myFloat32 float32 = 42.
	fmt.Println(myFloat32)

	myString := "Hello Go"
	fmt.Println(myString)

	myComplex := complex(2, 3)
	fmt.Println(myComplex)
	fmt.Println(real(myComplex))
}
