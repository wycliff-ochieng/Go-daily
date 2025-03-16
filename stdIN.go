package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println("Grace", scanner.Text())
	}*/

	f := os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	fmt.Println("Grace", scanner.Text())

}
