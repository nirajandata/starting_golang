package main

import (
	"fmt"
)

func sum(nums ...int) {
	fmt.Println("nums are ", nums)
	total := 0
	for _, num := range nums {
		fmt.Println(num)
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2, 3)
	sum(1, 2, 23, 3, 3, 3, 3, 3, 3)
	nums := []int{1, 2, 2, 2, 2, 4555}
	sum(nums...)
}
