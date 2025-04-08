package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	x := 10
	f := os.Stdin
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	day := scanner.Text()
	fmt.Println("Please type a day:", day)
	if day == "Saturday" {
		fmt.Println("Its a Weekend")
	} else {
		fmt.Println("Its a weekday")
	}
	//shadowing
	if x > 5 {
		fmt.Println(x)
		x, y := 5, 20
		fmt.Println(x, y)
	}
	fmt.Println(x)

	n := rand.Intn(10)
	if n == 0 {
		fmt.Println("Thats too low:", n)
	} else if n > 5 {
		fmt.Println("Thats a little bit big:", n)
	} else {
		fmt.Println("Thats a fair one:", n)
	}
	//for loop
	//switch (just like couple if else) 1)discrete 2)switch on true
	//
}
