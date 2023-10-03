package main

import "fmt"

func newGenericFunc[age any](myAge age) {
	fmt.Println(myAge)
}

func main() {
	fmt.Println("Go genrics")

	var testAge int64 = 29
	var testAge2 float32 = 28.6
	newGenericFunc(testAge)
	newGenericFunc(testAge2)
}