package ard

import (
	"fmt"
	"github.com/gorilla/feeds"
	"strings"
)

func CreateRSSFeed(results []Result, feedUrl string) *feeds.Feed {
	baseResult := results[0]
	imgUrl := strings.ReplaceAll(baseResult.Image.URL, "{width}", "300")

	feed := &feeds.Feed{
		Title:       baseResult.Title,
		Link:        &feeds.Link{Href: fmt.Sprintf("https://ard2rss.mauamy.de/%s", feedUrl)},
		Description: baseResult.Description,
		Id:          baseResult.ID,
		Image:       &feeds.Image{Url: imgUrl},
	}
	var feedItems []*feeds.Item

	for _, result := range results {
		for _, item := range result.Items.Nodes {
			audioUrl := item.Audios[0].DownloadURL
			if audioUrl == "" {
				audioUrl = item.Audios[0].URL
			}

			feedItem := &feeds.Item{
				Id:      item.ID,
				Title:   item.Title,
				Link:    &feeds.Link{Href: fmt.Sprintf("https://ard2rss.mauamy.de/%s", feedUrl)},
				Created: item.PublishDate,

				Description: item.Summary,
				Enclosure: &feeds.Enclosure{
					Url:    audioUrl,
					Length: fmt.Sprintf("%d", item.Duration),
					Type:   "audio/mpeg",
				},
			}
			feedItems = append(feedItems, feedItem)
		}
	}

	feed.Items = feedItems
	return feed
}
