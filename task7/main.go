package main

import "fmt"

func main() {
	arr := []int{3, 2, 2, 3}
	val := 3
	fmt.Println(removeElement(arr, val))
}

func removeElement(nums []int, val int) int {
	i := 0
	for _, num := range nums {
		if num != val {
			nums[i] = num
			i++
		}
	}
	return i
}
