package main

import (
	"fmt"
	"math/rand"
)

var googleDomains = map[string]string{

}

type SearchResult struct {
	ResultRank  int
	ResultUrl   string
	ResultTitle string
	ResultDesc  string
}

var userAgents []string {

}

func randomUserAgnet() string {
	randNum := rand.Int() % len(userAgents)
	return userAgents[randNum]
}

func GoogleScraper() ([]SearchResult, error) {
	results := SearchResult{}
	resultCounter := 0
	buildGoogleUrls()
}

func main() {
	res, err := GoogleScraper("ahmad moshtaghi")
	if err == nil {
		for _, res := range res {
			fmt.Println(res)
		}
	}
}
