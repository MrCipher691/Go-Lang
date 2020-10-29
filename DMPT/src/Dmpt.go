package main

import "fmt"

// This function returns true if woman 'w' prefers man 'm1' over man 'm'
func newProposal(prefer [8][4]int, w, m, m1 int) bool {
	// Check if w prefers m over her current engagment m1
	for i := 0; i < 4; i++ {
		// If m1 comes before m in list of w, then w prefers her
		// Current engagement, don't do anything
		if prefer[w][i] == m1 {
			return true
		}
		// If m comes before m1 in w's list, then free her current engagement and engage her with m
		if prefer[w][i] == m {
			return false
		}
	}
	return false
}

// Prints stable matching for 4 men & women. Men are numbered as 0 to 3. Women are numbereed as 4 to 7.
func stablemarriage(prefer [8][4]int) {
	// Stores partner of women. This is our output array that stores pairing information. The value of wPartner[i] indicates the partner assigned to woman 4+i.
	// Note that the woman numbers between 4 and 7. The value -1 indicates that (4+i)'th woman is free. Initialized as free.
	var wPartner = [4]int{-1, -1, -1, -1}
	// An array to store availability of men. If mFree[i] is false, then man 'i' is free, otherwise engaged. Initialized as free.
	var mFree = [4]bool{false, false, false, false}
	var singleMen int = 4

	// While there are free men
	for singleMen > 0 {
		// Pick the first free man
		var m int
		for m = 0; m < 4; m++ {
			if mFree[m] == false {
				break
			}
		}
		// One by one go to all women according to m's preferences.
		// Here m is the picked free man.
		for i := 0; i < 4 && mFree[m] == false; i++ {
			var w int
			w = prefer[m][i]

			// The woman of preference is free, w and m become
			// partners (partnership may change later).
			if wPartner[w-4] == -1 {
				wPartner[w-4] = m
				mFree[m] = true
				singleMen--
			} else { // If w is not free
				var m1 int
				// Find current engagement of w
				m1 = wPartner[w-4]

				// If w prefers m over her current engagement m1, then break the engagement between w and m1 and engage m with w.
				if newProposal(prefer, w, m, m1) == false {
					wPartner[w-4] = m
					mFree[m] = true
					mFree[m1] = false
				}
			}
		}
	}
	// Printing the Solution
	fmt.Printf("Man\tWoman\n")
	for i := 0; i < 4; i++ {
		fmt.Printf("%d\t%d\n", wPartner[i], i+4)
	}
}

//Driver Function
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
