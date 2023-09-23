# Go Basics
### Packages: 
collection of common source code files

two types of packages:
- executable: generate a file that we can run (package **main** makes executable package -> generate project binaries) -> must have func called name (entrypoint of the program)
- reusable: code used as helpers
___
### Variables:
- statically types

```go
var x string = "string";
x := "string"
```
Primitive types in Go:
- bool
- string
- int
- float64
___
### Functions Literals:
- It's lambda function (anonymous function)
```go
go func() {
	time.Sleep(5 * time.Second)
	checkLink(l, c)
 }()
```
___
### Array vs Slices:
- slice: array that can grow or shrink
- array: fixed size

append to a slice:
```go
cards := []string{"Ace of diamonds", newCard()}
cards = append(cards, "Six of Spades")
```


iterate of slice:
```go
	cards := []string{"Ace of diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	for index, element := range cards {
		fmt.Println(index, element)
	}
```

slicing:
```go
slice[startIndexIncluding: upToNotIncluding]
```

```go
fruits := []string{"apple", "banana", "mango"}
fruits[0:2] // "["apple", "banana"]"
```

____
### Receiver Functions:
sets methods for type we create
```go
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
```

d: actual copy of the deck variable -> like this or self 
___
### Struct:
- Like `dict` in py or as a plain `object` in js
- To define a struct
```go
type person struct {
	firstName string
	listName  string
}
// usage by order of the struct fields
alex := person{"Alex", "Anderson"}

// usage by key-value
alex := person{firstName: "Alex", lastName: "Anderson"}

```
 - Zero value struct

 ```go
 func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// alex.lastName = "Chris"
	var alex person         // value of type person with a zero value
	fmt.Println(alex)       // { }
	fmt.Printf("%+v", alex) // {firstName: lastName:}% -> firstName: "", lastName: ""
}
```
 - Embed structs

```go

```
___
### Pointers:
-  this is a pass by value (a copy from the real variable)
  ```go
// here we update the copy of p
func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
```
- `&variable` -> give me address of that variable
- `*pointer` -> give me value of this memory address
```go
func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 12356,
		},
	}
	jimPointer := &jim
	println((*jimPointerPointer))
	jimPointer.updateName("Jimmy")
	jim.print()
}
// *person is a type of a pointer to type person
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
```

- Pointer Shortcut: 
```go
func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 12356,
		},
	}
// Pointer shortcut
	jim.updateName("Jimmy")
	jim.print()
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
```
____
### Reference vs Value Types:
NOTE THAT: Everything in go is pass by value
```go
package main

import "fmt"
 
func main() {
 name := "bill"
 
 namePointer := &name
 
 fmt.Println(&namePointer)// different address
 printPointer(namePointer)
}
 
func printPointer(namePointer *string) {
 fmt.Println(&namePointer)// different address as everything in go is a pass by value
}
```
- These are reference types in go (don't worry about using pointers with these):
	- `slice`
	- `map`
	- `channel`
	- `pointer`
	- `function`
- These are value types in go (Use pointers to change values in these types):
	- `int`
	- `float`
	- `string`
	- `bool`
	- `struct`
____
### Maps:
- to declare a map:`map[keyType]valueType{}`
```go
// to construcr an empty map
var colors map[string]string
// or
colors := make(map[string]string)

// to construct map with initial values
colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
```
- delete a key in a map
```go
	delete(colors, "white")
```
- iterate over a map:
```go
func printMap(m map[keyType]valType){
	for key, val := range m {
		fmt.Println(key, val)
	} 
}
```
____
### Interfaces:
- concrete type -> I can create a val directly from
- interface type
- interfaces are not generic types
- interfaces are implicit -> no code needed to explicit define some type satisfy custom interface
- interfaces are contract to help us manage types
- interfaces are tough, understand how to read them
Rules of interfaces
```go
type bot interface {
	fn(parameterTypes) (returnTypes)
}
``` 
____
### Channels and Go Routines:
- **Go routines**:
	-  a lightweight thread managed by the Go runtime.
	- spinned by a `go` keyword
	- go could have thousands of go routines running concurrently
	- `go fn()` -> spins a new goroutine(lightweight thread) and runs `fn()` inside it
	- scheduler runs **one** routine until it finishes or makes a blocking call
	- Go uses one CPU core by default
	- **concurrency**: we can have multiple threads executing code. If one thread blocks, another one is picked up and executed, one thread is running at time. One CPU
	- **parallelism**: multiple threads executed at the same time. Requires multiple CPU's
	- main goroutine: launched at the start of the program. Controls when our programs ends.
	- child goroutines: created by `go` keyword
- **Channels**:
	- mediator between goroutines
	- the only way to communicate between running goroutines. Two Way Messaging
	- it's a Go type
	- has a type for their messages (can't share `int` messages in `string` channels)
	- `channel <- 5`: send value 5 to this channel
	- `myNumber <- channel`: wait for a value to be sent into the channel. When we get one, assign to `myNumber`
	- `fmt.Println(<- channel)`: wait for a value to be sent into the channel. When we get one, print it immediately
- **Blocking Channels**:
	- receiving messages from channels is blocking
	- our program might pauses the execution while waiting for goroutines
___
### Testing:
- file ending must be in `_test.go`
- run with `go test`
- always do the cleanup