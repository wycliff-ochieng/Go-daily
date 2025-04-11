package main

import "fmt"

func main() {
	foods := []string{"mango", "fish", "soda", "ugali", "Matumbo"}
	for _, food := range foods {
		switch food {
		case "mango":
			fmt.Println(food, ": A fruit")
		case "fish":
			fmt.Println(food, ": Protein")
		case "soda":
			fmt.Println(food, ": Soft drink")
		case "ugali":
			fmt.Println(food, ": Carbos")
			break
		default:
			fmt.Println("just eat")
		}

	}
	fmt.Println("")
}
