package main

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

const (
	//RSSURL RSSフィードの URL
	RSSURL = "http://feeds.feedburner.com/AmazonWebServicesBlog"
)

func main() {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(RSSURL)

	items := feed.Items
	for _, item := range items {
		fmt.Println(item.Title)
		fmt.Println(item.Link)
		fmt.Println(item.PublishedParsed)
	}
}
