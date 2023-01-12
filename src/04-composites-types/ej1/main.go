package main

import (
	"fmt"
)

func main() {

	var a [3]int

	//Print the array
	fmt.Println(a)
	//Print the first and last element
	fmt.Println(a[0])
	fmt.Println(a[2])

	a[1] = 1

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

}
