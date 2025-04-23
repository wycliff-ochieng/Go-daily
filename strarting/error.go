package main

import (
	"fmt"
	"errors"
)

func main() {
	fmt.Println("Error Handling")
	numerator := 20
	denominator := 3
	remainder, mod, err := calcRemainderAndMod(numerator,denominator)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(remainder,mod)
	if err := userJustShit(); err != nil {
		fmt.Println(changePants())
	}
}

func userJustFart() (int, error) {
	return 10, nil
}

func userJustShit() error {
	return nil
}
func changePants(string) {
	fmt.Println("Pants have been changed.... no smell anymore")
}

func calcRemainderAndMod(numerator int, denominator int) (int,int, error){
	if denominator == 0{
		return 0,0,errors.New("denominator cannot be zero")
	}
	return numerator / denominator, numerator % denominator, nil
}
