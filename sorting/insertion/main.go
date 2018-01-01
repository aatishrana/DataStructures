package main

import (
	"fmt"
	u "utils"
)

func main() {
	a := u.RandomArray(50000)
	fmt.Println("Sorting", len(a), "items")
	insertionSort(a)
}

/*
*	In insertion sort we consider our data divided in 2 groups.
*	First group contains partially sorted items and second group contains unsorted items.
*	Our goal is to take an item from unsorted group and put it in sorted group.
*	While doing so we shift all the items in sorted group until the right place of our selected item is created.
*	We start our algorithm considering our first item as the sorted group(one item is sorted, i.e: no one to compare)
*	When there are no items left in unsorted group algorithm terminates
*
 */
func insertionSort(a []int) {
	defer u.Elapsed("Insertion sort")()

	// outer loop iterate from second item to last item (first item is considered as a sorted group containing only one item)
	for i := 1; i < len(a); i++ {

		// we hold our current selected data and it's index in temporary variables
		temp := a[i]
		j := i

		// inner loop iterate from our selected index to an index at which our selected data can be placed,
		// if not found then first index
		for j > 0 && a[j-1] >= temp {
			a[j] = a[j-1]
			j--
		}

		// finally we place our selected data at its appropriate position
		a[j] = temp
	}
}
