package modules

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (ar *ArticlesResponse)FetchArticles () error{

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

func (ad *ContentMarketingResponse)FetchContentMarketingData() error {
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

func MergeArticlesWithMarketing (articles []Article, ads []Ad) Response{
	// by spec we need map each 5 articles to 1 ad, so the number of articles must be no less then 5*len(ads)
	if len(articles) < len(ads) * 5{
		result := Response{}
		return result
	}

	result := Response{
		Items: make([]ResponseItem, len(ads)),
	}
	for adIndex, ad := range ads {
		result.Items[adIndex].ContentMarketing = ad
		result.Items[adIndex].Articles = make([]Article, 5)
		result.Items[adIndex].Articles = articles[5*adIndex : 5*adIndex+5]
	}

	return result
}