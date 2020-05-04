package main

import "fmt"

// Person struct.
type Person struct {
	firstName string
	lastName  string
	gender    string
	age       int
	city      string
}

func main() {
	// person1 := Person{firstName: "Henry", lastName: "Zhang", gender: "male", age: 30, city: "Tokyo"}
	person1 := Person{"Henry", "Zhang", "male", 30, "Tokyo"}
	fmt.Println(person1)
}
