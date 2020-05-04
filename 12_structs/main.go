package main

import (
	"fmt"
	"strconv"
)

// Person struct.
type Person struct {
	// firstName string
	// lastName  string
	// gender    string
	// age       int
	// city      string

	firstName, lastName, gender, city string
	age                               int
}

// Greeting (value receiver)
func (person Person) greet() string {
	return "hello my name is " + person.firstName + " and I am " + strconv.Itoa(person.age)
}

// (pointer receiver)
func (person *Person) hasBirthday() {
	person.age++
}

func (person *Person) getsMarried(newLastName string) {
	if person.gender == "f" {
		person.lastName = newLastName
	} else {
		return
	}

}
func main() {
	// person1 := Person{firstName: "Henry", lastName: "Zhang", gender: "male", age: 30, city: "Tokyo"}

	person1 := Person{"Henry", "Zhang", "m", "Tokyo", 30}
	fmt.Println(person1)
	fmt.Println(person1.gender)
	person1.age++
	fmt.Println(person1.age)

	fmt.Println(person1.greet())
	person1.hasBirthday()
	fmt.Println(person1.age)
	person1.getsMarried("Cavill")
	fmt.Println(person1.lastName)
}
