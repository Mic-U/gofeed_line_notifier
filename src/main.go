package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Mic-U/gofeed_line_notifier/src/notifier"
	"github.com/Mic-U/gofeed_line_notifier/src/selector"
	"github.com/mmcdole/gofeed"
)

var (
	//RSSURL RSSフィードの URL
	RSSURL = os.Getenv("RSS_URL")
)

func main() {
	fp := gofeed.NewParser()

	for {
		exec(fp)
		time.Sleep(time.Minute * 5)
	}
}

func exec(fp *gofeed.Parser) {
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

	// TODO 多すぎる場合にどうするかは要検討
	if len(selected) > 5 {
		selected = selected[0:5]
	}
	messages := convertItemsToMessages(selected)
	postMessages := notifier.PostMessages{
		Messages: messages,
	}
	notifier.PostBroadcast(postMessages)
}

func convertItemsToMessages(items []*gofeed.Item) []notifier.Message {
	messages := make([]notifier.Message, 0)
	for _, i := range items {
		message := notifier.Message{
			Type: "text",
			Text: i.Title + "\n" + i.Link,
		}
		messages = append(messages, message)
	}
	return messages
}
