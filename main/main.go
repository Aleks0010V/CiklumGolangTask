package main

import (
	"CiklumGolangTask/modules"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
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

func handleRequests(ctx context.Context) (err error) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", listApiHandler)
	srv := &http.Server{
		Addr:    listeningString,
		Handler: mux,
	}
	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%s\n", err)
		}
	}()
	log.Printf("STARTING API on port %s", listeningString[1:])
	<-ctx.Done()
	log.Print("STOPPING API")
	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err = srv.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed:%+s", err)
	}

	log.Printf("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}

	return
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-c
		log.Printf("system call:%+v", oscall)
		cancel()
	}()

	if err := handleRequests(ctx); err != nil {
		log.Printf("failed to serve:+%v\n", err)
	}
}
