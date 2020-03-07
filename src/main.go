package main

import (
	"fmt"
	"os"

	"github.com/Mic-U/gofeed_line_notifier/src/selector"
	"github.com/mmcdole/gofeed"
)

var (
	//RSSURL RSSフィードの URL
	RSSURL = os.Getenv("RSS_URL")
)

func main() {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(RSSURL)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	items := feed.Items
	selected := selector.SelectNewNotifications(items)

	if len(selected) == 0 {
		fmt.Println("No new articles")
		return
	}
	// message := notifier.Message{
	// 	Type: "text",
	// 	Text: item.Title + "\n" + item.Link,
	// }
	// messages := notifier.PostMessages{Messages: []notifier.Message{message}}
	// notifier.PostBroadcast(messages)
}
