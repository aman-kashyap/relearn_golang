package main

import (
	"fmt"
)

func sliceBasic() {
	a := []int{23, 24, 25, 26}
	b := a // b is pointing to slice a any change in underlying data of b will change a
	b[2] = 67
	//a[2] = 77

	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("Length: %v \n", len(a))
	fmt.Printf("Capacity: %v \n", cap(a))
}

func slice2() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //
	b := a[:]
	c := a[3:]
	d := a[:3]
	e := a[4:8]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}

func makeSlice() {
	a := make([]int, 3, 100) // Slice of Length 3 with capacity of 100 element
	fmt.Println(a)
	fmt.Printf("Length: %v \n", len(a))
	fmt.Printf("Capacity: %v \n", cap(a))
}

func sliceCapacity() {
	a := []int{}
	fmt.Println(a)
	fmt.Printf("Length: %v \n", len(a))
	fmt.Printf("Capacity: %v \n", cap(a))
	a = append(a, 1) //append allows to push element to stack
	fmt.Println(a)
	fmt.Printf("Length: %v \n", len(a))
	fmt.Printf("Capacity: %v \n", cap(a))
	a = append(a, 1, 2, 3, 4, 5)
	b := []int{44, 33, 22}
	a = append(a, b...)
	fmt.Println(a, b)
	fmt.Printf("Length: %v \n", len(a))
	fmt.Printf("Capacity: %v \n", cap(a))
}

func popSlice() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Printf("Lenght %v\n", len(a))
	fmt.Println("Element in a", a)
	b := a[1:]
	fmt.Println(b)
	c := a[:len(a)-1]
	fmt.Println(c)
	d := append(a[:2], a[3:]...)
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)
	fmt.Println(d)
}
func main() {
	fmt.Println("This is main Function")
	sliceBasic()
	slice2()
	makeSlice()
	sliceCapacity()
	popSlice()
}
