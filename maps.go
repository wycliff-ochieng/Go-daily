package main

import "fmt"

func main() {
	//var m map[key_type]value_type
	//make - m := make(map[key_type]value_type)

	teams := map[string][]string{
		"England": []string{"Arsenal", "Manchester united", "WestHam"},
		"Spain":   []string{"Real Madrid", "Barcelona", "Atletico Madrid"},
		"Germany": []string{"Bayern Munich", "Dortmund", "Leverkusen"},
	}
	fmt.Println(teams)

	//comma ,ok idiom
	//using maps as sets

	intSet := map[int]bool{}
	vals := []int{3, 9, 5, 8, 5, 1, 9, 4, 5, 3, 6, 7, 8, 4, 3}

	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
}
