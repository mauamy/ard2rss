package ard

type GraphQlRequestData struct {
	OperationName string                  `json:"operationName"`
	Query         string                  `json:"query"`
	Variables     GraphQlRequestVariables `json:"variables"`
}

type GraphQlRequestVariables struct {
	Id     string `json:"id"`
	Count  int    `json:"count"`
	Offset int    `json:"offset"`
}

func NewProgramSetEpisodesQuery(showId string, count, offset int) GraphQlRequestData {
	return GraphQlRequestData{
		OperationName: "ProgramSetEpisodesQuery",
		Query:         ProgramSetEpisodesQuery,
		Variables:     GraphQlRequestVariables{Id: showId, Count: count, Offset: offset},
	}
}

type GraphqlResponse struct {
	Data struct {
		Result Result `json:"result"`
	} `json:"data"`
}

var ProgramSetEpisodesQuery = `
query ProgramSetEpisodesQuery($id:ID!,$offset:Int!,$count:Int!){
  result:programSet(id:$id){
    numberOfElements
    id
    title
    description
    items(
      offset:$offset 
      first:$count 
      filter:{isPublished:{equalTo:true},itemType:{notEqualTo:EVENT_LIVESTREAM}}
    ){
      pageInfo{hasNextPage endCursor}
      nodes{
        id
        coreId
        title
        isPublished
        tracking 
        publishDate 
        summary 
        duration 
        path 
        image{
          url 
          url1X1 
          description 
          attribution
        }
        programSet{
          id 
          coreId 
          title 
          path 
          publicationService{
            title 
            genre 
            path 
            organizationName
          }
        }
        audios{url mimeType downloadUrl allowDownload}
      }
    }
  }
}`
