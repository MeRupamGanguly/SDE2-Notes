package main

import "fmt"

func transformer(nums []int, trans func(int) int) (returned_values []int) {
	returned_values = make([]int, len(nums))
	for i, v := range nums {
		returned_values[i] = trans(v)
	}
	return
}
func adder(x int) func(int) int {
	return func(i int) int {
		return x + i
	}
}
func main() {
	numbers := []int{2, 4, 8, 16}
	doubls := transformer(numbers, func(i int) int { return i * 2 })
	fmt.Println(doubls)
	tripls := transformer(numbers, func(i int) int { return i * 3 })
	fmt.Println(tripls)
	add := adder(10)
	fmt.Println(add(2))
	fmt.Println(add(3))
	add = adder(12)
	fmt.Println(add(2))
	fmt.Println(add(3))
}
