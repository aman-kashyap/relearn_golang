package main

import (
	"fmt"
)

func main() {
	a := [...]int{23, 24, 25}
	b := a  //b copies complete value of array a
	c := &a //c is pointer to a Means c will poit to the data that a has
	b[1] = 26
	c[2] = 28 //change the value of a
	fmt.Println(a, b, c)
}
