package main

import "fmt"

func QuadB(x, y int) {
	// If dimensions are not positive, print nothing
	if x <= 0 || y <= 0 {
		return
	}

	// Loop through each row
	for i := 0; i < y; i++ {
		// Loop through each column
		for j := 0; j < x; j++ {
			// Top row
			if i == 0 {
				// First column of top row
				if j == 0 {
					fmt.Print("/")
				} else if j == x-1 {
					// Last column of top row
					fmt.Print("\\")
				} else {
					// Middle columns of top row
					fmt.Print("*")
				}
			} else if i == y-1 {
				// Bottom row
				// First column of bottom row
				if j == 0 {
					fmt.Print("\\")
				} else if j == x-1 {
					// Last column of bottom row
					fmt.Print("/")
				} else {
					// Middle columns of bottom row
					fmt.Print("*")
				}
			} else {
				// Middle rows
				// Left column
				if j == 0 {
					fmt.Print("*")
				} else if j == x-1 {
					// Right column
					fmt.Print("*")
				} else {
					// Interior (space)
					fmt.Print(" ")
				}
			}
		}
		// New line after each row
		fmt.Println()
	}
}
func main() {
	QuadB(5, 3)
}
