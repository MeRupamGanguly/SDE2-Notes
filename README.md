# SDE2-Notes

## Introduce Yourself.
- I have a background in BTech IT, passout of 2020. After College I joined Sensibol as a GOlang Backed Developer. My Primary role involved Developing Microservices for Our Clients such as PDL and SingShala. We use Golang MongoDB AWS for Building Robust and Scalable Bussiness Logics. I Resigned in April as my Father Diagonises with CANCER, Now He is Well and Now I Seeking new Opportunities.

## What are Advantages and Disadvantages of Monolith and Microservices?
- Microservices are Better for BIG project Structure where Scalling and Almost ZERO Downtime required. Fixing Bug and Maintain Large Code base is easy with Microservices. Only Disadvantages of Microservice is Inter Service Network Call which can produce sometime Slow Response time.

## Describe GOlang Garbage Collection
- GOlang use Automatic Garabge Collection to Manage Memory. So developer do not need manual allocate and deallocate of Memory. So Coding is much more simple and a coder faced Less Memory realated error.

## What are the Difference between GoRoutine and Thread?
- Go Routine are designed for Concurency which Mean No task can Run Parallel, Multiple Tasks are run using Context Switching. Thread are designed for Parallelism which mean Tasks are executed simultaneously on Multiple CPU Cores.
- GO-Routines have Dynamicall Stack Size, Go Routines are Managed by Go-Runtime. Thread have Fixed Stack Size, Thread are Managed by OS-kernel.
- Go Routines communicate using Channels. Threads communicate using Shared-Memory. Communication using Channel is Safer.

## What is Lexical Scoping in Programming Language?
- Lexical Scoping is a Principle which determine Scope of Variables and Expressions. Scope is determined by where Variable is Declared or Defined within the Code rather than where it is Called or used during Execution.

## What is Closure in GOlang?
- A Closure is a Special type of Anonymous Fuction that can references variables declared outside of the Function itself. Closure treat Functions as Values - by which the Function can be assigned to Variables, passed as Arguments and Retured from another Function.

```go
func main(){
    /*
    x_name is a Function which can assigned to a V variable, x_name fuction return y_name function. 
    //c is defined at x_name block, y_name function can access/refrence c. 

    v:=func x_name(_no_arguments_) func y_name(_no_arguments_) (return_type int){
        c:=0
        return y_name(_no_arguments_) (return_type int){
            return_type = c+10
            return return_type
        }
    }  
    */
    v:=func () func () int{
        c:=0
        return func() int{
            c = c+10
            return c
        }
    } 
    i:=v()
    fmt.Println(i())
    fmt.Println(i())
    fmt.Println(i())
}
```

## Describe Functional Programming With GOLANG.
Personally In GOlang I did not use Functional Programming much, It takes lot of Memory as basically, we make copies of the data, we want to work on, and also Hard to Read for Freshers. But the two importants features I used sometimes are:

A function takes another Fuction as Arguments : Function as First Class Citigen
```go
func transformer(numbers []int,  trans func(int)int) (returnedValues []int){
    res:=make([]int, len(numbers))
    for i,v:=range numbers{
        res[i]=trans(v)
    }
    return res
}
```
In this Code Blocks transformer is a function which takes 2 arguments one is integer array another is a function which take int as arguments and returned int.
transformer function return an array of Int.
```go
func adder(x int) func(int)int{
    return func(y int) int{
        x+y
    }
}
```

```go
func main(){
    numbers:=[]int{2,4,8,16}
    doub:=transformer(numbers, func(x int)int{
        return x*2
    })
    trip:=transformer(numbers, func(x int)int{
        return x*3
    })
    add10:=adder(10)
    add10(2)
    add10(5)
    add12:=adder(12)
    add12(2)
    add12(5)
}
```
## Explain Exception Handeling with Panic Defer and Recover.
- Panic is use to cause a Runtime Error and Stop the execution.
- When a Function(van be Main Function) Returns or Panic the Defer block execute just before the return or panic by Last in First out manner. The last difer will execute fast.
- Recover is used to regain control of a Panicking situation. Recover used inside Defer block.

```go
func reco(){
    // use recover to check if a panic occurred.
    r:=recover()
    if r!=nil{ 
    // If it did, handle it appropriately (such as logging an error message)
    // and allow the program to continue executing.
     fmt.Println("Recovered from panic:", r)
}

func doSomething(i int)float{
    // Use defer to schedule reco() function immediately after doSomething function starts executing.
    defer reco() 
    if i==0{
        //  use panic to halt execution and potentially trigger reco().
        panic("Not Divisible by 0") 
    }else{
        return 27/i
    }
}
```

## Array vs Slice in Golang
- Array are Fixed in size and not shrink and grow on Runtime. Slice are not fixed in size and can grow and shrink on Runtime.
- In Golang slices are designed to act like dynamic arrays however they are just references to an existing array of a fixed length.

## Explain How Method Dispatching work on Golang.
- Method Dispatching means Selecting an implementation of a Method at Runtime. Go uses a mechanism called receiver type binding for method dispatching. In Golang we have Two way to Create a Function 
- Value Receiver func(svc serviceStruct) Method_Name(). Operates on a copy of the instance of the associated type. When you define a method with a value receiver, any modifications made to the receiver inside the method will not affect the original instance.
- Pointer Receiver func(svc *serviceStruct) Method_Name(). Operates directly on the instance of the associated type. When you define a method with a pointer receiver, any modifications made to the receiver inside the method will directly affect the original instance.
-  Pointer receivers directly operate on the instance, which is more memory-efficient and can result in faster execution, especially for large structs.

## Describe Concurency Primitives in GOlang.
Concurency Primitives are fundamental tools or mechanism provided by programming languages to help, manage and control the execution behaviours of concutent tasks.

Some common Concurency Primitives are Mutex, Semaphore, Channels etc.

- Mutex 
    - Mutex are used to protect shared Resources such as Variables and Structures, from being accessed simultaneously by Multiple threads or goroutines. 
    - Mutex in golang use Lock and Unlock function call to Lock and Unlock shared Resources to Protect access.
```go
// Here in this Example we try to Synchronize MAP 
var m map[string]int
var mu sync.Mutex 
// var mu sync.RWMutex : It allows multiple goroutines to read the map simultaneously but ensures exclusive access for writes.

// Writing to map
mu.Lock()
m["key"] = 123
mu.Unlock()

// Reading from map
mu.Lock() 
// mu.RLock() : Acquires a read lock. Multiple goroutines can hold a read lock simultaneously, allowing for concurrent read access to a shared resource.
value := m["key"]
mu.Unlock() // mu.RUnlock(

```

- Semaphore is an integer Counter.
    - Counter can be Decremented when a thrade wants to Enter a critical section or access a shared resources.
    -  If the semaphore's value is greater than zero (> 0), the semaphore decrements its value (semaphore_value--) and allows the thread to proceed. 
    - If the semaphore's value is zero (== 0), the thread is blocked (put into a waiting state) until the semaphore's value becomes positive.
    - And Increemnted when Thrade wants to Exit a critical section or completed access a shared resources. This operation increments the semaphore's value (semaphore_value++). 
    - If there are threads waiting (blocked on P), it unblocks one of them, allowing it to proceed.

- Channels 
    - Channels are Higher level concurency primitive, facilitate communication and synchronisation between Concurrent threads or goroutines by allowing them to Send and Receive values. 
    - Channels help prevent race conditions by ensuring that only one goroutine can access data at a time. 
    - Channels can be buffered, allowing goroutines to send multiple values without blocking until the buffer is full. This can improve performance in scenarios where the producer and consumer operate at different speeds. 

## How you use WaitGroup?
- wg.Add(1) increments the WaitGroup counter before each goroutine starts. 
- Each worker goroutine calls wg.Done() when it completes its task, which decrements the counter. 
- wg.Wait() blocks the main goroutine until the counter becomes zero, indicating that all workers have finished.
```go
func worker(i int, wg *sync.WaitGroup){
	defer wg.Done()
	time.Sleep(time.Second)
}
func main(){
	wg:=sync.WaitGroup{}
	for i:0;i<5;i++{
		wg.Add(1)
		go worker(i,&wg)
	}
	wg.Wait()
}
```
## Describe Map Synchronisation.
- In Go (Golang), maps are not inherently safe for concurrent access by multiple goroutines. If one goroutine is writing to a map while another goroutine is reading from or writing to the same map concurrently, it can result in unpredictable behavior or panics. They require synchronization mechanisms such as mutexes to ensure safe concurrent read and write operations.

![Screenshot from 2024-07-26 13-00-11](https://github.com/user-attachments/assets/f0286ec3-e3fd-4577-83bb-fdda90f1892c)
![Screenshot from 2024-07-26 13-03-33](https://github.com/user-attachments/assets/4399f402-d98f-4020-b2da-8cce57dae741)
![Screenshot from 2024-07-26 13-09-40](https://github.com/user-attachments/assets/5e3e7431-3e03-4351-9778-e299076dabe5)
![Screenshot from 2024-07-26 13-18-34](https://github.com/user-attachments/assets/36c2f9d3-c71e-420a-a15f-cdd3822a137e)

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	testMap := make(map[string]string)
	testMap["Jio"] = "Reliance"
	testMap["Airtel"] = "Bharti Airtel"

	for k, v := range testMap {
		fmt.Println("Key is: ", k, "Value is: ", v)
	}
	var wg sync.WaitGroup

	mu := sync.RWMutex{}

	for i := 0; i < 10; i++ {
		wg.Add(2)
		go mapWriter(&testMap, &wg, &mu)
		go mapConsumer(&testMap, &wg, &mu)
	}
	wg.Wait()
}

func mapConsumer(testMap *map[string]string, wg *sync.WaitGroup, mu *sync.RWMutex) {
	// Go does not allow you to use range directly on a pointer to a map because it expects a map type.
	// you need to dereference the pointer to get the actual map value, and then you can range over it.
	defer wg.Done()
	vTestMap := *testMap
	for k, v := range vTestMap {
		mu.RLock()
		fmt.Println("Consumer Key is : ", k, "Value is : ", v)
		mu.RUnlock()
	}
}
func mapWriter(testMap *map[string]string, wg *sync.WaitGroup, mu *sync.RWMutex) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		// Go does not allow you to use range directly on a pointer to a map because it expects a map type.
		// you need to dereference the pointer to get the actual map value, and then you can range over it.
		mu.Lock()
		vTestMap := *testMap
		vTestMap["Agni"] = fmt.Sprint("Jal", i)
		testMap = &vTestMap
		mu.Unlock()
	}
}
```