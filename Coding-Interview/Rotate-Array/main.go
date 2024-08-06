package main

import "fmt"

/* Problem Statement:
[2,4,8,16,32,64]
rotate right 3 step
Output: [16,32,64,2,4,8]
*/

// Define the Contract
type Rotation interface {
	rotate(arr []int, step int) (res []int)
}

// Define the class
type service struct{}

// Implementing the Rotation interface by receiver function
func (svc *service) rotate(arr []int, step int) (res []int) {
	length := len(arr)
	res = append(res, arr[length-step:]...) // apending all the elements after step
	res = append(res, arr[:length-step]...) // appending all the elements before the step
	return
}
func main() {
	arr := []int{2, 4, 8, 16, 32, 64}
	step := 3
	s := service{} // create object of class
	fmt.Println(s.rotate(arr, step))
}
