package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

/*
Modify dup2 to print the names of all files in wich each duplicated lines occurs
*/
func main() {

	inputFiles := []string{
		"../data/args.dat",
		"../data/args2.dat",
		"../data/args3.dat"}

	counts := make(map[string]int)

	var filesWithDuplicates []string

	if len(inputFiles) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range inputFiles {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v \n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
			// Using Search function
			i := sort.Search(len(inputFiles), func(i int) bool { return arg == inputFiles[i] })

			if i < len(inputFiles) && inputFiles[i] == arg {
				filesWithDuplicates = append(filesWithDuplicates, inputFiles[i])
			}
		}
	}

	fmt.Println(filesWithDuplicates)

	// for line, n := range counts {
	// 	if n > 1 {
	// 		fmt.Printf("%d\t%s\n", n, line)
	// 	}
	// }
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
	}
}
