package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {

	readinput := bufio.NewScanner(os.Stdin)
	fmt.Printf("\nPlease enter the filename containing a list of names: ")
	readinput.Scan()
	filename := readinput.Text()

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("\nFile could not be opened. Please make sure you have the")
		fmt.Println("correct filename and path before running again.")
	}

	nameslice := make([]name, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		namestrings := strings.Fields(scanner.Text())
		p := name{fname: namestrings[0], lname: namestrings[1]}
		nameslice = append(nameslice, p)
	}

	fmt.Println("\nHere are your list of names from", filename)

	for i := 0; i < len(nameslice); i++ {
		fmt.Println(nameslice[i].fname, nameslice[i].lname)
	}

}
