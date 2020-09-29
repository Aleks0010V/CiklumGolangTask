package modules

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func (ar *ArticlesResponse) FetchArticles() error {

	resp, err := http.Get("https://storage.googleapis.com/aller-structure-task/articles.json")
	if err != nil {
		log.Printf("Articles fetching error: %v", err)
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Printf("Articles fetching error: %v", err)
		return err
	}

	err = json.Unmarshal(body, ar)
	return err
}

func (cm *ContentMarketingResponse) FetchContentMarketingData() error {

	resp, err := http.Get("https://storage.googleapis.com/aller-structure-task/contentmarketing.json")
	if err != nil {
		log.Printf("Marketing fetching error: %v", err)
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Printf("Marketing fetching error: %v", err)
		return err
	}

	err = json.Unmarshal(body, cm)
	return err
}
