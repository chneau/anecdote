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
var Sources map[string]func() ([]Anecdote, error)

// Descriptor of source ...
type Descriptor struct {
	URL     string
	Content string
	Title   string
	Summary string
	Example string
	Replace string
}

func init() {
	Sources = map[string]func() ([]Anecdote, error){}
	Sources["SCMB"] = Builder(Descriptor{
		URL:     "https://secouchermoinsbete.fr/random",
		Content: ".anecdote",
		Title:   "header h1 a",
		Summary: ".summary a",
		Replace: "\n                En savoir plus",
	})
	Sources["SI"] = Builder(Descriptor{
		URL:     "https://www.savoir-inutile.com/",
		Summary: "#phrase",
	})
	Sources["D2R"] = Builder(Descriptor{
		URL:     "http://www.dico2rue.com/mots-au-hasard/",
		Content: ".word_center",
		Title:   "td.word",
		Summary: "tbody > tr:nth-child(2)",
		Example: "td.example",
	})
	Sources["RW"] = Builder(Descriptor{
		URL:     "https://randomword.com/",
		Title:   "#random_word",
		Summary: "#random_word_definition",
	})
}

// Builder returns annecdotes funcs
func Builder(d Descriptor) func() ([]Anecdote, error) {
	return func() ([]Anecdote, error) {
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
}
