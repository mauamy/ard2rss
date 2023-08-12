package main

import (
	"ard_audiothek_rss/ard"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"time"
)

const ardUrl = "https://www.ardaudiothek.de/sendung/levels-und-soundtracks/12642475/"

// <script id="__NEXT_DATA__"

func GetRSSFeed(w http.ResponseWriter, r *http.Request) {
	mediaType := readParam(r, "type")
	name := readParam(r, "name")
	id := readParam(r, "id")

	if mediaType != "sendung" && mediaType != "sammlung" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	ardUrl := fmt.Sprintf("https://www.ardaudiothek.de/%s/%s/%s", mediaType, name, id)

	showData, err := ard.GetArdAudiothekShowData(ardUrl)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error: %s", err.Error())
		return
	}
	showDataResult := showData.Props.PageProps.InitialData.Data.Result

	rssFeed := ard.CreateRSSFeed(showDataResult, fmt.Sprintf("%s/%s/%s", mediaType, name, id))
	xml, err := rssFeed.ToRss()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error: %s", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(xml))
}

func routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/:type/:name/:id", GetRSSFeed)

	return rateLimit(router)
}

func main() {
	logger := log.New(os.Stdout, "", 0)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", 8080),
		Handler:      routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err := srv.ListenAndServe()
	logger.Fatal(err)
}
