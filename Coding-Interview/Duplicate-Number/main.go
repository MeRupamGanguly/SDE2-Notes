package main

import "fmt"

func searchFirstDuplicate(arr []int) int {
	unique := make(map[int]bool) //contains Unique elemets
	for _, a := range arr {
		// If the element present in the map then it is Duplicate
		if _, exist := unique[a]; exist {
			return a
		} else {
			unique[a] = true
		}
	}
	return 0
}
func main() {
	arr := []int{2, 4, 6, 8, 5, 6, 2, 1, 8, 5}
	fmt.Println(searchFirstDuplicate(arr))
}
