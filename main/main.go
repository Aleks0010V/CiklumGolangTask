package main

import (
	"CiklumGolangTask/modules"
	"fmt"
	"log"
	"net/http"
)

func apiHandler(w http.ResponseWriter, r *http.Request) {
	var (
		articles modules.ArticlesResponse
		ads modules.ContentMarketingResponse
		res modules.Response
	)
	artErr := articles.FetchArticles()
	adsErr := ads.FetchContentMarketingData()

	if artErr == nil && adsErr == nil {
		res = modules.MergeArticlesWithMarketing(articles.Response.Items, ads.Response.Items)
	}

	fmt.Println(res)
}

func handleRequests() {
	http.HandleFunc("/my-api", apiHandler)
	log.Fatal(http.ListenAndServe("127.0.1.1:8888", nil))
}

func main() {
	handleRequests()
}
