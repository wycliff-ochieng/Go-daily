package main

import (
	"errors"
	"fmt"
)

func main() {
	/*fmt.Println("Error Handling")
	numerator := 20
	denominator := 0
	remainder, mod, err := calcRemainderAndMod(numerator, denominator)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(remainder, mod)
	/*
		if err := userJustShit(); err != nil {
			fmt.Println(changePants())
		}*/
	a := 2
	b := 2
	err := returnError(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("returnError() function ended normally")
	}
	first := "Wyclif"
	last := "Ochieng"
	fmt.Println(fullName(first, last))
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

func calcRemainderAndMod(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("denominator cannot be zero")
	}
	return numerator / denominator, numerator % denominator, nil
}

func returnError(a int, b int) error {
	if a == b {
		err := errors.New("Error in returnError() function")
		return err
	} else {
		return nil
	}
}

func fullName(first string, last string) (string, error) {
	fullname := first + " " + last
	if len(fullname) > 4 {
		err := errors.New("Invalid fullname")
		return fullname, err
	} else {
		return fullname, nil
	}
}
