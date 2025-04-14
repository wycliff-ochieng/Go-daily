package main

import "fmt"

func main() {
	fmt.Println("answer")
	nums := []int{2, 7, 4, 5, 7}
	target := 9
	ans := twoSum(nums, target)
	fmt.Println(ans)
}
func twoSum(nums []int, target int) []int {
	num2index := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, ok := num2index[complement]; ok {
			return []int{j, i}
		}
		num2index[num] = i
	}
	return []int{}
}
