package main

import (
	"CiklumGolangTask/modules"
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

var cmPosition = 5            // position of Content Marketing in pattern
var listeningString = ":8888" // addr on machine to start API

func listApiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(405)
		log.Printf("%v %v %v : 405 Method Not Allowed", r.Method, r.Host, r.URL.Path)
		return
	}

	var (
		articles modules.ArticlesResponse
		cm       modules.ContentMarketingResponse
		res      modules.ResponseByList
		err      error
		wg       sync.WaitGroup
	)
	w.Header().Set("content-type", "application/json")

	wg.Add(2)
	go func() {
		if err = articles.FetchArticles(); err != nil {
			log.Fatalf("Articles was not received: %v", err)
		}
		wg.Done()
	}()

	go func() {
		if err = cm.FetchContentMarketingData(); err != nil {
			log.Fatalf("ContentMarketing was not received: %v", err)
		}
		wg.Done()
	}()
	wg.Wait()

	res.MergeArticlesWithMarketing(articles.Response.Items, cm.Response.Items, cmPosition)
	if resJSON, err := json.Marshal(res.Items); err == nil {
		_, err := w.Write(resJSON)
		if err != nil {
			log.Printf("%v %v %v : JSON was not sent: %v", r.Method, r.Host, r.URL.Path, err)
		} else {
			log.Printf("%v %v %v : 200 OK", r.Method, r.Host, r.URL.Path)
		}
	}
}

func handleRequests() {
	// ToDO - ideally, it should have graceful shutdown method
	http.HandleFunc("/", listApiHandler)
	log.Printf("STARTING API on port %s", listeningString[1:])
	log.Fatal(http.ListenAndServe(listeningString, nil))
}

func main() {
	handleRequests()
}
