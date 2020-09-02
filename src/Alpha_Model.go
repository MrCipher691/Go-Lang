package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	var val int
	var err error
	var data string
	slice := make([]int, 0, 3)

	for {
		fmt.Println("Enter integer (press 'x' to exit):")
		fmt.Scan(&data)

		if data == "x" {
			return
		}

		if val, err = strconv.Atoi(data); err != nil {
			fmt.Println("! Invalid Input !")
			continue
		}

		slice = append(slice, val)
		sort.Ints(slice)
		fmt.Println(slice)

		if len(slice) == cap(slice) {
			newData := make([]int, cap(slice), (cap(slice) + 1))
			copy(newData, slice)
			slice = newData
		}
	}
}
