package main

import "fmt"

func greeting(name string) string {
	return "hello " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func variadicSum(nums ...int) {
	fmt.Print(nums, "  ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	fmt.Println(greeting("Henry"))
	fmt.Println(getSum(2, 3))
	variadicSum(1, 2)
	variadicSum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	variadicSum(nums...)
}
