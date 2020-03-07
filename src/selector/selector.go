package selector

import (
	"fmt"
	"time"

	"github.com/mmcdole/gofeed"
)

var (
	latestPublished = time.Now().UTC()
)

// SelectNewNotifications select items later than latestPublished
func SelectNewNotifications(items []*gofeed.Item) []*gofeed.Item {
	selected := selectMap(isLaterThanLatestPublished, items)
	updateLatestPublished(selected)
	return selected
}

func selectMap(f func(i *gofeed.Item) bool, s []*gofeed.Item) []*gofeed.Item {
	ans := make([]*gofeed.Item, 0)
	for _, x := range s {
		if f(x) == true {
			additionalItem := &x
			ans = append(ans, *additionalItem)
		}
	}
	return ans
}

func isLaterThanLatestPublished(item *gofeed.Item) bool {
	publishedAt := item.PublishedParsed.UTC()
	return publishedAt.After(latestPublished)
}

func updateLatestPublished(selected []*gofeed.Item) {
	for _, item := range selected {
		fmt.Println(item.PublishedParsed)
		if item.PublishedParsed.UTC().After(latestPublished) {
			latestPublished = item.PublishedParsed.UTC()
		}
	}
	fmt.Printf("Current latestPublished: %s\n", latestPublished.String())
}
