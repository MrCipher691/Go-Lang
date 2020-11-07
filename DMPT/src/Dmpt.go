package main

import "fmt"

// Prints stable matching for 4 men & women. Men are numbered as 0 to 3. Women are numbered as 4 to 7.
func stableMarriage(prefer [8][4]int) {
	// Stores partner of women. This is our output array that stores pairing information. The value of womenPartner[i] indicates the partner assigned to woman 4+i.
	// Note that the woman numbers between 4 and 7. The value -1 indicates that (4+i)'th woman is free. Initialized as free.
	var womenPartner = [4]int{-1, -1, -1, -1}
	// An array to store availability of men. If menFree[i] is false, then man 'i' is free, otherwise engaged. Initialized as free.
	var menFree = [4]bool{false, false, false, false}
	var singleMen int = 4

	// While there are free men
	for singleMen > 0 {
		// Pick the first free man
		var m int
		for m = 0; m < 4; m++ {
			if menFree[m] == false {
				break
			}
		}
		// One by one go to all women according to m's preferences.
		// Here m is the picked free man.
		for i := 0; i < 4 && menFree[m] == false; i++ {
			var w int
			w = prefer[m][i]

			// The woman of preference is free, w and m become
			// partners (partnership may change later).
			if womenPartner[w-4] == -1 {
				womenPartner[w-4] = m
				menFree[m] = true
				singleMen--
			} else { // If w is not free
				var mNext int
				// Find current engagement of w
				mNext = womenPartner[w-4]

				// If w prefers m over her current engagement mNext, then break the engagement between w and mNext and engage m with w.
				if newProposal(prefer, w, m, mNext) == false {
					womenPartner[w-4] = m
					menFree[m] = true
					menFree[mNext] = false
				}
			}
		}
	}
	// Printing the Solution
	fmt.Printf("Man\tWoman\n")
	for i := 0; i < 4; i++ {
		fmt.Printf("%d\t%d\n", womenPartner[i], i+4)
	}
}

// This function returns true if woman 'w' prefers man 'mNext' over man 'm'
func newProposal(prefer [8][4]int, w, m, mNext int) bool {
	// Check if w prefers m over her current engagment mNext
	for i := 0; i < 4; i++ {
		// If mNext comes before m in list of w, then w prefers her
		// Current engagement, don't do anything
		if prefer[w][i] == mNext {
			return true
		}
		// If m comes before mNext in w's list, then free her current engagement and engage her with m
		if prefer[w][i] == m {
			return false
		}
	}
	return false
}

// This function uses Eculid's Algorithm
func euclid(x, y int) int {
	// Stopping condition for recursive loop
	if y == 0 {
		return x
	}
	// Using recursion to calculate GCD of 2 given numbers
	return (euclid(y, x%y))
}

// This function calculates multiplicative inverse of two co-prime numbers
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

// Main/Driver Function
func main() {
	// Declaring all variables required for running the program
	var ch, x, y, z, temp int
	var n, m, ph, phInt float64

	// Preference order of men and women for stable marriage problem
	// Men are numbered 0-3; (A-0, B-1, C-2, D-3)
	// Women are numbered 4-7; (a-4, b-5, c-6, d-7)
	var prefer = [8][4]int{{6, 5, 7, 4}, // Preference of A = c, b, d, a
		{5, 4, 6, 7}, // Preference of B = b, a, c, d
		{6, 4, 7, 5}, // Preference of C = c, a, d, b
		{6, 4, 7, 5}, // Preference of D = c, a, d, b
		{0, 1, 3, 2}, // Preference of a = A, B, D, C
		{2, 0, 3, 1}, // Preference of b = C, A, D, B
		{2, 1, 3, 0}, // Preference of c = C, B, D, A
		{1, 0, 2, 3}} // Preference of d = B, A, C, D

	fmt.Println("| -- MENU -- |")
	fmt.Printf("1. Stable Marriage\n2. Euclid's Algorithm\n3. Pigeonhole Principle\n4. Multiplicative Inverse\n")
	fmt.Printf("\nEnter your choice: ")
	fmt.Scanf("%d", &ch)
	fmt.Println("")

	switch ch {
	case 1:
		// Exectue stable marriage function
		stableMarriage(prefer)
	case 2:
		// Taking 2 integer inputs to calculate their GCD
		fmt.Printf("Enter 1st no.: ") // x = 8923794876682362
		fmt.Scanf("%d", &x)
		fmt.Printf("Enter 2nd no.: ") // y = 9384764854422
		fmt.Scanf("%d", &y)

		// Since function is designed in such a way that first variable has to be > or = to other one.
		// This if-else case is used to decide which variable goes first.
		if x > y {
			// Since x>y, this function is executed.
			z = euclid(x, y)
			// Printing Solution.
			fmt.Printf("GCD of %d & %d is %d\n", x, y, z)
		} else {
			// Since y>=x, this function is executed.
			z = euclid(y, x)
			// Printing Solution.
			fmt.Printf("GCD of %d & %d is %d\n", x, y, z)
		}
	case 3:
		// Taking n and m variables for pigeon hole algorithm formula.
		fmt.Printf("Enter n; no. of pigeons: ") // n = 500
		fmt.Scanf("%f", &n)
		fmt.Printf("Enter m; no. of pigeonholes: ") // m = 45
		fmt.Scanf("%f", &m)

		// Pigeon Hole Algorithm Formula
		ph = (((n - 1) / m) + 1)
		// Converting float answer to int type, to get rid to decimal places.
		temp = int(ph)
		// Converting int back to float since Golang does not allow comparison of variable of 2 different data types.
		phInt = float64(temp)

		// Checking pigeon hole algorithms condition: n>m.
		if n > m {
			// This if-else is defined to round up values so that we get final integer answers.
			if ph > phInt {
				fmt.Printf("If %.2f pigeons are accommodated in %.2f pigeonholes and since %.2f>%.2f,\nthen one of the pigeonhole must contain at least %.2f pigeons (Acc. to formula %.2f)\n", n, m, n, m, phInt+1, ph)
			} else {
				fmt.Printf("If %.2f pigeons are accommodated in %.2f pigeonholes and since %.2f>%.2f,\nthen one of the pigeonhole must contain at least %.2f pigeons (Acc. to formula %.2f)\n", n, m, n, m, phInt, ph)
			}
		}
	case 4:
		fmt.Printf("For multiplicative inverse to exist both the numbers have to be co-prime.\n")
		// Taking 2 integer inputs to calculate multiplicative inverse
		fmt.Printf("Enter 1st no.: ") //x = 392
		fmt.Scanf("%d", &x)
		fmt.Printf("Enter 2nd no.: ") //y = 27
		fmt.Scanf("%d", &y)

		// Since function is designed in such a way that first variable has to be > or = to other one.
		// This if-else case is used to decide which variable goes first.
		if x > y {
			// Since x>y, this function is executed.
			z = euclid(x, y)
			// Printing Solution.
			fmt.Printf("GCD of %d & %d is %d\n", x, y, z)
		} else {
			// Since y>=x, this function is executed.
			z = euclid(y, x)
			// Printing Solution.
			fmt.Printf("GCD of %d & %d is %d\n", x, y, z)
		}

		// Condition to see if multiplicative inverse; the two numbers should be co-prime.
		if z != 1 {
			// The numbers are not co-prime, hence their solution does not exist.
			fmt.Printf("! -- NO SOLUTION -- !\n")
		} else {
			// Since function is designed in such a way that first variable has to be > or = to other one.
			// This if-else case is used to decide which variable goes first.
			if x > y {
				// Printing solution.
				fmt.Printf("Solution: %d\n", multipliInverse(y, x))
			} else {
				// Printing solution.
				fmt.Printf("Solution: %d\n", multipliInverse(x, y))
			}
		}
	}
}
