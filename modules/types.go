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
	HarvesterId   string  `json:"harvester_id"`
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
	HarvesterId       string  `json:"harvester_id"`
	CommercialPartner string  `json:"commercial_partner"`
	LogoURL           string  `json:"logo_url"`
	Cerebro_score     float32 `json:"cerebro-score"`
	URL               string  `json:"url"`
	Title             string  `json:"title"`
	CleanImage        string  `json:"cleanImage"`
}

type ResponseByObjects struct {
	Items []ResponseItem `json:"items"`
}

type ResponseByList struct {
	Items []interface{} `json:"items"` // just a slice of articles and ads by pattern "5art-1ad"
}

type ResponseItem struct {
	Articles         []Article   `json:"articles"`
	ContentMarketing interface{} `json:"content_marketing"` // can be either ContentMarketing type or EmptyAd type
}

type EmptyAd struct {
	Type string `json:"type"`
}

var EAd = EmptyAd{Type: "Ad"}
