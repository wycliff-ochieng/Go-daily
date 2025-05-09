package main

import (
	"errors"
	"fmt"
)

func main() {
	n := -2
	numerator := 10
	denominator := 5
	fmt.Println(division(numerator, denominator))
	fmt.Println(negativeCheck(n))
}

func negativeCheck(n int) error {
	if n < 0 {
		return errors.New("n cannot be a negative")
	}
	return nil
}

func division(numerator int, denominator int) (int, error) {
	if numerator == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return numerator / denominator, nil
}
 type CustomErr struct{
	Status Status
	message string
 }

 func pageErr(name string, passcode string) (string , error){
	success, err := pageErr(name,passcode)
	if err != nil{
		return nil, CustomErr{
			Status : Invalidlogin,
			message : fmt.Println("Try again later")
		}
	}
	return success,nil
 }

 func origNo(num *int) int{
	*num = 5
 }
 func hikeNo(num int) int{
	num = 6
 }