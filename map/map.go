package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is map's main function")
	fmt.Println("This is mapBasic function")
	mapBasic()
	fmt.Println("This is mapKey function")
	mapKey()
	mapMake()
}

func mapBasic() {
	statePopulation := map[string]int{ //map takes a key(string) and maps that to a value(int)
		"UP":    230000000000,
		"Bihar": 8000000000,
		"MP":    4000000000,
	}
	fmt.Println(statePopulation)         //map doesnot return the value as provided
	statePopulation["Delhi"] = 600000000 // add another key to map
	fmt.Println(statePopulation)
	fmt.Println(statePopulation["Delhi"])
	delete(statePopulation, "Delhi") //delete a key from map
	fmt.Println(statePopulation)
	pop0 := statePopulation["Bih"] //verifying key in the map
	fmt.Println(pop0)
	pop, ok := statePopulation["Bih"] //verifying key in the map
	fmt.Println(pop, ok)
	pop1, ok := statePopulation["Bihar"]
	fmt.Println(pop1, ok)

}

func mapKey() {
	statePopulation := map[string]int{ //map takes a key(string) and maps that to a value(int)
		"UP":    230000000000,
		"Bihar": 8000000000,
		"MP":    4000000000,
	}
	//m := map[[]int]string{}	//slice is an invalid map key type
	m := map[[4]int]string{}
	fmt.Println(statePopulation, m)
	fmt.Println(statePopulation["UP"])
}

func mapMake() {
	statePopulation := make(map[string]int)
	statePopulation = map[string]int{ //map takes a key(string) and maps that to a value(int)
		"UP":    230000000000,
		"Bihar": 8000000000,
		"MP":    4000000000,
	}
	fmt.Println(statePopulation)
	_, ok := statePopulation["UP"]
	fmt.Println(ok)
	fmt.Println(len(statePopulation))
}
