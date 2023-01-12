package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/TwiN/go-color"
)

/*
Experimento to measure the different in running
time between our potencially inefficiente version
and the one that use strings.Join illustrate
part of the time package.
In orde to provide more significative test case, we use a external file with thousand of random text lines.
*/
func main() {

	start := time.Now()

	//First version.
	s := ""

	// get file from terminal
	// if len(os.Args) != 2 {
	// 	fmt.Printf("Usage: ./main.go <filename> ")
	// 	return
	// }

	filePath := "args.dat" //os.Args[1]
	fileLines := readfile(filePath)

	t := time.Now()
	for value := range fileLines {
		s += strconv.Itoa(value)
	}
	//Only we interest measure the processing time, no include the print time.
	firstTime := t.Sub(start).Nanoseconds()
	fmt.Printf(color.Colorize(color.Green, "First version time: %v nseg\n"), firstTime)
	//fmt.Println(s)

	//Second version using strings.Join
	t = time.Now()
	s = strings.Join(fileLines, "\n")
	secondTime := t.Sub(start).Nanoseconds()
	fmt.Printf(color.Colorize(color.Green, "Second version time: %v nseg \n"), secondTime)
	//fmt.Println(s)
	fmt.Printf("Second version  is %v %% of first version.", (secondTime/firstTime)*100)

}

func readfile(filePath string) []string {

	// read the whole content of file and pass it to file variable, in case of error pass it to err variable
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Could not open the file due to this %s error \n", err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	if err = file.Close(); err != nil {
		fmt.Printf("Could not close the file due to this %s error \n", err)
	}

	return fileLines
}
