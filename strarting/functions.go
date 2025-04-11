package main

import (
	"fmt"
	"os"
)

var tasks []string

func addTask(task string) string {
	tasks = append(tasks, task)
	fmt.Println("Task added :", task)
}

func main() {
	/*//ans := add(12, 12)
	//fmt.Println(ans)
	var ans = add(20, 20)
	fmt.Println(ans)
	fmt.Println(add(20, 20))
	greet("Melisa")*/

	if len(os.Args) < 2 {
		fmt.Println("Kindly add a valid task")
		return
	}
	task := os.Args[1]
	addTask(task)

	fmt.Println("\n📝 Current tasks:")
	for i, t := range tasks {
		fmt.Printf("%d. %s\n", i+1, t)

	}
	/*func add(a int, b int) int {
	  	return a + b
	  }
	  func greet(name string) string {
	  	maneno := fmt.Sprintf("Hello my good friend %s", name)
	  	//fmt.Println("helloo my good friend %s \n", name)
	  	fmt.Println(maneno)
	  	return maneno*/
}
