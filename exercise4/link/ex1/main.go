package main

import (
	"fmt"
	"gostudio/exercise4/link"
	"strings"
)

var exampleHTML = `
<html>

<body>
    <h1>Hello!</h1>
	<a href="/other-page">A link 
	<span>to another</span>
	 page</a>
	<a href="/page-two">A link to a second page</a>
</body>

</html>
`

func main() {
	r := strings.NewReader(exampleHTML)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)

}
