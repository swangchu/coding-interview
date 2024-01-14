// Solve the problem [Python]
// Given an array of integers arr, create a function that
// returns an array containing the values of arr without duplicates (the order doesn't matter).

// Example 1:

// Input: arr = [4, 2, 5, 3, 3, 1, 2, 4, 1, 5, 5, 5, 3, 1]

// Output: [4, 2, 5, 3, 1]

// Example 2:

// Input: arr = [1, 1, 1, 1, 1, 1, 1, 1]

// Output: [1]

// Example 3:

// Input: arr = [4, 4, 2, 3, 2, 2, 4, 3]

// Output: [4, 2, 3]

package main

import "fmt"

func main() {

	dublicate := []int{23, 34, 56, 12, 56, 23}
	r := removeDub(dublicate)
	fmt.Println(r)
}

func removeDub(arr []int) []int {
	noDuplicate := []int{}
	visited := make(map[int]bool)

	for _, element := range arr {
		if !visited[element] {
			noDuplicate = append(noDuplicate, element)
			visited[element] = true
		}
	}
	return noDuplicate
}
