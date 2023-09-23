package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	// very custom logic
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	// very custom logic
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
