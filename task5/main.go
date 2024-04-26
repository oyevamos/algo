package main

import "fmt"

func main() {
	arr := []int{3, 2, 4}
	tar := 6
	fmt.Println(twoSum(arr, tar))
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
