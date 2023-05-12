package ard

import "time"

type Result struct {
	NumberOfElements    int    `json:"numberOfElements"`
	ID                  string `json:"id"`
	Title               string `json:"title"`
	Description         string `json:"description"`
	EditorialCategories struct {
		Nodes []struct {
			Title string `json:"title"`
			ID    string `json:"id"`
		} `json:"nodes"`
	} `json:"editorialCategories"`
	Path  string `json:"path"`
	Image struct {
		URL         string `json:"url"`
		URL1X1      string `json:"url1X1"`
		Description string `json:"description"`
		Attribution string `json:"attribution"`
	} `json:"image"`
	PublicationService struct {
		Title            string `json:"title"`
		Genre            string `json:"genre"`
		Path             string `json:"path"`
		OrganizationName string `json:"organizationName"`
	} `json:"publicationService"`
	Items struct {
		PageInfo struct {
			HasNextPage bool   `json:"hasNextPage"`
			EndCursor   string `json:"endCursor"`
		} `json:"pageInfo"`
		Nodes []struct {
			ID          string    `json:"id"`
			Title       string    `json:"title"`
			PublishDate time.Time `json:"publishDate"`
			Summary     string    `json:"summary"`
			Duration    int       `json:"duration"`
			Path        string    `json:"path"`
			Image       struct {
				URL         string `json:"url"`
				URL1X1      string `json:"url1X1"`
				Description string `json:"description"`
				Attribution string `json:"attribution"`
			} `json:"image"`
			ProgramSet struct {
				ID                 string `json:"id"`
				Title              string `json:"title"`
				Path               string `json:"path"`
				PublicationService struct {
					Title            string `json:"title"`
					Genre            string `json:"genre"`
					Path             string `json:"path"`
					OrganizationName string `json:"organizationName"`
				} `json:"publicationService"`
			} `json:"programSet"`
			Audios []struct {
				URL           string `json:"url"`
				DownloadURL   string `json:"downloadUrl"`
				AllowDownload bool   `json:"allowDownload"`
			} `json:"audios"`
		} `json:"nodes"`
	} `json:"items"`
}

type ShowData struct {
	Props struct {
		PageProps struct {
			InitialData struct {
				Data struct {
					Result Result `json:"result"`
				} `json:"data"`
			} `json:"initialData"`
		} `json:"pageProps"`
		DeviceType  string `json:"deviceType"`
		RubricsData struct {
			Data struct {
				EditorialCategories struct {
					Title    string `json:"title"`
					Tracking struct {
						Title    string `json:"title"`
						Chapter1 string `json:"chapter1"`
						Page     struct {
							PageTitle    string `json:"page_title"`
							PageID       string `json:"page_id"`
							PageChapter1 string `json:"page_chapter1"`
							PageChapter2 string `json:"page_chapter2"`
						} `json:"page"`
					} `json:"tracking"`
					Nodes []struct {
						ID    string `json:"id"`
						Title string `json:"title"`
						Links struct {
							Self struct {
								Href      string `json:"href"`
								Name      string `json:"name"`
								Templated bool   `json:"templated"`
							} `json:"self"`
							MtImage struct {
								Href      string `json:"href"`
								Name      string `json:"name"`
								Templated bool   `json:"templated"`
							} `json:"mt:image"`
						} `json:"_links"`
						Image struct {
							URL    string `json:"url"`
							URL1X1 any    `json:"url1X1"`
						} `json:"image"`
					} `json:"nodes"`
					Links struct {
						Self struct {
							Href string `json:"href"`
							Name string `json:"name"`
						} `json:"self"`
					} `json:"_links"`
				} `json:"editorialCategories"`
			} `json:"data"`
		} `json:"rubricsData"`
		ChannelData struct {
			Data struct {
				Organizations struct {
					Nodes []struct {
						Name string `json:"name"`
					} `json:"nodes"`
				} `json:"organizations"`
			} `json:"data"`
		} `json:"channelData"`
		NSsp bool `json:"__N_SSP"`
	} `json:"props"`
	Page  string `json:"page"`
	Query struct {
		Title  string `json:"title"`
		ShowID string `json:"showId"`
	} `json:"query"`
	BuildID      string `json:"buildId"`
	IsFallback   bool   `json:"isFallback"`
	Gssp         bool   `json:"gssp"`
	AppGip       bool   `json:"appGip"`
	ScriptLoader []any  `json:"scriptLoader"`
}
