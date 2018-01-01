package main

import (
	"fmt"
	u "utils"
)

func main() {
	a := u.RandomArray(50000)
	fmt.Println("Sorting", len(a), "items")
	bubbleSort(a)
}

/*
*	In bubble sort we compare 2 items and swap them if smaller one is present on the right side.
*	we do this from first item till last, this is called (pass)
*	with every increase in no. of passes we decrease our last index (largest value reaches end with every pass, so no need to compare)
*	we repeat this process untill our last index is equal to our first index, then we terminate.
*
 */
func bubbleSort(a []int) {
	defer u.Elapsed("Bubble sort")()

	// outer loop iterate from last item to second item
	for i := len(a) - 1; i > 1; i-- {

		// inner loop iterate from first item to (value of outer), will get shorter as outer loop inc
		for j := 0; j < i; j++ {

			// when a bigger item found swap it
			if a[j] > a[j+1] {
				u.Swap(&a[j], &a[j+1])
			}
		}
	}
}
