package ard

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/antchfx/htmlquery"
	"io"
	"net/http"
)

func GetArdAudiothekShowData(showUrl string) (*ShowData, error) {
	doc, err := htmlquery.LoadURL(showUrl)
	if err != nil {
		return nil, err
	}

	jsonScript := htmlquery.FindOne(doc, "//script[@id='__NEXT_DATA__']")
	jsonStr := htmlquery.InnerText(jsonScript)

	var audioData ShowData
	err = json.Unmarshal([]byte(jsonStr), &audioData)

	return &audioData, err
}

func GetArdAudiothekShowDataWithGraphApi(showId string, count, offset int) (*GraphqlResponse, error) {
	endpoint := "https://api.ardaudiothek.de/graphql"

	requestData := NewProgramSetEpisodesQuery(showId, count, offset)
	jsonRequestData, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonRequestData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respData, _ := io.ReadAll(resp.Body)
	fmt.Printf("resp data:\n %v\n", string(respData))

	var graphQlResp GraphqlResponse
	err = json.Unmarshal(respData, &graphQlResp)
	if err != nil {
		return nil, err
	}

	return &graphQlResp, nil
}
