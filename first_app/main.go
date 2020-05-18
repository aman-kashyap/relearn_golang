package main

import "fmt"

func main() {
	statement := "Hello Home!"
	fmt.Println(len(statement))
	fmt.Printf("%T\n%v\n", statement, len(statement))
	fmt.Println("Hello Go HOME")
}
