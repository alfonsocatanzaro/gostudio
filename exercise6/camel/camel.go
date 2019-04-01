package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	answer := int32(1)
	for _, ch := range input {
		str := string(ch)
		if strings.ToUpper(str) == string(str) {
			answer++
		}
	}
	fmt.Println(answer)

	// scanner := bufio.NewScanner(os.Stdin)

	// for scanner.Scan() {
	// 	word := scanner.Text()
	// 	n := 1
	// 	for i := 0; i < len(word); i++ {
	// 		if word[i] >= 'A' && word[i] <= 'Z' {
	// 			n++
	// 		}
	// 	}
	// 	fmt.Println(n)
	// }

}
