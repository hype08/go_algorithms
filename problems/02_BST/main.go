/*
Write a function that takes in a Binary Search Tree and a target integer value and returns the closest value to that target value.package main

You can assume that there will only be one closest value.
*/

package main

import "math"

// BST struct.
type BST struct {
	Value int
	Left  *BST
	Right *BST
}

func main() {
	FindClosestValue()
}

// FindClosestValue to target.
func (tree *BST) FindClosestValue(target int) int {
	return tree.findClosestValue(target, math.MaxInt32)
}

func (tree *BST) findClosestValue(target, closest int) int {
	currentnode := tree
	for currentnode != nil {
		if absdiff(target, closest) > absdiff(target, currentnode.Value) {
			closest = currentnode.Value
		}
		if target < currentnode.Value {
			currentnode = currentnode.Left
		} else if target > currentnode.Value {
			currentnode = currentnode.Right
		}
		break
	}
	return closest
}

func absdiff(a, b int) int {
	out := math.Abs(float64(a) - float64(b))
	return int(out)
}
