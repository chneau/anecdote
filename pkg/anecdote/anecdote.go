package anecdote

import (
	"crypto/tls"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Anecdote ...
type Anecdote struct {
	Title   string
	Summary string
	Example string
}

// String ...
func (a *Anecdote) String() string {
	result := ""
	if a.Title != "" {
		result += a.Title
	}
	if a.Summary != "" {
		if result != "" {
			result += "\n"
		}
		result += a.Summary
	}
	if a.Example != "" {
		if result != "" {
			result += "\n"
		}
		result += a.Example
	}
	return result
}

// Sources ...
var Sources map[string]Descriptor

// Descriptor of source ...
type Descriptor struct {
	URL     string
	Desc    string
	Content string
	Title   string
	Summary string
	Example string
	Replace string
}

// Anecdotes ...
func (d *Descriptor) Anecdotes() ([]Anecdote, error) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	found := []Anecdote{}
	res, err := client.Get(d.URL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	if d.Content == "" {
		d.Content = "body"
	}
	doc.Find(d.Content).Each(func(i int, s *goquery.Selection) {
		a := Anecdote{}
		if d.Title != "" {
			a.Title = strings.TrimSpace(s.Find(d.Title).Text())
		}
		if d.Summary != "" {
			a.Summary = strings.TrimSpace(s.Find(d.Summary).Text())
			a.Summary = strings.Replace(a.Summary, d.Replace, "", -1)
		}
		if d.Example != "" {
			a.Example = strings.TrimSpace(s.Find(d.Example).Text())
			a.Example = strings.Replace(a.Example, d.Replace, "", -1)
		}
		found = append(found, a)
	})
	return found, nil
}

func init() {
	Sources = map[string]Descriptor{
		"SCMB": Descriptor{
			URL:     "https://secouchermoinsbete.fr/random",
			Desc:    "Se Coucher Moins Bete",
			Content: ".anecdote",
			Title:   "header h1 a",
			Summary: ".summary a",
			Replace: "\n                En savoir plus",
		},
		"SI": Descriptor{
			URL:     "https://www.savoir-inutile.com/",
			Desc:    "Savoir Inutile",
			Summary: "#phrase",
		},
		"D2R": Descriptor{
			URL:     "http://www.dico2rue.com/mots-au-hasard/",
			Desc:    "Dico 2 Rue",
			Content: ".word_center",
			Title:   "td.word",
			Summary: "tbody > tr:nth-child(2)",
			Example: "td.example",
		},
		"RW": Descriptor{
			URL:     "https://randomword.com/",
			Desc:    "Random Word",
			Title:   "#random_word",
			Summary: "#random_word_definition",
		},
		"OSD": Descriptor{
			URL:     "http://onlineslangdictionary.com/random-word/",
			Desc:    "Online Slang Dictionary",
			Content: ".term",
			Title:   "h2 > a",
			Summary: ".definitions",
		},
		"LQ": Descriptor{
			URL:     "http://www.litquotes.com/Random-Quote.php",
			Desc:    "Lit Quotes",
			Summary: ".purple > p",
		},
	}
}
