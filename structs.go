package main

import "fmt"

func main() {
	//structs-ideal for when you have releated data that you want to group together
	//type nameOfStruct struct
	//once a struct type is declared we can define variablees of that type

	type parent struct {
		name    string
		age     int
		country string
	}
	//fmt.Println(parent)
	var fred parent
	fmt.Println(fred)
	mike := parent{}
	fmt.Println(mike)
	ochola := parent{
		"Peter Ochola",
		42,
		"Kenya",
	}
	fmt.Println(ochola)
	fmt.Println(ochola.name)
}
