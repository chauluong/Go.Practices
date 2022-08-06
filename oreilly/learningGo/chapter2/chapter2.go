package main

import (
	"fmt"
)

func main() {
	fmt.Println("------Chapter 02----------")

	complexType()
	runeliteralType()

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

func runeliteralType() {
	fmt.Println("--rune--")
	var runetype = '\141'
	fmt.Println(runetype)
}
