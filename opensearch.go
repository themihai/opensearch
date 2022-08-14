package opensearch

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/mail"
	"net/url"
	"strings"

	"golang.org/x/text/language"
)

type Description struct {
	XMLName     xml.Name `xml:"OpenSearchDescription"`
	ShortName   string
	Description string
	Tags        string
	Contact     mail.Address
	URL         []URL `xml:"Url"`
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

func (d *Description) ResolveReference(base *url.URL) error {
	if base.Scheme == "" {
		return errors.New("baseURI is missing scheme")
	}
	for k := range d.URL {
		nu, err := url.Parse(d.URL[k].Template)
		if err != nil {
			return err
		}
		if nu.Scheme == "" {
			d.URL[k].Template = base.ResolveReference(nu).String()
		}
	}
	return nil
}

// <Url type="application/atom+xml"
// template="http://example.com/?q={searchTerms}&amp;pw={startPage?}&amp;filters={filters}&amp;format=atom"/>
type URL struct {
	XMLName  xml.Name `xml:"Url"`
	Type     string   `xml:"type,attr"`
	Template string   `xml:"template,attr"`
}

// <Image height="64" width="64" type="image/png">http://example.com/websearch.png</Image>

type Image struct {
	XMLName xml.Name `xml:"Image"`
	Height  int      `xml:"height,attr"`
	Width   int      `xml:"width,attr"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
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

func Decode(b []byte) (*Description, error) {
	var d = &Description{}
	err := xml.Unmarshal(b, d)
	return d, err
}

func (d *Description) Write(w io.Writer) error {
	return xml.NewEncoder(w).Encode(d)
}

// typ is the media type (i.e. application/json or atom)
func (d *Description) Request(typ, searchTerms, filters string, startPage string) (*url.URL, error) {
	tpl := d.TemplateByType(typ)
	if tpl == "" {
		return nil, fmt.Errorf("No matching URL %#v", d)
	}
	tpl = FillTemplate(tpl, searchTerms, filters, startPage)
	return url.Parse(tpl)

}
func FillTemplate(tpl, searchTerms, filters, startPage string) string {
	if searchTerms != "" {
		searchTerms = url.QueryEscape(searchTerms)
	}
	if startPage != "" {
		startPage = url.QueryEscape(startPage)
	}
	if filters != "" {
		filters = url.QueryEscape(filters)
	}
	tpl = strings.Replace(tpl, `{searchTerms}`, searchTerms, 1)
	tpl = strings.Replace(tpl, `{filters}`, filters, 1)
	tpl = strings.Replace(tpl, `{startPage}`, startPage, 1)
	tpl = strings.Replace(tpl, `{startPage?}`, startPage, 1)
	return tpl
}
func (d *Description) TemplateByType(typ string) string {
	for _, v := range d.URL {
		if v.Type == typ {
			return v.Template
		}
	}
	return ""
}
