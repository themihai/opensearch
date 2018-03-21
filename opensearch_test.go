package opensearch

import (
	"encoding/xml"
	"net/mail"
	"testing"
)

func TestEncoding(t *testing.T) {
	d := Description{
		XMLName: xml.Name{
			Space: "http://a9.com/-/spec/opensearch/1.1/",
			Local: "OpenSearchDescription"},
		ShortName:   "Web Search",
		Description: "Use Example.com to search the Web.",
		Tags:        "example web",
		Contact:     mail.Address{Address: "admin@example.com"},
		URL: []URL{
			URL{
				Type:     "application/atom+xml",
				Template: "http://example.com/?q={searchTerms}&pw={startPage?}&format=atom",
			},
			URL{
				Type:     "application/rss+xml",
				Template: "http://example.com/?q={searchTerms}&pw={startPage?}&format=rss",
			},
			URL{
				Type:     "application/text/html",
				Template: "http://example.com/?q={searchTerms}&pw={startPage?}",
			},
		},
		LongName: "Example.com Web Search",
		Image: []Image{
			Image{
				Height: 64,
				Width:  64,
				Type:   "image/png",
				URL:    "http://example.com/websearch.png",
			},
			Image{
				Height: 16,
				Width:  16,
				Type:   "image/vnd.microsoft.icon",
				URL:    "http://example.com/websearch.ico",
			},
		},
	}
	b, err := xml.Marshal(d)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != e {
		t.Errorf("got %s", string(b))
	}
}

const e = ` <?xml version="1.0" encoding="UTF-8"?>
 <OpenSearchDescription xmlns="http://a9.com/-/spec/opensearch/1.1/">
   <ShortName>Web Search</ShortName>
   <Description>Use Example.com to search the Web.</Description>
   <Tags>example web</Tags>
   <Contact>admin@example.com</Contact>
   <Url type="application/atom+xml"
        template="http://example.com/?q={searchTerms}&amp;pw={startPage?}&amp;format=atom"/>
   <Url type="application/rss+xml"
        template="http://example.com/?q={searchTerms}&amp;pw={startPage?}&amp;format=rss"/>
   <Url type="text/html" 
        template="http://example.com/?q={searchTerms}&amp;pw={startPage?}"/>
   <LongName>Example.com Web Search</LongName>
   <Image height="64" width="64" type="image/png">http://example.com/websearch.png</Image>
   <Image height="16" width="16" type="image/vnd.microsoft.icon">http://example.com/websearch.ico</Image>
   <Query role="example" searchTerms="cat" />
   <Developer>Example.com Development Team</Developer>
   <Attribution>
     Search data Copyright 2005, Example.com, Inc., All Rights Reserved
   </Attribution>
   <SyndicationRight>open</SyndicationRight>
   <AdultContent>false</AdultContent>
   <Language>en-us</Language>
   <OutputEncoding>UTF-8</OutputEncoding>
   <InputEncoding>UTF-8</InputEncoding>
 </OpenSearchDescription>`
