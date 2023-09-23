package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo // equivalent to contactInfo contactInfo
}

func main() {
	// jim := person{
	// 	firstName: "Jim",
	// 	lastName:  "Party",
	// 	contactInfo: contactInfo{
	// 		email:   "jim@gmail.com",
	// 		zipCode: 12356,
	// 	},
	// }

	// jim.updateName("Jimmy")
	// jim.print()
	mySlice := []string{"Hi", "There", "!"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}
func updateSlice(s []string) {
	s[0] = "Bye"
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
