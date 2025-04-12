package main

import "fmt"

func main() {
	myDog := Dog{name: "Tommy", age: 10}
	myCat := Cat{name: "Rose", age: 13}
	actions(myDog)
	actions(myCat)
	fmt.Println("Kwak")
}

type Animal interface {
	speak()
	eat()
}

type Dog struct {
	name string
	age  int64
}

type Cat struct {
	name string
	age  int64
}

func (d Dog) speak() {
	fmt.Println(d.name, "who is", d.age, "years old is barking wooof")
}
func (d Dog) eat() {
	fmt.Println(d.name, "eating")
}
func (c Cat) speak() {
	fmt.Println(c.name, "is allso meowing")
}
func (c Cat) eat() {
	fmt.Println(c.name, "can't stop eating though he is ", c.age)
}
func actions(a Animal) {
	a.speak()
	a.eat()
}
