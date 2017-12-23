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

func insertionSort(a []int) {
	defer u.Elapsed("Insertion sort")()
	for i := 1; i < len(a); i++ {
		temp := a[i]
		j := i
		for j > 0 && a[j-1] >= temp {
			a[j] = a[j-1]
			j--
		}
		a[j] = temp
	}
}
