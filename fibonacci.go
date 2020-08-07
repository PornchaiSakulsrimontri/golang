package main

import (
	"fmt"
)

//Implement a fibonacci function that returns a function (a closure) that returns successive
//fibonacci numbers
//F0, the "0" is omitted (0, 1, 1, 2, 3, 5, ...)

// fibonacci is a function that return
// a function that returns an int.
func fibonacci() func() int {
	prev, current, next := -1, 1, 0

	return func() int {
		next = prev + current
		prev, current = current, next
		return next
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
