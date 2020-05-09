package main

import (
	"fmt"
)

func main() {
	// Arrays
	// initialize an empty array.
	var a [5]int
	fmt.Println("empty", a)

	// set array value by key.
	a[4] = 100
	fmt.Println("set a[4]", a)
	fmt.Println("len", len(a))

	// initialize an array.
	b := [4]int{1, 2, 3, 4}
	fmt.Println("b: ", b)
	fmt.Println("b[3]: ", b[3])

	// Two Dimensional Array
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	//Slices

	// make empty slice of length 3.
	s := make([]string, 10)
	fmt.Println("s: ", s)

	s[3] = "a"
	fmt.Println("s: ", s)
	s = append(s, "e", "f")
	fmt.Println("s: ", s)

	// copy slice.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("c: ", c)

	l := s[(len(s))-1]
	fmt.Println("l: ", l)

	// two dimensional slice with varying length of inner slice

	twoS := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoS[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoS[i][j] = i + j
		}
	}
	fmt.Println(twoS)
}
