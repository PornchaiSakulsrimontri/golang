package main

import(
	"fmt"
)

//Implement a fibonacci function that returns a function (a closure) that returns successive
//fibonacci numbers
//F0, the "0" is omitted (0, 1, 1, 2, 3, 5, ...)

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int{ 
	i, prev, current := -1, 0, 1

	return func() int {
		if i==-1 || i==0{
			i++
			return i
		}
		next:= prev + current
		prev = current
		current = next

		return next
	}
}

func main(){
	f:= fibonacci()
	for i:=0;i<10;i++{
		fmt.Println(f())
	}
}
