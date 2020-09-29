package main

import (
	"CiklumGolangTask/modules"
	"encoding/json"
	"log"
	"net/http"
)

func apiHandler(w http.ResponseWriter, r *http.Request) {
	var (
		articles modules.ArticlesResponse
		ads modules.ContentMarketingResponse
		res modules.ResponseByObjects
		i modules.MarketingResponse
	)
	i = &res
	w.Header().Set("content-type", "application/json")
	artErr := articles.FetchArticles()
	adsErr := ads.FetchContentMarketingData()

	if artErr == nil && adsErr == nil {
		i.MergeArticlesWithMarketing(articles.Response.Items, ads.Response.Items)
	}
	resJSON, resErr := json.Marshal(res.Items)
	if resErr == nil {
		w.Write(resJSON)
	}
}

func handleRequests() {
	http.HandleFunc("/my-api", apiHandler)
	log.Fatal(http.ListenAndServe("127.0.1.1:8888", nil))
}

func main() {
	handleRequests()
}
