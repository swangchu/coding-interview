// Given a string str, create a function that returns the
// first repeating character.If such character doesn't exist,
// return the null character '\0'.

// Example 1:

// Input: str = "inside code"

// Output: 'i'

// Example 2:

// Input: str = "programming"

// Output: 'r'

// Example 3:

// Input: str = "abcd"

// Output: '\0'

// Example 4:

// Input: str = "abba"

// Output: 'b'

package main

import "fmt"

func main() {
	s := "programming"
	r := firtRepeat(s)
	fmt.Printf("The repeating first repeating charater is %c", r)
}

func firtRepeat(str string) byte {
	present := make(map[byte]bool)

	for i := 0; i < len(str); i++ {
		ch := str[i]
		if present[ch] {
			return ch
		} else {
			present[ch] = true
		}
	}
	return '\x00'
}
