package main

import "fmt"

func main() {
	/*for i := 0; i <= 10; i++ {
		//fmt.Println(i)
	}
	j := 10
	for j <= 30 {
		//fmt.Println(j)
	}
	//for, four ways*/
	samples := []string{"hello", "Melisa", "Makena", "Karen"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
	}
}
