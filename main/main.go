package main

import (
	"CicklumGolangTask/modules"
	"fmt"
	"log"
	"net/http"
)

func apiHandler(w http.ResponseWriter, r *http.Request) {
	var (
		articles modules.ArticlesResponse
		ads modules.ContentMarketingResponse
	)
	artErr := articles.FetchArticles()
	adsErr := ads.FetchContentMarketingData()
	fmt.Println(articles, artErr)
	fmt.Println(ads, adsErr)
}

func handleRequests() {
	http.HandleFunc("/my-api", apiHandler)
	log.Fatal(http.ListenAndServe("127.0.1.1:8888", nil))
}

func main() {
	handleRequests()
}
