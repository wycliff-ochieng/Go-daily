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
	s_range := score[1:3]
	fmt.Println(s_range)
	fmt.Println(ranges)

	//slices arent comparable

	var x []int = []int{3, 4, 5, 7, 4, 9, 70, 23, 45}
	var y []int
	x[0] = 10
	fmt.Println(x[2])
	fmt.Println(y)
	fmt.Println(y == nil)
	y = append(y, 10)
	fmt.Println(y)

	var country []string
	fmt.Println(country, len(country))
	//fmt.Println(len(country))
	country = append(country, "Kenya", "Uganda", "South Africa")
	fmt.Println(country, len(country))
	//appending slices to another (...)
	county := []string{}
	county = append(county, "Nairobi", "Kitui")
	fmt.Println(county)
	central := []string{}
	central = append(central, "Kiambu", "Muranga")
	fmt.Println(central, len(central))
	central = append(central, county...)
	fmt.Println(central, ":", len(central))
}
