package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Roo")
	ans := div(20, 10)
	fmt.Println(ans)
	fmt.Println(addToBase(3))
	fmt.Println(addToBase(3, 4))
	fmt.Println(addToBase(3, 4, 8, 9))
	fmt.Println(addToBase(3, 4, 9, 8))
	//fmt.Println(div2(20, 0))
	result, remainder, err := div2(5, 3)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)
}

func div(numerator int, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

// variadic inputs
func addToBase(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

// multiple return values
func div2(nume int, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("Cannot divide by 0")
	}
	return nume / denom, nume % denom, nil
}
