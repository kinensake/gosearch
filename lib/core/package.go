package core

import (
	"io"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

const (
	baseURL = "https://pkg.go.dev/search?q="
	bound   = 5 // max number of packages
)

/*
 * Crawl search page html with keyword
 * @params (keyword of package)
 * @return (html body)
 */
func getSearchPage(keyword string) io.ReadCloser {
	endpoint := baseURL + url.QueryEscape(keyword)
	res, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}

	return res.Body
}

/*
 * Get package path list from crawled HTML
 * @params (keyword of package)
 * @return (package path list)
 */
func GetPackagePath(keyword string) []string {
	body := getSearchPage(keyword)
	defer body.Close()

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		panic(err)
	}

	paths := []string{}
	doc.Find(".SearchSnippet-header-path").Each(func(i int, s *goquery.Selection) {
		path := s.Text()
		paths = append(paths, path[1:len(path)-1])
	})

	limit := bound
	if len(paths) < limit {
		limit = len(paths)
	}
	return paths[0:limit]
}
