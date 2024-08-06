package main

import "fmt"

func sumOfNNumbers(arr []int) int {
	length := len(arr)
	sum := (length * (length + 1)) / 2
	return sum
}
func searchMissingNumber(arr []int) int {
	expectedSum := sumOfNNumbers(arr)
	actualSum := 0
	for i := range arr {
		actualSum += arr[i]
	}
	return expectedSum - actualSum
}
func main() {
	arr := []int{0, 1, 2, 3, 4, 6, 7, 8}
	fmt.Println(searchMissingNumber(arr))
}
