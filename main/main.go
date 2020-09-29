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
		err      error
	)
	w.Header().Set("content-type", "application/json")
	err = articles.FetchArticles()
	if err != nil {
		log.Fatalf("Articles was not received: %v", err)
	}
	err = ads.FetchContentMarketingData()
	if err != nil {
		log.Fatalf("Articles was not received: %v", err)
	}

	res.MergeArticlesWithMarketing(articles.Response.Items, ads.Response.Items, 6)
	resJSON, err := json.Marshal(res.Items)
	if err == nil {
		_, err := w.Write(resJSON)
		if err != nil {
			log.Fatalf("JSON was not sent: %v", err)
		}
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
