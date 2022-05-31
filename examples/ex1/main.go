package main

import (
	"fmt"
	"strings"

	"github.com/hankpeeples/linkParser"
)

func main() {
	r := strings.NewReader("string")
	data, _ := linkParser.Parse(r)

	fmt.Println(data)
}
