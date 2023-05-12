package main

import (
	"ard_audiothek_rss/ard"
	"fmt"
	"github.com/gorilla/feeds"
	"strings"
)

func CreateRSSFeed(result ard.Result, feedUrl string) *feeds.Feed {
	imgUrl := strings.ReplaceAll(result.Image.URL, "{width}", "300")

	feed := &feeds.Feed{
		Title:       result.Title,
		Link:        &feeds.Link{Href: fmt.Sprintf("https://ard2rss.mauamy.de/%s", feedUrl)},
		Description: result.Description,
		Id:          result.ID,
		Image:       &feeds.Image{Url: imgUrl},
	}
	var feedItems []*feeds.Item

	for _, item := range result.Items.Nodes {
		audioUrl := item.Audios[0].DownloadURL
		if audioUrl == "" {
			audioUrl = item.Audios[0].URL
		}

		feedItem := &feeds.Item{
			Id:          item.ID,
			Title:       item.Title,
			Link:        &feeds.Link{Href: fmt.Sprintf("https://ard2rss.mauamy.de/%s", feedUrl)},
			Created:     item.PublishDate,
			Description: item.Summary,
			Enclosure: &feeds.Enclosure{
				Url:    audioUrl,
				Length: fmt.Sprintf("%d", item.Duration),
				Type:   "audio/mpeg",
			},
		}
		feedItems = append(feedItems, feedItem)
	}

	feed.Items = feedItems
	return feed
}
