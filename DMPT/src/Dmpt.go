package main

import "fmt"

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

func euclid(x, y int) int {
	if y == 0 {
		return x
	}
	return (euclid(y, x%y))
}

func multipliInverse(x, y int) int {
	var z1, z2 int
	longGCD(x, y, &z1, &z2)
	return (((z1 % y) + y) % y)
}

func longGCD(x int, y int, x1 *int, x2 *int) int {
	var x3, x4, x5 int
	if x == 0 {
		*x1 = 0
		*x2 = 1
		return y
	}
	x5 = longGCD(y%x, x, &x3, &x4)

	*x1 = x4 - (y/x)*x3
	*x2 = x3
	return x5
}

//Driver Function
func main() {
	var ch, x, y, z, temp int
	var n, m, ph, phInt float64
	var prefer = [8][4]int{{7, 5, 6, 4},
		{5, 4, 6, 7},
		{4, 5, 6, 7},
		{4, 5, 6, 7},
		{0, 1, 2, 3},
		{0, 1, 2, 3},
		{0, 1, 2, 3},
		{0, 1, 2, 3}}

	fmt.Println("| -- MENU -- |")
	fmt.Printf("1. Stable Marriage\n2. Euclid's Algorithm\n3. Pigeonhole Principle\n4. Multiplicative Inverse\n")
	fmt.Printf("\nEnter your choice: ")
	fmt.Scanf("%d", &ch)
	fmt.Println("")

	switch ch {
	case 1:
		stablemarriage(prefer)
	case 2:
		fmt.Printf("Enter 1st no.: ") //x = 8923794876682362
		fmt.Scanf("%d", &x)
		fmt.Printf("Enter 2nd no.: ") //y = 9384764854422
		fmt.Scanf("%d", &y)

		if x > y {
			z = euclid(x, y)
			fmt.Printf("GCD of %d & %d is %d\n", x, y, z)
		} else {
			z = euclid(y, x)
			fmt.Printf("GCD of %d & %d is %d\n", x, y, z)
		}
	case 3:
		fmt.Printf("Enter n; no. of pigeons: ")
		fmt.Scanf("%f", &n)
		fmt.Printf("Enter m; no. of pigeonholes: ")
		fmt.Scanf("%f", &m)

		ph = (((n - 1) / m) + 1)
		temp = int(ph)
		phInt = float64(temp)

		if n > m {
			if ph > phInt {
				fmt.Printf("If %.2f pigeons are accommodated in %.2f pigeonholes and since %.2f>%.2f, then one of the pigeonhole must contain at least %.2f pigeons (Acc. to formula %.2f)\n", n, m, n, m, phInt+1, ph)
			} else {
				fmt.Printf("If %.2f pigeons are accommodated in %.2f pigeonholes and since %.2f>%.2f, then one of the pigeonhole must contain at least %.2f pigeons (Acc. to formula %.2f)\n", n, m, n, m, phInt, ph)
			}
		}
	case 4:
		fmt.Printf("For multiplicative inverse to exist both the numbers have to be co-prime.\n")
		fmt.Printf("Enter 1st no.: ") //x = 392
		fmt.Scanf("%d", &x)
		fmt.Printf("Enter 2nd no.: ") //y = 27
		fmt.Scanf("%d", &y)

		if x > y {
			z = euclid(x, y)
			fmt.Printf("GCD of %d & %d is %d\n", x, y, z)
		} else {
			z = euclid(y, x)
			fmt.Printf("GCD of %d & %d is %d\n", x, y, z)
		}

		if z != 1 {
			fmt.Printf("! -- NO SOLUTION -- !\n")
		} else {
			if x > y {
				fmt.Printf("Solution: %d\n", multipliInverse(y, x))
			} else {
				fmt.Printf("Solution: %d\n", multipliInverse(x, y))
			}
		}
	}
}
