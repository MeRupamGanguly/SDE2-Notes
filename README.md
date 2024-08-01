# SDE2-Notes

## Introduce Yourself.
- I have a background in BTech IT, passout of 2020. After College I joined Sensibol as a GOlang Backed Developer. My Primary role was Developing Microservices for Our Clients such as PDL and SingShala. We use Golang MongoDB AWS for Building Robust and Scalable Bussiness Logics. I Resigned in April as my Father Treatments of CANCER, Now He is Well and Now I Seeking new Opportunities.

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

- When Communication and Synchronisation between Go-Routines are Priority then Choose Channel. When Protecting Shared-Resources is Priority then Choose Mutex. When Limiting access to a Pool of Resources(Comapny have only 3 Printers for 25 Computers) then Choose Semaphore.

Some common Concurency Primitives are Mutex, Semaphore, Channels, Atomic etc.

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

- Atomic 
    - Atomic Operations are used on Primitive Data Types like int32,float64, Uint64 etc. Mainly use on Incrementing/Updates Counters/Flags. Atomic Opertaions are Lock free and Efficient for Primitives Data-types only.

- Condition Variables
    - Condition variables allow threads to wait until a certain condition is true before proceeding. They are typically used in conjunction with mutexes to manage complex synchronization patterns.

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
		mu.Lock()
		vTestMap := *testMap
		vTestMap["Agni"] = fmt.Sprint("Jal", i)
		testMap = &vTestMap
		mu.Unlock()
	}
}
```

## Describe Channel Comunication
- `<-` this syntax is used for Sends Value into Channel.
- `:=<-`  this syntax is used for Receives Value from Channel.
```go
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
	close(ch) // If we do not close the channel then Deadlock occured as Reader continiously try to read from channel until closed.
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
```

## Describe uses of Select in Golang.
Assume a development scenerio where we have 3 s3 Buckets. We spawn 3 GO-Routines each one uploading a File on each S3 bucket at same time. We have to Return SignedUrl of the file so user can stream the File as soon as possible. Now we do not have to wait for 3 S3 Upload operation, when one s3 upload done we can send the SignedUrl of the File to the User so he can Stream. And Other two S3 Upload will continue at same time. This is the Scenerio when Select Statement will work as a Charm.

Select Block Main Execution, untils one of its case completed. If multiple case completes at same time then Randomly one case is seclected for Output and Select block will complete.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		rand.NewSource(time.Now().Unix())
		i := rand.Intn(6)
		fmt.Println("ch1 Rand ", i)
		time.Sleep(time.Duration(i) * time.Second)
		ch1 <- 1
	}()
	go func() {
		rand.NewSource(time.Now().Unix())
		i := rand.Intn(6)
		fmt.Println("ch2 Rand ", i)
		time.Sleep(time.Second * time.Duration(i))
		ch2 <- 2
	}()

	select {
	case <-ch1:
		fmt.Print("ch1 is Done")
	case <-ch2:
		fmt.Print("ch2 is Done")
	case <-time.After(time.Second * time.Duration(4)):
		fmt.Print("Context Expire")
	}
}
```

```go
package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func coreProcess(ctx context.Context, i int, ch chan<- int) {
	rand.NewSource(time.Now().Unix())
	r := rand.Intn(4)
	t := time.Duration(r) * time.Second
	ctx, cancel := context.WithTimeout(ctx, t)
	defer cancel()
	fmt.Printf("Doing Some Work for Process: %d with random %d", i, r)
	fmt.Println()
	select {
	case <-time.After(t): //Send 1 via Channel to Caller Function After Wait for Random Seconds.
		ch <- 1
	case <-ctx.Done(): // If context expire then send 0 via Channel to Caller Function
		ch <- 0
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := make(chan int)
	defer close(ch)
	for i := 0; i < 6; i++ { // Calling CoreProcess concurently 6 times.
		go coreProcess(ctx, i, ch)
	}
	// gather result from al Concurent Processes.
	for i := 0; i < 6; i++ {
		select {
		case res := <-ch: // To store it a res is important , otherwise in Print statement channel read twice
			fmt.Println(res) // fmt.Println(<-ch)  X:Wrong Syntax
		case <-ctx.Done():
			fmt.Println("Context Done")
		}
	}
}
```

## What are SOLID Principles.
- SOLID principles are guidelines for Designing Code-base that are easy to understand, maintain and extend over time.

Main Points to Remeber are: 
1. Struct have only those fields which are strictly copuled, decoupled fields can be on another strictly copuled Struct. 
2. New features can be added without changing old features/Functions. 
3. Obeject of Superclass can be replacable by its subclass instances, without affecting corectness of the features.
4. Strcut should only implements what it really needs, instead of unrequired methods, so do not club different decouplable interfaces. 
5. Classes rely on interfaces instead of specific Implementations, that way we can achive tech-stack-migrations easily.

- Single Responsibility: A Struct/Class should have only a single reason to change. Fields of Book and Fields of Author should be on Different struct. 

```go
type Book struct{
  ISIN string
  Name String
  AuthorID string
}
type Author struct{
  ID string
  Name String
}
```
Assume One Author decided later, he does not want to Disclose its Real Name to Spread. So we can Serve Frontend by Alias instead of Real Name. Without Changing Book Class/Struct, we can add Alias in Author Struct. By that, Existing Authors present in DB will not be affected as Frontend will Change Name only when it Founds that Alias field is not empty.
```go
type Book struct{
  ISIN string
  Name String
  AuthorID string
}
type Author struct{
  ID string
  Name String
  Alias String
}
```
- Open/Closed: Coding Components like Struct, Functions etc should be Open for Extension but Closed for Modification. This principle encourages us to design our systems in a way that allows new functionality to be added without changing existing code.

```go
type Shape interface{
	Area() float64
}
type Rectangle struct{
	W float64
	H float64
}
type Circle struct{
	R float64
}
```
Now we want to Calculate Area of Rectangle and Circle, so Rectangle and Circle both can Implements Shape Interface by Write Body of the Area() Function.
```go
func (r Rectangle) Area()float64{
	return r.W * r.H
}
func (c Circle)Area()float64{
	return 3.14 * c.R * c.R
}
```
Now we can create a Function PrintArea() which take Shape as Arguments and Calculate Area of that Shape. So here Shape can be Rectangle, Circle. In Future we can add Triangle Struct which implements Shape interface by writing Body of Area. Now Traingle can be passed to PrintArea() with out modifing the PrintArea() Function.
```go
func PrintArea(shape Shape) {
	fmt.Printf("Area of the shape: %f\n", shape.Area())
}
```
```go
// In Future
type Triangle struct{
	B float64
	H float54
}
func (t Triangle)Area()float64{
	return 1/2 * t.B * t.H
}
```
```go
func main(){
	rect:= Rectangle{W:5,H:3}
	cir:=Circle{R:3}
	PrintArea(rect)
	PrintArea(cir)
	// In Future
	tri:=Triangle{B:4,H:8}
	PrintArea(tri)
}
```
- Liskov Substitution: Objects of a Super Class should be Replacable with Objects of its Sub Classes without affecting the correctness of the Program.
```go
type Bird interface{
	Fly() string
}
type Sparrow struct{
	Name string
}
type Penguin struct{
	Name string
}
```
Sparrow and Pengin both are Bird, But Sparrow can Fly, Penguin Not. ShowFly() function take argument of Bird type and call Fly() function. Now as Penguin and Sparrow both are types of Bird, they should be passed as Bird within ShowFly() function.
```go
func (s Sparrow) Fly() string{
	return "Sparrow is Flying"
}
func (p Penguin) Fly() string{
	return "Penguin Can Not Fly"
}
```
```go
func ShowFly(b Bird){
	fmt.Println(b.Fly())
}
func main() {
	sparrow := Sparrow{Name: "Sparrow"}
	penguin := Penguin{Name: "Penguin"}
  // SuperClass is Bird,  Sparrow, Penguin are the SubClass
	ShowFly(sparrow)
	ShowFly(penguin)
}
```
- Interface Segregation: Class should not be forced to Depends on Interfaces they do not want to use or Implements.  If we use a normal Printer machine, which Need only Print() function then that Struct/Class only implements Printer interface instead of NewTypeOfDevice interface.

```go
// The Printer interface defines a contract for printers with a Print method.
type Printer interface {
	Print()
}
// The Scanner interface defines a contract for scanners with a Scan method.
type Scanner interface {
	Scan()
}
// The NewTypeOfDevice interface combines Printer and Scanner interfaces for New type of devices which can Print and Scan with it new invented Hardware.
type NewTypeOfDevice interface {
	Printer
	Scanner
}
```
- Dependency Inversion: Classes depends on Interfaces not Implementations. Thats help Decoupling. That means if we change Our Primary Dtabase SQL to MongoDB then our Service layer should not be touched, only Repository layer Changes by Implementing Repository Interfaces, and Main file change where we create Mongodb Connection.
```go
// The MessageSender interface defines a contract for sending messages with a SendMessage method.
type MessageSender interface {
	SendMessage(msg string) error
}
// EmailSender and SMSClient structs implement the MessageSender interface with their respective SendMessage methods.
type EmailSender struct{}

func (es EmailSender) SendMessage(msg string) error {
	fmt.Println("Sending email:", msg)
	return nil
}
type SMSClient struct{}

func (sc SMSClient) SendMessage(msg string) error {
	fmt.Println("Sending SMS:", msg)
	return nil
}
type NotificationService struct {
	Sender MessageSender
}
```
The NotificationService struct depends on MessageSender interface, not on concrete implementations (EmailSender or SMSClient). This adheres to Dependency Inversion, because high-level modules (NotificationService) depend on abstractions (MessageSender) rather than details.
```go

func (ns NotificationService) SendNotification(msg string) error {
	return ns.Sender.SendMessage(msg)
}
func main() {
	emailSender := EmailSender{}
	smsClient := SMSClient{}

	emailNotification := NotificationService{Sender: emailSender}
	smsNotification := NotificationService{Sender: smsClient}

	emailNotification.SendNotification("Hello, this is an email notification!")
	smsNotification.SendNotification("Hello, this is an SMS notification!")
}
```
## What are the different Design Patterns you know and Explain each with GOLANG.
- Sigleton: Singleton pattern ensures a Class/Struct has only one Instance. And provide a global point of access to that instance. That instance should be designed to be Thread-safe so Multiple go-routines can access it concurrently without create multiple instances of the Class. Global Configs, Database Connection, Logging service, are where we can use Singleton approach.
```go
type Config struct{
	configs map[string] string
}
var (
	confs *Config
	once sync.Once
)
func (c *Config) initConfig() {
    // Load configuration settings from file, database, etc.
    c.configs["server_address"] = "localhost"
    c.configs["port"] = "8080"
    // Add more configuration settings as needed
}

// GetConfigManager returns the singleton instance of ConfigManager.
// sync.Once ensures that, the initialization code, inside once.Do() is executed exactly once, 
//  preventing multiple initializations even with concurrent calls.
func GetConfigs() *Config{
	once.Do(func(){obj=&Config{configs:make(map[string]string)}})
	obj.initConfig()
}

// GetConfig retrieves a specific configuration setting.
func (cm *ConfigManager) GetConfig(key string) string {
    return cm.config[key]
}

func main() {
    // Get the singleton instance of ConfigManager
    configManager := GetConfigManager()

    // Access configuration settings
    fmt.Println("Server Address:", configManager.GetConfig("server_address"))
    fmt.Println("Port:", configManager.GetConfig("port"))
}
```
- Builder: Construct/Builds Complex objects Step by Step.
```go
type Product struct{ // The Complex Object we want to Build and give back as a Response to our Primary API.
	Part_1 interface{}
	part_2 interface{}
	part_3 string
	part_4 bool
}
```
```go
type Builder interface{ // The Contracts use for building the Complex Product.
	set_part_1(uddt type interface{}) (resp interface{}, err error)
	set_part_2(uddt type interface{}) (resp interface{}, err error)
	set_part_3(uddt type interface{}) (resp string, err error)
	set_part_4(uddt type interface{}) (resp bool, err error)
	Build()Product
}
// the class which gonna implements Contracts
type service struct{  } 

func (svc *service) set_part_1(uddt type interface{}) (resp interface{}, err error) {
	fmt.Println("part_1 computation done")
	var ifs interface{}
	ifs = "Verify Complete"
	return ifs,nil
}
func (svc *service) set_part_2(uddt type interface{}) ( interface{},  error) {
	fmt.Println("part_2 computation done")
	var ifs interface{}
	ifs = "Downloading Complete"
	return ifs,nil
}
func (svc *service) set_part_3(uddt type interface{}) (resp string, err error) {
	fmt.Println("part_3 computation done")
	return "Converting completed",nil
}
func (svc *service) set_part_4(uddt type interface{}) (resp bool, err error)  {
	fmt.Println("part_4 computation done")
	return  true,nil
}
func (svc *service)Build()Product{

	var ifs_part_1 interface{}
	ifs_part_1="Input for Part_1"
	resp_part_1,_:=svc.set_part_1(ifs_part_1)

	var ifs_part_2 interface{}
	ifs_part_2="Input for Part_2"
	resp_part_2,_:=svc.set_part_2(ifs_part_2)

	var ifs_part_3 interface{}
	ifs_part_3="Input for Part_3"
	resp_part_3,_:=svc.set_part_3(ifs_part_3)

	var ifs_part_4 interface{}
	ifs_part_4="Input for Part_4"
	resp_part_4,_:=svc.set_part_1(ifs_part_4)

	return Product{
		Part_1:resp_part_1,
		Part_2:resp_part_2,
		Part_3:resp_part_3,
		Part_4:resp_part_4,
	}
}
```
```go
type Director struct{
	builder Builder
}
func NewDirector(builder Builder) *Director{
	return &Director{builder: builder}
}
func (d *Director) Run() Product{
	return d.builder.Build()
}
```
```go
func main(){
	builder_obj := &service{}
	director_obj:= NewDirector(builder_obj)
	out_product:= director_obj.Run()
	fmt.Println(out_product)
}
```
- Factory: In this Design patterns we create objects in a way that allow Flexibility to choose. Exact type of Object only known at Runtime. 

```go
type Repository interface{
	Get(id int)(interface{},error)
	Add(data interface{}) error
}
type MySqlImp struct{// MySQL connection details or any necessary configuration}

func(sq *MySqlImp)Get(id int)(interface{},error){
	// Implement MySQL specific logic to fetch data by ID
	return
}
func(sq *MySqlImp)Add(data interface{}) error{
	// Implement MySQL specific logic to Add data
	return
}

type MongoDBImp struct{// MongoDB connection details or any necessary configuration}

func(mo *MongoDBImp)Get(id int)(interface{},error){
	// Implement MongoDB specific logic to fetch data by ID
	return
}
func(mo *MongoDBImp)Add(data interface{}) error{
	// Implement MongoDB specific logic to Add data
	return
}
```
```go
type repo struct{}
func(r *repo) createRepo(type int) (Repository){
	switch type{
		case 0:
			return &MySqlImp{}
		case 1:
			return &MongoDBImp{}
	}
}
```
```go
func main(){ // Client
	re:=repo{}
	riposi:=re.createRepo(1)// at this time only we know Waht type of Get() method is called 
	dtata:=riposi.Get(56)
}
```
- Observer: In this design pattern chnages of a Feature will broadcast to its Subscriber.
Subscriber Observe to Data Change of the Producer.
```go
type Subscriber interface{ 
	Update(tick float64)
}

type Topic interface{
	Register(s Subscriber)
	Deregistrer(s Subscriber)
	Notify(tick float64)
}

type Publisher struct{ // Publisher Will implenets Topic interface
	subs []Subscriber
} 
func (p *Publisher) Register(s Subscriber){
	p.subs=append(p.subs, s)
}
func (p *Publisher) Deregister(s Subscriber){
	for i,sub:=range p.subs{
		if sub==s{
			p.subs=append(p.subs[:i],p.subs[i+1:]...)
			break
		}
	}
}
func (p *Publisher) Notify(tick float64){ 
	// Notify will get all Active Subscriber and Updates the tick of subscribers
	for _,sub:=range p.subs{
		sub.Update(tick)
	}
}
```
```go
type Client struct{ // Client Will implements Subscriber interface
	Name string
}
func (c *Client) Update(tick float64){
	fmt.Println(c.Name+" is listening. ping is :",tick)
}
```
```go
func main(){

	sub_1:=&Client{Name:"R"}
	sub_2:=&Client{Name:"u"}
	sub_3:=&Client{Name:"p"}
	sub_4:=&Client{Name:"a"}
	sub_5:=&Client{Name:"m"}

	pub:=&Publisher{}

	pub.Register(sub_1)
	pub.Register(sub_2)
	pub.Register(sub_3)
	pub.Register(sub_4)
	pub.Register(sub_5)

	go func(){ // Publisher will publish tick every seconds
		for i:=0;i<2;i++{
			tick:=float64(i)+0.5
			pub.Notify(tick)
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(6*time.Second)
	pub.Deregister(sub_3)
	// Simulate publishing another article after deregistration
	pub.Notify(23.67)
	time.Sleep(5*time.Second)
}
```
- Decorator: Most important patterns For Game Developemnts. This patterns allows add features Dynamically at Runtime without affecting the Core functionalities. 
```go

type Features interface {
	Attack() int
	Defense() int
}

type Player struct{ // Player will implements Features interface.
	attack int
	defense int
}
func (p *Player) Attack()int{
	return p.attack
}
func (p *Player) Defense() int{
	return p.defense
}
```
```go
type Weapon struct{ // Weapon enhances a player's attack power by adding more value.
	player *Player
	attack int
} 
type Helmet strcut{ // Helmet enhances a player's defense power by adding more value.
	player *Player
	defense int
}
func (p *Weapon) Attack()int{
	p.player.attack= p.player.attack + p.attack
	return p.player.attack 
}
func (p *Weapon) Defense() int{
	return p.player.defense
}
func (p *Helmet) Attack()int{
	return p.player.attack 
}
func (p *Helmet) Defense() int{
	p.player.defense=p.player.defense+p.defense
	return p.player.defense
}
```
```go
func main(){
	// Spawn with default Powers
	p:=&Player{
		attack:10,
		defense:10,
	}
	fmt.Println("Game Start")
	fmt.Println("Player Found Level3-Helmet")
	player_with_helemt:=&Helmet{
		player: p,
		defense: 5,
	}
	fmt.Println("Player has Attack ",player_with_helemt.Attack()," Defense ",player_with_helemt.Defense())
	fmt.Println("After Some time player found AK-47")
	player_with_helemt_weapon:=&Weapon{
		player: &player_with_helemt.player,
		attack: 5,
	}
	
	fmt.Println("Player has Attack ",player_with_helemt_weapon.Attack()," Defense ",player_with_helemt_weapon.Defense())
}	
```
```bash
Game Start
Player Found Level3-Helmet
Player has Attack  10  Defense  15
After Some time player found AK-47
Player has Attack  15  Defense  15
```

## Explain Domain Driven Development.
- In Domain Driven Development we focus on deep understanding of the Core problem we are try to Solve. We focus on understanding the Entities, Value-Objects, Aggregates that reflects the Business Logic. We start development of the Core then Iteratively development based on Feedback and Evolving understanding. We breakdown the systems into manageable parts(Micro-services) based on different business logic.

Imagine we are building Online Shopping Cart, where User can Browse and Add product into Cart, then can place Orders. So,

we try to understand deeply the Domain of Online Cart:

- Core Concepts: Products, Users, Orders, Cart, Payment, Shipping etc.

- Business Rules: Discount, Shipping-Polices, Payment-Gateways, Inventory-Management etc.

- Explain all the Terms to the Team Members. Define each term means and how they interacts.

- Identify the Boundaries: Creating Different Micro-services with separate abilities (have to create Product, Order, Auth, User, Payment etc Micro-services). Product and Order should not be mixed.

- Clear separation of Domain Models: Product, Order, User Models.

- Clear knowledge of Value-Object: Value object is a struct which have no identity without its own Fields value. Two Value object will be treated as equal if both have same values of all the fields. CartItem is a Struct which have Product_id and Quantity fields, These attributes determine the nature of the Cart-item. In domain-driven design, equality of value objects is determined based on their attribute values rather than their identity.

- Clear understanding of Aggregates: Aggregates are the Cluster of Related Objects that are treated as a Single unit. An Order often comprises multiple related entities and value objects such as Order-items, Customer-details, and Payment-details. when an order is placed or updated, all related data within the aggregate is managed consistently. Order struct acts as the Root of the Aggregate, managing the life-cycle of associated Objects like Order-items and PaymantDetails.

- Clear understanding of Services: PaymentService (handling payment transactions), Shipping-service (calculating shipping costs).

- Create proper Repositories Contracts using Interfaces. And Separate them based on the Services.

- Create Some important Events like Order-placed event, PayamentSuccess event, Order-delivered event, Inventory-updated event etc. Using these events we can trigger other services like Advertisement, Notify, Subscriptions, Report/Analytics etc.

- Clear understanding of Orchestration: Placed-Order-Service Coordinates with Inventory and Payments services. Cart-Service Coordinates with Pricing of Product and Inventory for adaptation of future price and quantity changes.

### Define Entities and Value Objects

```go


// Product entity

type Product struct {
	ID uuid.UUID
	Name string
	Price float64
	Quantity int
}

/*
CartItem value object
A CartItem is defined by its attributes such as ProductID and Quantity.
These attributes determine the nature of the CartItem,
but the CartItem itself does not have an identity independent of these attributes.
Two CartItem instances are considered equal if their attributes
(ProductID and Quantity) are the same. In domain-driven design,
equality of value objects is determined based on their attribute
values rather than their identity.
*/

type CartItem struct {
	ProductID uuid.UUID
	Quantity int
}

// Order entity

type Order struct {
	ID uuid.UUID
	CustomerID uuid.UUID
	Items []CartItem
	Total float64
	CreatedAt time.Time
}


```

 ### Aggregates group related entities and value objects together, with an aggregate root enforcing consistency.

```go
/*
Order aggregate root
An Order often comprises multiple related entities
and value objects such as OrderItems, CustomerDetails, and PaymentDetails
when an order is placed or updated, all related data
within the aggregate is managed consistently.
An order may have rules regarding minimum quantities,
total prices, or eligibility for discounts.
By defining these rules within the aggregate,
you ensure they are consistently enforced whenever an order is manipulated.
The CustomerDetails and PaymentDetails are correctly associated with the order.
*/
type OrderAggregate struct {
	Order *Order
}

func (oa *OrderAggregate) AddItem(productID uuid.UUID, quantity int) error {
	// Business logic to add items to the order
}


```

 ### Repositories abstract the data persistence layer for entities.

```go

// ProductRepository interface
type ProductRepository interface {
	Save(product *Product) error
	FindByID(id uuid.UUID) (*Product, error)
	// Other methods like FindByName, Delete, etc.

}

// OrderRepository interface
type OrderRepository interface {
	Save(order *Order) error
	FindByID(id uuid.UUID) (*Order, error)
	// Other methods like FindByCustomerID, UpdateStatus, etc.
}

```

 ### Application services orchestrate interactions between domain models and external systems.

```go

// ProductService handles operations related to products
type ProductService struct {
	productRepo ProductRepository
}

func (ps ProductService) AddProduct(name string, price float64, quantity int) (Product, error) {
	// Business logic to add a new product
}

// OrderService handles operations related to orders
type OrderService struct {
	orderRepo OrderRepository
}

func (os OrderService) PlaceOrder(customerID uuid.UUID, items []CartItem) (Order, error) {
	// Business logic to place a new order
}


```

### Implement the infrastructure layer for persistence, messaging, etc. For simplicity, we'll use in-memory repositories.

```go
// InMemoryProductRepository implements ProductRepository using in-memory storage

type InMemoryProductRepository struct {
	products map[uuid.UUID]*Product
}

func (r InMemoryProductRepository) Save(product Product) error {
	r.products[product.ID] = product
	return nil

}

func (r InMemoryProductRepository) FindByID(id uuid.UUID) (Product, error) {
	if product, ok := r.products[id]; ok {
		return product, nil
	}
	return nil, errors.New("product not found")

}
// InMemoryOrderRepository implements OrderRepository using in-memory storage

type InMemoryOrderRepository struct {
	orders map[uuid.UUID]*Order
}

func (r InMemoryOrderRepository) Save(order Order) error {
	r.orders[order.ID] = order
	return nil

}

func (r InMemoryOrderRepository) FindByID(id uuid.UUID) (Order, error) {
	if order, ok := r.orders[id]; ok {
		return order, nil
	}
	return nil, errors.New("order not found")

}

```

### Now, let's see how we would use these components in a hypothetical scenario:

```go
func main() {

	// Initialize repositories
	productRepo := &InMemoryProductRepository{
		products: make(map[uuid.UUID]*Product),
	}
	orderRepo := &InMemoryOrderRepository{
		orders: make(map[uuid.UUID]*Order),
	}
	// Initialize services
	productService := &ProductService{productRepo}
	orderService := &OrderService{orderRepo}
	// Example usage: Add a new product
	newProduct, err := productService.AddProduct("Smartphone", 999.99, 50)
	if err != nil {
		fmt.Println("Error adding product:", err)
		return
	}
	fmt.Println("New Product ID:", newProduct.ID)
	// Example usage: Place a new order
	customerID := uuid.New()
	items := []CartItem{
		{ProductID: newProduct.ID, Quantity: 2},
	}
	newOrder, err := orderService.PlaceOrder(customerID, items)
	if err != nil {
		fmt.Println("Error placing order:", err)
		return
	}
	fmt.Println("New Order ID:", newOrder.ID)
}

```

## Explain Event Driven Development.
