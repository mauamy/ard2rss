package main

import (
	"ard_audiothek_rss/ard"
	"github.com/gorilla/feeds"
)

func CreateRSSFeed(result ard.Result) *feeds.Feed {
	feed := &feeds.Feed{
		Title:       result.Title,
		Link:        &feeds.Link{Href: "your feed link"},
		Description: result.Description,
		Id:          result.ID,
		Image:       &feeds.Image{Url: result.Image.URL},
	}
	var feedItems []*feeds.Item

	for _, item := range result.Items.Nodes {
		feedItem := &feeds.Item{
			Id:      item.ID,
			Title:   item.Title,
			Link:    &feeds.Link{Href: item.Audios[0].DownloadURL},
			Created: item.PublishDate,
		}
		feedItems = append(feedItems, feedItem)
	}

	feed.Items = feedItems
	return feed
}
