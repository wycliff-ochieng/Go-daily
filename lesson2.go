package main

import "fmt"

func main() {
	var ages [3]int = [3]int{20, 10, 34}
	fmt.Println(ages, len(ages))

	name := [3]string{"Ochorokodi", "Cliff", "Ohangla"}
	fmt.Println(name, len(name))

	//slices(arrays)
	score := []int{3, 24, 56, 78}
	fmt.Println(score)
	score = append(score, 245)
	fmt.Println(score)

	//slice ranges
	//arrays do not change but slices do change
	ranges := name[1:3]
	fmt.Println(ranges)
}
