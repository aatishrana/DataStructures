package main

import (
	"fmt"
	u "utils"
)

func main() {
	a := u.RandomArray(50000)
	fmt.Println("Sorting", len(a), "items")
	selectionSort(a)
}

/*
*	In Selection sort we pick smallest item and place it at first index
*	We repeat this and find second smallest item and place it at second index
*	we repeat this process again and again till we reach end
 */
func selectionSort(a []int) {
	defer u.Elapsed("Selection sort")()
	var min int

	//	outer loop iterate from first item to second last item (no need to check last item, because all smaller items would be shifted)
	for i := 0; i < len(a)-1; i++ {

		//	store current selected item index in min
		min = i

		//	inner loop iterate from selected item to last item
		for j := i + 1; j < len(a); j++ {

			//	if value is less then selected min value then select this index as min
			if a[j] < a[min] {
				min = j
			}

			//	finally swap the value of min value with current pass index
			u.Swap(&a[i], &a[min])
		}
	}
}
