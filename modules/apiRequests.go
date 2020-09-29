package modules

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (ar *ArticlesResponse) FetchArticles() error {

	client := http.DefaultClient
	resp, err := client.Get("https://storage.googleapis.com/aller-structure-task/articles.json")
	if err != nil {
		fmt.Println("Articles error")
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Articles error")
		return err
	}
	defer resp.Body.Close()

	err = json.Unmarshal(body, ar)
	return err
}

func (ad *ContentMarketingResponse) FetchContentMarketingData() error {
	client := http.DefaultClient

	resp, err := client.Get("https://storage.googleapis.com/aller-structure-task/contentmarketing.json")
	if err != nil {
		fmt.Println("Articles error")
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Articles error")
		return err
	}
	defer resp.Body.Close()

	err = json.Unmarshal(body, ad)
	return err
}

func (result *ResponseByList) MergeArticlesWithMarketing(articles []Article, contentMarketing []ContentMarketing) {
	// by spec we need map each 5 articles to 1 ad, so the number of articles must be no less then 5*len(contentMarketing)
	if len(articles) < len(contentMarketing)*5 {
		return
	}

	result.Items = make([]interface{}, 0)

	cmCounter := 0
	for i := 0; i < len(articles); i++ {
		if i%5 == 0 && i != 0 && cmCounter < len(contentMarketing) {
			result.Items = append(result.Items, contentMarketing[cmCounter])
			cmCounter++
		} else if i%5 == 0 && cmCounter >= len(contentMarketing) {
			result.Items = append(result.Items, EAd)
		} else {
			result.Items = append(result.Items, articles[i])
		}
	}
}
