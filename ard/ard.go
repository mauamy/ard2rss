package ard

import (
	"encoding/json"
	"github.com/antchfx/htmlquery"
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
