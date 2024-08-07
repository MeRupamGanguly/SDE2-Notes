package main

import (
	"fmt"
	"sync"
	"time"
)

// Divide and Conquer: Parallel Task Aggregation
// The program divides a time-consuming summing task into smaller, concurrent tasks to improve efficiency and reduce total execution time.
// This function simulates a time-consuming task.
// task func([]int) int: A function that takes a slice of integers and
// returns an integer (e.g., computes the sum).
func slowTask(arr []int, task func([]int) int, ch chan<- int, wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	fmt.Println("Slow Task Called")
	defer wg.Done()
	res := task(arr)
	ch <- res // sends result to a channel.
}
func main() {
	startTime := time.Now()
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	spawnCapacity := 4 //Number of concurrent tasks (goroutines) to spawn.
	length := len(arr)
	patition := length / spawnCapacity  // Determines the size of each partition to be processed by each goroutine.
	ch := make(chan int, spawnCapacity) //  Creates a buffered channel with capacity spawnCapacity to collect results from goroutines.
	var wg sync.WaitGroup
	for i := 0; i < spawnCapacity; i++ {
		start := i * patition   // Determines the starting index of the current partition.
		end := start + patition // Determines the ending index of the current partition.
		if end > length {
			end = length // Adjusts the ending index to ensure it does not exceed the length of the array.
		}
		wg.Add(1)

		t := func(a []int) int { // Writing logic of the Task
			fmt.Println("logic called for slice: ", a)
			// Computes the sum of the slice and return.
			t := 0
			for i := range a {
				t += a[i]
			}
			return t
		}

		go slowTask(arr[start:end], t, ch, &wg)
	}
	go func() { // Starts a goroutine that waits for all tasks to complete and then closes the channel to signal that no more results will be sent.
		wg.Wait()
		close(ch)
	}()
	//Aggregates the results from the channel and computes the total sum.
	sum := 0
	for partialTaskResult := range ch {
		sum += partialTaskResult
	}
	fmt.Println("Sum is: ", sum)
	fmt.Println("Time Taken: ", time.Since(startTime))
}

/*
Slow Task Called
Slow Task Called
logic called for slice:  [16 17 18 19 20]
logic called for slice:  [6 7 8 9 10]
Slow Task Called
logic called for slice:  [1 2 3 4 5]
Slow Task Called
logic called for slice:  [11 12 13 14 15]
Sum is:  210
Time Taken:  1.000637245s
*/
/*
Purpose: The program divides a time-consuming summing task into smaller, concurrent tasks to improve efficiency and reduce total execution time.

Task Division:

    The array is partitioned into smaller segments based on a specified number of concurrent tasks (spawnCapacity).
    Each segment is processed by a separate goroutine to compute the sum of its elements.

Concurrency Management:

    Goroutines: Used to execute the summing tasks concurrently. Each goroutine handles a segment of the array and computes its sum.
    Channels: Employed for communication between goroutines. Results are sent through a channel to be aggregated later.
    WaitGroup: Manages synchronization by keeping track of the number of active goroutines. It ensures that the main program waits for all goroutines to finish before closing the channel.

Execution Flow:

    Start Time: The execution time is measured from the start of the program to calculate performance.
    Task Execution: Each goroutine processes its assigned segment, computes the sum, and sends the result to a channel.
    Aggregation: Once all goroutines complete their tasks, the main program aggregates the results from the channel to compute the total sum.
    Completion Time: The total time taken for the entire process is printed at the end.

Efficiency:

    Parallel Processing: By executing tasks concurrently, the program typically reduces the overall processing time compared to a sequential approach where tasks are executed one after another.
    Resource Utilization: Concurrent execution can make better use of system resources, potentially leading to improved performance for time-consuming tasks.

*/
