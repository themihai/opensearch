package opensearch

import (
	"encoding/xml"
	"fmt"
	"golang.org/x/text/language"
	"io"
	"net/mail"
	"net/url"
	"strings"
)

type Description struct {
	XMLName     xml.Name `xml:"OpenSearchDescription"`
	ShortName   string
	Description string
	Tags        string
	Contact     mail.Address
	URL         []URL
	// Example.com Web Search
	LongName string
	Image    []Image
	Query    Query
	// Example.com Development Team
	Developer string
	// Search data Copyright 2005, Example.com, Inc., All Rights Reserved
	Attribution string
	// open
	SyndicationRight string
	AdultContent     bool
	Language         language.Tag
	// UTF-8
	OutputEncoding string
	InputEncoding  string
}

// <Url type="application/atom+xml"
// template="http://example.com/?q={searchTerms}&amp;pw={startPage?}&amp;format=atom"/>
type URL struct {
	Type     string
	Template string
}

// <Image height="64" width="64" type="image/png">http://example.com/websearch.png</Image>

type Image struct {
	Height int
	Width  int
	Type   string
	URL    string
}

// <Query role="example" searchTerms="cat" />
type Query struct {
	Role       string
	SearchTerm string
}

func New(r io.Reader) (*Description, error) {
	var d = &Description{}
	err := xml.NewDecoder(r).Decode(d)
	return d, err
}

func (d *Description) Write(w io.Writer) error {
	return xml.NewEncoder(w).Encode(d)
}

func (d *Description) Request(typ, q, startPage string) (*url.URL, error) {
	q = url.QueryEscape(q)
	startPage = url.QueryEscape(startPage)
	for _, v := range d.URL {
		if v.Type == typ {
			v.Template = strings.Replace(v.Template, `{searchTerms}`, q, 1)
			v.Template = strings.Replace(v.Template, `{startPage}`, startPage, 1)
			return url.Parse(v.Template)
		}
	}
	return nil, fmt.Errorf("No matching URL %#v", d)
}
