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
### Testing:
- file ending must be in `_test.go`
- run with `go test`
- always do the cleanup