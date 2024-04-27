package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 6}
	target := 5
	fmt.Println(searchInsert(arr, target))
}

func searchInsert(nums []int, target int) int {
	for i, j := range nums {
		if j == target {
			return i
		} else if j > target {
			return i
		}
	}
	return len(nums)
}
