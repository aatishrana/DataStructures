package main

import (
	"fmt"
	u "utils"
)

func main() {
	a := u.RandomArray(50000)
	fmt.Println("Sorting", len(a), "items")
	shellSort(a)
}

func shellSort(a []int) {
	defer u.Elapsed("Shell sort")()
	var temp int
	l := len(a)

	h := 1
	for h <= l/3 {

		//	increment gap sequence till half of array length
		h = h*3 + 1
	}

	fmt.Println(h)

	//	loop till h is reduced to zero
	for h > 0 {

		// outer loop
		for i := h; i < l; i++ {
			temp = a[i]
			j := i
			for j > h-1 && a[j-h] >= temp {
				a[j] = a[j-h]
				j -= h
			}
			a[j] = temp
		}
		h = (h - 1) / 3
	}
}
