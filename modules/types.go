package modules


// data types for response from articles API
type ArticlesResponse struct {
	HttpStatus int
	Response ArticlesList
}

type ArticlesList struct {
	Items []Article
}

type Article struct {
	Type string
	HarvesterId string
	Cerebro_score float32  // ToDo - rename this field, Unmarshal doesn't recognize it
	URL string
	Title string
	CleanImage string
}


// data types for response from content marketing API
type ContentMarketingResponse struct {
	HttpStatus int
	Response AdsList
}

type AdsList struct {
	Items []Ad
}

type Ad struct {
	Type string
	HarvesterId string
	CommercialPartner string
	LogoURL string
	Cerebro_score float32  // ToDo - rename this field, Unmarshal doesn't recognize it
	URL string
	Title string
	CleanImage string
}

type Response struct {
	Items []ResponseItem
}

type ResponseItem struct {
	Articles []Article
	ContentMarketing interface{}  // can be either Ad type or EmptyAd type
}

type EmptyAd map[string]string
var EAd = EmptyAd {"Type": "Ad"}