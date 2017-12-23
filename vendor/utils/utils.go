package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomArray(l int) []int {
	m := 100000
	a := make([]int, l, l)
	for i := 0; i < l; i++ {
		a[i] = rand.Intn(m)
	}
	return a
}

func Swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func Elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}
