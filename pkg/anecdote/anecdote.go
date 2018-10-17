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
}

// String ...
func (a *Anecdote) String() string {
	result := ""
	if a.Title != "" {
		result += a.Title + "\n"
	}
	result += a.Summary
	return result
}

// Sources ...
var Sources map[string]func() ([]Anecdote, error)

func init() {
	Sources = map[string]func() ([]Anecdote, error){}
	Sources["SCMB"] = scmb
	Sources["SI"] = si
	Sources["D2R"] = d2r
}

func scmb() ([]Anecdote, error) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	found := []Anecdote{}
	res, err := client.Get("https://secouchermoinsbete.fr/random")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	doc.Find(".anecdote").Each(func(i int, s *goquery.Selection) {
		title := s.Find("header h1 a").Text()
		summary := s.Find(".summary a").Text()
		summary = strings.Replace(summary, "\n                En savoir plus", "", -1)
		found = append(found, Anecdote{
			Summary: summary,
			Title:   title,
		})
	})
	return found, nil
}

func si() ([]Anecdote, error) {
	found := []Anecdote{}
	res, err := http.Get("https://www.savoir-inutile.com/")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	found = append(found, Anecdote{
		Summary: doc.Find("#phrase").First().Text(),
		Title:   "",
	})
	return found, nil
}

func d2r() ([]Anecdote, error) {
	found := []Anecdote{}
	res, err := http.Get("http://www.dico2rue.com/mots-au-hasard/")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	body := doc.Find(".word_center").First()
	word := strings.TrimSpace(body.Find("td.word").First().Text())
	def := strings.TrimSpace(body.Find("tbody > tr:nth-child(2)").First().Text())
	ex := strings.TrimSpace(body.Find("td.example").First().Text())
	found = append(found, Anecdote{
		Summary: def + "\n" + ex,
		Title:   word,
	})
	return found, nil
}
