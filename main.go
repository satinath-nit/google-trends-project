// File: main.go
// Code: GoLang

package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
)

type RSS struct {
	Channel *Channel `xml:"channel"`
	XMLName xml.Name `xml:"rss"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	ItemList    []Item `xml:"item"`
}

type Item struct {
	Title     string `xml:"title"`
	Link      string `xml:"link"`
	Traffic   string `xml:"approx_traffic"`
	NewsItems []News `xml:"news_item"`
}

type News struct {
	Headline     string `xml:"news_item_title"`
	HeadlineLink string `xml:"news_item_url"`
}

func main() {
	var r RSS

	data := readGoogleTrends()

	err := xml.Unmarshal(data, &r)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("\n Rss Feed: \n")

	for _, item := range r.Channel.ItemList {
		fmt.Println("Title: ", item.Title)
		fmt.Println("Link: ", item.Link)
		fmt.Println("Traffic: ", item.Traffic)
		fmt.Println("News Items: ")
		for _, news := range item.NewsItems {
			fmt.Println("Headline: ", news.Headline)
			fmt.Println("Headline Link: ", news.HeadlineLink)
		}
		fmt.Println()
	}

}

func readGoogleTrends() []byte {
	resp := getGoogleTrends()
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return data
}

func getGoogleTrends() *http.Response {
	resp, err := http.Get("https://trends.google.com/trends/trendingsearches/daily/rss?geo=US")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return resp

}
