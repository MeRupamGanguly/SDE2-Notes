package main

import (
	"fmt"
	"time"
)

func slowTask(arr []int) int {
	fmt.Println("Slow Task Called")
	time.Sleep(time.Second)
	sum := 0
	for i := range arr {
		sum += arr[i]
	}
	return sum
}
func main() {
	startTime := time.Now()
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	spawnCapacity := 4
	length := len(arr)
	patition := length / spawnCapacity
	sum := 0
	for i := 0; i < 4; i++ { // The array is divided into partitions, with each partition being processed by the slowTask function.
		start := i * patition
		end := start + patition
		if end > length {
			end = length
		}
		sum += slowTask(arr[start:end])
	}
	fmt.Println("Sum is: ", sum)
	fmt.Println("Time Taken: ", time.Since(startTime))
}

/*
Slow Task Called
Slow Task Called
Slow Task Called
Slow Task Called
Sum is:  210
Time Taken:  4.002764194s
*/
