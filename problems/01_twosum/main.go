/*
Write a function that takes a non-empty array of distinct integers and an integer representing a target sum.
If any two numbers in the input array sum up to the target sum, the function should return them in an array, in any order.
If no two numbers sum up to the target sum, the function should return an empty array.
You cannot add a single integer to itself to obtain the target sum.
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	target := 10
	array := []int{3, 5, -4, 8, 11, 1, -1, 6}
	result := TwoNumberSum2(array, target)
	fmt.Println(result)
}

// TwoNumberSum1 O(n^2) time | O(1) space.
func TwoNumberSum1(array []int, target int) []int {
	for i := 0; i < len(array)-1; i++ {
		firstNum := array[i]
		for j := i + 1; j < len(array); j++ {
			secondNum := array[j]
			if firstNum+secondNum == target {
				return []int{firstNum, secondNum}
			}
		}
	}
	return []int{}
}

// TwoNumberSum2 O(nlog(n)) time | O(1) space.
func TwoNumberSum2(array []int, target int) []int {
	sort.Ints(array)
	fmt.Println(array)
	left, right := 0, len(array)-1
	for left < right {
		currentSum := array[left] + array[right]
		if currentSum == target {
			return []int{array[left], array[right]}
		} else if currentSum < target {
			left++
		}
		right--
	}
	return []int{}
}
