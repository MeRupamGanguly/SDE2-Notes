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
    x_name is a Function which can assigned to a V variable, x_name fuction return y_name function. c is defined at x_name block, y_name function can access/refrence c. 

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
