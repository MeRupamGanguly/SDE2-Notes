package main

import (
	"fmt"
	"sync"
)

func map_consumer(testMap map[string]string, wg *sync.WaitGroup, mu *sync.RWMutex) {
	for k, v := range testMap {
		mu.RLock()
		fmt.Println("--> ", k, v)
		mu.RUnlock()
	}
	wg.Done()
}
func map_writer(testMap map[string]string, wg *sync.WaitGroup, mu *sync.RWMutex) {
	for i := 0; i < 10; i++ {
		mu.Lock()
		testMap["Agni"] = fmt.Sprint("Jal", i)
		mu.Unlock()
	}
	wg.Done()
}
func main() {
	testMap := make(map[string]string)
	testMap["Jio"] = "Reliance"
	testMap["Airtel"] = "Bharti Airtel"
	for k, v := range testMap {
		fmt.Println("--> ", k, v)
	}
	fmt.Println("-------------------------------")
	var wg sync.WaitGroup
	var mu sync.RWMutex
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go map_writer(testMap, &wg, &mu)
		go map_consumer(testMap, &wg, &mu)
	}
	wg.Wait()
}
