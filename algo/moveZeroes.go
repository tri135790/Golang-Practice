package main

import "fmt"

func main() {
	var arr = []int{0, 1, 0, 3, 0, 0, 12, 0}
	moveZeroes(arr)
	fmt.Println(arr)

}

func moveZeroes(nums []int) {
	var countZeroes int = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i-countZeroes] = nums[i]
			if countZeroes != 0 {
				nums[i] = 0
			}
		} else {
			countZeroes++
		}
	}
}
