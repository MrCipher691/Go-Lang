package main

import "fmt"

func main() {
	x := [...]int{11, 234, 123, 56, 345}

	for i, v := range x {
		fmt.Println("Index =", i, "| Values =", v)
	}
}
