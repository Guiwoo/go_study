package main

import "designpattern/state"

/**
Complicated object ts aren't designed from scratch
- They reiterate exisitng desings

An existing desing is a Prototype
We make a copy of the prototype and customize it
Requires 'deep copy' support
We make the cloning convenient via a Factory
*/

// Deep copyting

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{j, i}
		} else {
			m[v] = i
		}
	}
	return nil
}

func main() {
	state.Start()
}
