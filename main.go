package main

import (
	"ard_audiothek_rss/ard"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const ardUrl = "https://www.ardaudiothek.de/sendung/levels-und-soundtracks/12642475/"

// <script id="__NEXT_DATA__"

func GetRSSFeed(c *gin.Context) {
	showName := c.Param("showName")
	showId := c.Param("showId")
	ardUrl := fmt.Sprintf("https://www.ardaudiothek.de/sendung/%s/%s/", showName, showId)

	showData, err := ard.GetArdAudiothekShowData(ardUrl)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	showDataResult := showData.Props.PageProps.InitialData.Data.Result

	rssFeed := CreateRSSFeed(showDataResult)
	xml, err := rssFeed.ToRss()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.String(200, xml)
}

func main() {
	r := gin.Default()

	r.GET("/sendung/:showName/:showId", GetRSSFeed)

	r.Run(":8080")

}
