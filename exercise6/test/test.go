package main

import (
	"fmt"
)

func main() {
	arr := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n := nonDivisibleSubset(4, arr)
	fmt.Printf("%d\n", n)

}

func nonDivisibleSubset(k int32, S []int32) int32 {
	resti := make(map[int32]int32)
	for i := k - 1; i > -1; i-- {
		resti[i] = 0
	}
	for _, n := range S {
		resti[n%k]++
	}

	if resti[0] > 1 {
		resti[0] = 1
	}
	if (k%2 == 0) && resti[k/2] > 1 {
		resti[k/2] = 1
	}

	for n := k - 1; n > k/2; n-- {
		n1 := k - n
		if resti[n] > resti[n1] {
			resti[n1] = 0
		} else {
			resti[n] = 0
		}
	}

	ret := int32(0)
	for _, c := range resti {
		ret += c
	}

	return ret

}
