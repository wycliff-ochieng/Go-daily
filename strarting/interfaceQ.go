package main

import "fmt"

func main() {
	fmt.Println("quest")
	rec := rectangle{length: 64.38, width: 8.38}
	circ := circle{radius: 76.0}
	operations(rec)
	operations(circ)
	hum := Human{language: "English", kind: "European"}
	speaking(hum)
	got := Goat{language: "Bleets", kind: "Malian"}
	speaking(got)
}

type Shape interface {
	Area() float64
}

type rectangle struct {
	length float64
	width  float64
}
type circle struct {
	radius float64
}

func (r rectangle) Area() float64 {
	jibu := r.length * r.width
	return jibu
}
func (c circle) Area() float64 {
	ans := 3.14 * c.radius * c.radius
	return ans
}
func operations(s Shape) {
	fmt.Println(s.Area())
}

func PrintArea(s Shape) {
	fmt.Println(s.Area())
}

type Speaker interface {
	speak() string
}

type Human struct {
	language string
	kind     string
}
type Goat struct {
	language string
	kind     string
}

func (h Human) speak() string {
	how := fmt.Sprintln("A person who speaks", h.language, "is an", h.kind)
	return how
}
func (g Goat) speak() string {
	how2 := fmt.Sprintln("An animal that", g.language, "is a goat from", g.kind)
	return how2
}
func speaking(s Speaker) {
	fmt.Println(s.speak())
}
