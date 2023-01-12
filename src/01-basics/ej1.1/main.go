package main

import (
	"fmt"
	"os"
)

/*
Modify the echo program to also print os.Args[0], the name of the command line invoked
*/
func main() {

	s, sep := "", ""

	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}
