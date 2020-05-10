package main

import "fmt"

func main() {
	// make an empty map of key string, value int.
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map: ", m)

	value1 := m["k1"]
	fmt.Println("value of k1: ", value1)
	fmt.Println("length of m: ", len(m))

	// built-in delete
	delete(m, "k2")
	fmt.Println("m after deletion: ", m)

	// if key is present
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// declare and init a new map in the same line.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("n: ", n)
}
