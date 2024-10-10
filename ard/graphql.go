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

// https://api.ardaudiothek.de/graphql?query=query%20ProgramSetEpisodesQuery(%24id%3AID!%2C%24offset%3AInt!%2C%24count%3AInt!)%7Bresult%3AprogramSet(id%3A%24id)%7Bitems(offset%3A%24offset%20first%3A%24count%20filter%3A%7BisPublished%3A%7BequalTo%3Atrue%7D%2CitemType%3A%7BnotEqualTo%3AEVENT_LIVESTREAM%7D%7D)%7BpageInfo%7BhasNextPage%20endCursor%7Dnodes%7Bid%20coreId%20title%20isPublished%20tracking%20publishDate%20summary%20duration%20path%20image%7Burl%20url1X1%20description%20attribution%7DprogramSet%7Bid%20coreId%20title%20path%20publicationService%7Btitle%20genre%20path%20organizationName%7D%7Daudios%7Burl%20mimeType%20downloadUrl%20allowDownload%7D%7D%7D%7D%7D&variables=%7B%22id%22%3A%2212449787%22%2C%22offset%22%3A24%2C%22count%22%3A12%7D%27
var ProgramSetEpisodesQuery = `
query ProgramSetEpisodesQuery($id:ID!,$offset:Int!,$count:Int!){
  result:programSet(id:$id){
    numberOfElements
    id
    title
    description
	image {
      url
    }
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
