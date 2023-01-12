package main

import (
	"fmt"
	"os"
)

/*
Modify the echo program to print the index and he value of each for its arguments, line per line.
*/
func main() {

	for i, arg := range os.Args[0:] {
		fmt.Printf("%v: %v\n", i, arg)
	}

}
