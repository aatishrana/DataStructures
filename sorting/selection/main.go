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

func selectionSort(a []int) {
	defer u.Elapsed("Selection sort")()
	var min int
	for i := 0; i < len(a)-1; i++ {
		min = i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[min] {
				min = j
			}
			u.Swap(&a[i], &a[min])
		}
	}
}
