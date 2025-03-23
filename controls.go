package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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
}
