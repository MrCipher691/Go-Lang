package main

import "fmt"

func newProposal(prefer [8][4]int, w, m, m1 int) bool {
	for i := 0; i < 4; i++ {
		if prefer[w][i] == m1 {
			return true
		}
		if prefer[w][i] == m {
			return false
		}
	}
	return false
}

func stablemarriage(prefer [8][4]int) {
	var wPartner = [4]int{-1, -1, -1, -1}
	var mFree = [4]bool{false, false, false, false}
	var freeCount int = 4

	for freeCount > 0 {
		var m int
		for m = 0; m < 4; m++ {
			if mFree[m] == false {
				break
			}
		}
		for i := 0; i < 4 && mFree[m] == false; i++ {
			var w int
			w = prefer[m][i]

			if wPartner[w-4] == -1 {
				wPartner[w-4] = m
				mFree[m] = true
				freeCount--
			} else {
				var m1 int
				m1 = wPartner[w-4]
				if newProposal(prefer, w, m, m1) == false {
					wPartner[w-4] = m
					mFree[m] = true
					mFree[m1] = false
				}
			}
		}
	}
	fmt.Printf("Man\tWoman\n")
	for i := 0; i < 4; i++ {
		fmt.Printf("%d\t%d\n", wPartner[i], i+4)
	}
}

func main() {
	var prefer = [8][4]int{{7, 5, 6, 4},
		{5, 4, 6, 7},
		{4, 5, 6, 7},
		{4, 5, 6, 7},
		{0, 1, 2, 3},
		{0, 1, 2, 3},
		{0, 1, 2, 3},
		{0, 1, 2, 3}}

	stablemarriage(prefer)
}
