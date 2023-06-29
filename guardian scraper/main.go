package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"strings"
)

func getRequest(targetUrl string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", targetUrl, nil)
	if err!= nil

}

func discoverLinks(response *http.Response, baseUrl string) []string {
	if response != nil {
		doc, _ := goquery.NewDocumentFromResponse(response)
		foundUrls := []string{}
		if doc != nil {
			doc.Find("a").Each(func(i int, selection *goquery.Selection) {
				res, _ := selection.Attr("href")
				foundUrls = append(foundUrls, res)
			})
		}
		return foundUrls
	} else {
		return []string{}
	}
}

func checkRelative(href string, baseUrl string) string {
	if strings.HasPrefix(href, "/") {
		return fmt.Sprintf("%s%s", baseUrl, href)
	} else {
		return href
	}
}

func resolveRelativeLinks(href string, baseUrl string) (bool, string) {
	resultHref := checkRelative(href, baseUrl)
	baseParse, _ := url.Parse(baseUrl)
	resultParse, _ := url.Parse(resultHref)
	if baseParse != nil && resultParse != nil {
		if baseParse.Host == resultParse.Host {
			return true, resultHref
		} else {
			return false, ""
		}
	}
}

var tokens = make(chan struct{}, 5)

func Crawl(targetUrl string, baseUrl string) []string {
	fmt.Println(targetUrl)
	tokens <- struct{}{}
	resp, _ := getRequest(targetUrl)
	<-tokens
	links := discoverLinks(resp, baseUrl)
	foundUrls := []string{}

	for _, link := range links {
		ok, correctLink := resolveRelativeLinks(link, baseUrl)
		if ok {
			if correctLink != "" {
				foundUrls = append(foundUrls, correctLink)
			}
		}
	}
	return foundUrls
}

func ParseHTML(response *http.Response) {

}

func main() {
	worklist := make(chan []string)
	var n int
	n++
	baseDomain := "https://www.theguardian.com"
	go func() { worklist <- []string{"https://www.theguardian.com"} }()
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string, baseUrl string) {
					foundLinks := Crawl(link, baseDomain)
					if foundLinks != nil {
					}
					worklist <- foundLinks
				}(link, baseDomain)
			}
		}
	}
}
