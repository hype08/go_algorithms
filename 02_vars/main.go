package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	name, email := "Henry", "henry@email.com"

	// Using var
	age := 30
	const isCool = true

	fmt.Println(name, email, age, isCool)
	fmt.Printf("%T\n", isCool)
}
