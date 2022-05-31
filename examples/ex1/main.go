package main

import (
	"fmt"
	"strings"

	"github.com/hankpeeples/linkParser"
)

var exampleHtml = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">
    A link to another page
    <span> some span  </span>
  </a>
  <a href="/page-two">A link to a second page</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHtml)
	links, _ := linkParser.Parse(r)

	fmt.Printf("%+v\n", links)
}
