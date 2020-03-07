package main

import (
	"fmt"

	"github.com/Mic-U/gofeed_line_notifier/src/notifier"
	"github.com/mmcdole/gofeed"
)

const (
	//RSSURL RSSフィードの URL
	RSSURL = "http://feeds.feedburner.com/AmazonWebServicesBlog"
)

func main() {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(RSSURL)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	items := feed.Items
	item := items[0]
	message := notifier.Message{
		Type: "text",
		Text: item.Title + "\n" + item.Link,
	}
	messages := notifier.PostMessages{Messages: []notifier.Message{message}}
	fmt.Println(messages)
	notifier.PostBroadcast(messages)
}
