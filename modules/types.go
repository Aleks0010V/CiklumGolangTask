package modules

// data types for response from articles API
type ArticlesResponse struct {
	HttpStatus int          `json:"http_status"`
	Response   ArticlesList `json:"response"`
}

type ArticlesList struct {
	Items []Article `json:"items"`
}

type Article struct {
	Type          string  `json:"type"`
	HarvesterId   string  `json:"harvesterId"`
	Cerebro_score float32 `json:"cerebro-score"`
	URL           string  `json:"url"`
	Title         string  `json:"title"`
	CleanImage    string  `json:"cleanImage"`
}

// data types for response from content marketing API
type ContentMarketingResponse struct {
	HttpStatus int     `json:"http_status"`
	Response   AdsList `json:"response"`
}

type AdsList struct {
	Items []ContentMarketing `json:"items"`
}

type ContentMarketing struct {
	Type              string  `json:"type"`
	HarvesterId       string  `json:"harvesterId"`
	CommercialPartner string  `json:"commercialPartner"`
	LogoURL           string  `json:"logoURL"`
	Cerebro_score     float32 `json:"cerebro-score"`
	URL               string  `json:"url"`
	Title             string  `json:"title"`
	CleanImage        string  `json:"cleanImage"`
}

type ResponseByList struct {
	Items []interface{} `json:"items"` // just a slice of articles and ads by pattern "5art-1ad"
}

type EmptyAd struct {
	Type string `json:"type"`
}

var EAd = EmptyAd{Type: "Ad"}
