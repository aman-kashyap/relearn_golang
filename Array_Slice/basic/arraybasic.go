package main

import (
	"fmt"
)

func main() {
	grades := [3]int{98, 97, 96}
	names := [...]string{"mohan", "sohan", "ram", "shyam"}
	var students [5]string
	fmt.Printf("Grades: %v\nNames: %v\nStudents: %v\n", grades, names, students)
	students[0] = "Geeta"
	students[2] = "ramesh"
	students[1] = "suresh"
	students[3] = "parth"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Students: #0 %v\n", students[0])
	fmt.Printf("Number of Students: %v\n", len(students))
}
