package linkParser

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link represents a link (<a href="">...</a>) in an HTML document.
type Link struct {
	Href string
	Text string
}

// Parse will take in an HTML document and return a slice of Links parsed from it.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	nodes := linkNodes(doc)
	var links []Link

	for _, node := range nodes {
		links = append(links, buildLink(node))
	}

	// dfs(doc, "")
	return links, nil
}

func buildLink(n *html.Node) Link {
	var ret Link

	// get attribute (href)
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}
	ret.Text = text(n)
	return ret
}

func text(n *html.Node) string {
	// base case, no need to look at children
	if n.Type == html.TextNode {
		return n.Data
	}

	if n.Type != html.ElementNode {
		return ""
	}

	var ret string
	// for every child, get the text and add it to the return variable
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c)
	}
	return strings.Join(strings.Fields(ret), " ")
}

func linkNodes(n *html.Node) []*html.Node {
	// base case
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}

// func dfs(n *html.Node, padding string) {
// 	msg := n.Data
// 	if n.Type == html.ElementNode {
// 		msg = "<" + msg + ">"
// 	}
// 	fmt.Println(padding, msg)
// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		dfs(c, padding+"  ")
// 	}
// }
