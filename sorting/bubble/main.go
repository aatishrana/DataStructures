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

func bubbleSort(a []int) {
	defer u.Elapsed("Bubble sort")()
	for i := len(a) - 1; i > 1; i-- {
		for j := 0; j < i; j++ {
			if a[j] > a[j+1] {
				u.Swap(&a[j], &a[j+1])
			}
		}
	}
}
