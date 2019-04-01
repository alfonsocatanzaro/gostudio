package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	//output := ""
	n1 := 10
	n2 := n1 << 2
	n3 := n2 << 2
	fmt.Printf("%b \t %d \n ", n1, n1)
	fmt.Printf("%b \t %d \n ", n2, n2)
	fmt.Printf("%b \t %d \n ", n3, n3)
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
