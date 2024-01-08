// Given an array of integers arr and an integer k, create a boolean function
// that checks if there exist two elements in arr such that we get k when we
// add them together.

// Example 1:

// Input: arr = [4, 5, 1, -3, 6], k = 11

// Output: true

// Explanation: 5 + 6 is equal to 11

package main

import (
	"fmt"
	"slices"
)

func main() {
	n := []int{12, 5, 7, 2, 9}
	k := 11

	if fPair(n, k) {
		fmt.Printf("There is a pair that sums up to %d in the list", k)
	} else {
		fmt.Println("There is no pair in the list that sums up ", k)
	}

}

func fPair(list []int, k int) bool {
	sortedList := slices.Sort(list)
	left := 0
	right := len(sortedList)

	for left < right {
		if sortedList[left]+sortedList[right] == k {
			return true
		} else if sortedList[left]+sortedList[right] < k {
			left++
		} else {
			right--
		}
	}
	return false
}
