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

	/*scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)

	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		words[scan.Text()]++
	}
	fmt.Println(len(words))*/

}
