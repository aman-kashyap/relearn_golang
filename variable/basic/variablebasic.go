package main

import (
	"fmt"
	"strconv"
)

var i int = 918
var n string = "42"
var o string = "42.7"
var p int = 89

func main() {
	var i int
	i = 78
	var j string
	j = "Ram"
	k := "RAMAYAN"
	l := 44.44
	var m float32 = 66.2
	//p := 90
	var q float64
	q = float64(m)
	var r string
	r = strconv.Itoa(i)

	var s int
	s, _ = strconv.Atoi(n)

	fmt.Println(i, j, k, l, m, n, o, p, q, r, s)
	fmt.Printf("i has value %v of %T type, '%v %T' \n'%v %T' '%v %T' '%v %T' '%v %T' '%v %T' '%v %T' '%v %T' \n'%v %T' '%v %T' \n", i, i, j, j, k, k, l, l, m, m, n, n, o, o, p, p, q, q, r, r, s, s)
}
