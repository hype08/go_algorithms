package main

import "fmt"

func greeting(name string) string {
	return "hello " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Henry"))
	fmt.Println(getSum(2, 3))
}
