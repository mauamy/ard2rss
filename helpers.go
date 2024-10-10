package main

import (
	"ard_audiothek_rss/ard"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func readParam(r *http.Request, name string) string {
	params := httprouter.ParamsFromContext(r.Context())
	return params.ByName(name)
}

func getAllShowDataResults(showId string) ([]ard.Result, error) {
	hasNextPage := true
	count := 50
	offset := 0
	var results []ard.Result

	for hasNextPage {
		showData, err := ard.GetArdAudiothekShowDataWithGraphApi(showId, count, offset)
		if err != nil {
			return nil, err
		}

		showDataResult := showData.Data.Result
		results = append(results, showDataResult)

		hasNextPage = showDataResult.Items.PageInfo.HasNextPage
		if hasNextPage {
			offset += len(showDataResult.Items.Nodes)
		}
	}

	return results, nil
}
