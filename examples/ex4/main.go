package main

import (
	"fmt"
	"strings"

	"github.com/hankpeeples/linkParser"
)

var exampleHtml = `
<html>
<body>
  <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHtml)
	links, _ := linkParser.Parse(r)

	fmt.Printf("%+v\n", links)
}
