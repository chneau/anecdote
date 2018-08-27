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

// SCMB Se Coucher Moins Bete
func SCMB() ([]Anecdote, error) {
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

// SI Savoir Inutile
func SI() ([]Anecdote, error) {
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
	doc.Find("#phrase").Each(func(i int, s *goquery.Selection) {
		summary := s.Text()
		found = append(found, Anecdote{
			Summary: summary,
			Title:   "",
		})
	})
	return found, nil
}
