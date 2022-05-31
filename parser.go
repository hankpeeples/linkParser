package linkParser

import "io"

// Link represents a link (<a href="">...</a>) in an HTML document.
type Link struct {
	Href string
	Text string
}

// Parse will take in an HTML document and return a slice of Links parsed from it.
func Parse(r io.Reader) ([]Link, error) {
	return nil, nil
}
