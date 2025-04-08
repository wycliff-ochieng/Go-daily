package main

import "fmt"

func main() {

	fmt.Println("First go line!")

	//variables

	//strings
	var name string = "Wycliff"
	var name2 = "Ochieng"
	name3 := "Onyango"
	fmt.Println(name, name2, name3)

	//integers
	var number int = 65
	fmt.Println(number)

	//bits and memory
	var age1 int8 = 21
	var age2 int16 = 250
	fmt.Println(age1, age2)

	//float
	var score1 float32 = 33.44
	var score2 float64 = 38389474647.393
	fmt.Println(score1, score2)

	//printing and formating string
	fmt.Print("Hello, ")
	fmt.Print("World.\n")
	fmt.Print("Welcome\n")
	fmt.Println("My name is ", name, "I am ", age1, "years old")
	fmt.Printf("My name is %v and i ame %v years old", name2, age2)
	fmt.Printf("My name is %q and i ame %q years old", name2, age2)
}
