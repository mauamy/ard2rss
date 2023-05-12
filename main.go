package main

import (
	"ard_audiothek_rss/ard"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const ardUrl = "https://www.ardaudiothek.de/sendung/levels-und-soundtracks/12642475/"

// <script id="__NEXT_DATA__"

func GetRSSFeed(c *gin.Context) {
	mediaType := c.Param("type")
	name := c.Param("name")
	id := c.Param("id")

	log.Println("Request URL: ")

	ardUrl := fmt.Sprintf("https://www.ardaudiothek.de/%s/%s/%s", mediaType, name, id)

	showData, err := ard.GetArdAudiothekShowData(ardUrl)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	showDataResult := showData.Props.PageProps.InitialData.Data.Result

	rssFeed := CreateRSSFeed(showDataResult, fmt.Sprintf("%s/%s/%s", mediaType, name, id))
	xml, err := rssFeed.ToRss()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.String(200, xml)
}

func main() {
	r := gin.Default()

	r.GET("/:type/:name/:id", GetRSSFeed)

	r.Run(":8080")

}
