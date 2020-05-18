package main

import (
	"fmt"
)

func main() {
	var n complex128 = 1 + 2i
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", real(n), real(n))
	fmt.Printf("%v, %T\n", imag(n), imag(n))

	var p complex128 = complex(6, 3)
	fmt.Printf("%v, %T\n", p, p)

	a := 2 + 3i
	b := 1 + 3.2i
	fmt.Printf("%v, %T\n %v, %T\n", a, a, b, b)
	fmt.Println("addition : ", a+b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	//fmt.Println(a % b)
}
