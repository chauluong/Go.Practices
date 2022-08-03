package main

import (
	"fmt"
)

func main() {
	fmt.Println("------Chapter 02----------")

	complexType()

}

func complexType() {
	//Complex Types

	x := complex(2.5, 3.1)

	y := complex(10.2, 2)

	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(real(x))
	fmt.Println(imag(y))

}
