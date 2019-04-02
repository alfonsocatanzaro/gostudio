package main

import (
	"fmt"
	"strconv"
)

func main() {
	var length int
	var shiftBy int32
	var input1, input2, input3 string

	fmt.Scanf("%s\n", &input1)
	if n, err := strconv.Atoi(input1); err == nil {
		length = int(n)
	}
	fmt.Scanf("%s\n", &input2)

	fmt.Scanf("%s\n", &input3)
	if n, err := strconv.Atoi(input3); err == nil {
		shiftBy = int32(n)
	}

	const minLower, maxLower, minUPPER, maxUPPER = 'a', 'z', 'A', 'Z'
	var output string
	for i := 0; i < length; i++ {
		ch := int32(input2[i])
		if ch >= minLower && ch <= maxLower {
			ch += shiftBy
			if ch > maxLower {
				ch = ch + minLower - maxLower - 1
			}
		}
		if ch >= minUPPER && ch <= maxUPPER {
			ch += shiftBy
			if ch > maxUPPER {
				ch = ch + minUPPER - maxUPPER - 1
			}
		}
		output += string(ch)
	}
	fmt.Println(input2)
	fmt.Println(output)
}
