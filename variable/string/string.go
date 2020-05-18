package main

import (
	"fmt"
)

func main() {
	s := "This is a string" //string are immutable
	s2 := "This is also a string"
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", s[3], s[3])
	fmt.Printf("%v, %T\n", string(s[3]), s[3])
	fmt.Printf("%v, %T\n", s+s2, s+s2)

	//string conversion byte slice
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)
}
