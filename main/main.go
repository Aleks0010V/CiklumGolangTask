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
		cm       modules.ContentMarketingResponse
		res      modules.ResponseByList
		err      error
	)
	w.Header().Set("content-type", "application/json")
	if err = articles.FetchArticles(); err != nil {
		log.Fatalf("Articles was not received: %v", err)
	}
	if err = cm.FetchContentMarketingData(); err != nil {
		log.Fatalf("Articles was not received: %v", err)
	}

	res.MergeArticlesWithMarketing(articles.Response.Items, cm.Response.Items, 6)
	resJSON, err := json.Marshal(res.Items)
	if err == nil {
		_, err := w.Write(resJSON)
		if err != nil {
			log.Printf("%v %v %v : JSON was not sent: %v", r.Method, r.Host, r.URL.Path, err)
		} else {
			log.Printf("%v %v %v : OK", r.Method, r.Host, r.URL.Path)
		}
	}
}

func handleRequests() {
	listeningString := ":8888"
	http.HandleFunc("/", listApiHandler)
	log.Print("STARTING API")
	log.Fatal(http.ListenAndServe(listeningString, nil))
}

func main() {
	handleRequests()
}
