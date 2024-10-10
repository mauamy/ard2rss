package ard

import (
	"testing"
)

//func TestGetArdAudiothekShowData(t *testing.T) {
//	testShowUrl := "https://www.ardaudiothek.de/sendung/dark-matters-geheimnisse-der-geheimdienste/12449787/"
//
//	showData, err := GetArdAudiothekShowData(testShowUrl)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//
//	results := showData.Props.PageProps.InitialData.Data.Result
//
//	json, err := json.MarshalIndent(results, "", "  ")
//	if err != nil {
//		t.Error(err)
//	}
//
//	os.WriteFile("results.json", json, 0644)
//
//	t.Logf("Number of Items: %d", results.NumberOfElements)
//	t.Logf("Has Next Page: %t | endCursor: %s", results.Items.PageInfo.HasNextPage, results.Items.PageInfo.EndCursor)
//	t.Logf("Nodes: %d", len(results.Items.Nodes))
//}

func TestGetArdAudiothekShowDataGraphApi(t *testing.T) {
	showID := "12449787"
	count := 5

	resp, err := GetArdAudiothekShowDataWithGraphApi(showID, 5, 0)
	if err != nil {
		t.Error(err)
	}
	result := resp.Data.Result
	if len(result.Items.Nodes) != count {
		t.Errorf("Result node length should be %d, got %d", count, len(result.Items.Nodes))
	}
}
