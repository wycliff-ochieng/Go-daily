package main

import "fmt"

func main() {
	fmt.Println("Questions on go")
	person := Person{name: "Wambulu", age: 56}
	person.introduce()
	rect := Rectangle{length: 56.9, width: 65.3}
	fmt.Println(rect.area())
	review := Book{author: "Leornado DiCaprio", title: "Killers of the flower moon"}
	fmt.Println(review.summary())
}

type Person struct {
	name string
	age  int64
}

func (p Person) introduce() {
	fmt.Println("My name is", p.name, "and i am", p.age, "years old")
}

type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) area() float64 {
	return r.length * r.width
}

type Book struct {
	author string
	title  string
}

func (b Book) summary() string {
	summ := fmt.Sprintln("The book", b.title, "by", b.author, "narrates about how a small village called olage was the richest how trying to illegally acquire the riches led to the rise of the FBI")
	return summ
}
