package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	arr := []int32{19, 10, 12, 10, 24, 25, 22}
	n := nonDivisibleSubset(4, arr) - 1
	fmt.Printf("%d\n", n)

}

func nonDivisibleSubset(k int32, S []int32) int32 {
	comb := int32(math.Pow(2, float64(len(S))) - 1)
	ret := 0
	for n := comb; n > 0; n-- {
		form := "%0" + strconv.Itoa(len(S)) + "b"
		bits := fmt.Sprintf(form, n)
		c := strings.Count(bits, "1")
		if c <= ret {
			continue
		}
		if arrIsOk(bits, &S, k) {
			ret = c
		}
	}

	return int32(ret)

}

func arrIsOk(bits string, s2 *[]int32, k int32) bool {
	for a := 0; a < len(bits); a++ {
		if bits[a] == '0' {
			continue
		}
		for b := a + 1; b < len(*s2); b++ {
			if bits[b] == '0' {
				continue
			}
			if (((*s2)[a] + (*s2)[b]) % k) == 0 {
				return false
			}
		}
	}
	return true
}
