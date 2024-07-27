package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(ch chan<- int, wg *sync.WaitGroup) {
	fmt.Println("Producer Call")
	for i := 0; i < 9; i++ {
		time.Sleep(time.Second)
		ch <- i
	}
	close(ch)
	wg.Done()
}
func consumer(ch <-chan int, wg *sync.WaitGroup) {
	fmt.Println("Consumer Call")
	for data := range ch {
		fmt.Println(data)
	}
	wg.Done()
}
func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go producer(ch, &wg)
	go consumer(ch, &wg)
	fmt.Println("Done")
	wg.Wait()
}
