package main

import (
	"CiklumGolangTask/modules"
	"encoding/json"
	"log"
	"net/http"
)

func listApiHandler(w http.ResponseWriter, r *http.Request) {
	var (
		articles modules.ArticlesResponse
		ads      modules.ContentMarketingResponse
		res      modules.ResponseByList
	)
	w.Header().Set("content-type", "application/json")
	artErr := articles.FetchArticles()
	adsErr := ads.FetchContentMarketingData()

	if artErr == nil && adsErr == nil {
		res.MergeArticlesWithMarketing(articles.Response.Items, ads.Response.Items)
	}
	resJSON, resErr := json.Marshal(res.Items)
	if resErr == nil {
		w.Write(resJSON)
	}
}

func handleRequests() {
	listeningString := ":8888"
	http.HandleFunc("/", listApiHandler)
	log.Fatal(http.ListenAndServe(listeningString, nil))
}

func main() {
	handleRequests()
}
