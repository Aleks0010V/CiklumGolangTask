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

func (result *ResponseByObjects)MergeArticlesWithMarketing (articles []Article, ads []Ad) {
	// by spec we need map each 5 articles to 1 ad, so the number of articles must be no less then 5*len(ads)
	if len(articles) < len(ads) * 5{
		return
	}

	// map all ads
	result.Items = make([]ResponseItem, len(ads))
	for adIndex, ad := range ads {
		result.Items[adIndex].ContentMarketing = ad
		result.Items[adIndex].Articles = make([]Article, 5)
		result.Items[adIndex].Articles = articles[5*adIndex : 5*adIndex+5]
	}
	articles = articles[len(ads) * 5 : ]

	// map all other articles with empty ad
	for ;; {
		if len(articles) >= 5 {
			nextResponseItem := ResponseItem{
				Articles:         articles[:5],
				ContentMarketing: EAd,  // empty Ad
			}
			result.Items = append(result.Items, nextResponseItem)
			articles = articles[5:]
		} else {
			nextResponseItem := ResponseItem{
				Articles:         articles,
				ContentMarketing: EAd,  // empty Ad
			}
			result.Items = append(result.Items, nextResponseItem)
			break
		}
	}
}

func (result *ResponseByList)MergeArticlesWithMarketing (articles []Article, ads []Ad) {
	// by spec we need map each 5 articles to 1 ad, so the number of articles must be no less then 5*len(ads)
	if len(articles) < len(ads) * 5{
		return
	}

	var iItems uint16 = 0
	// I want to allocate exactly needed amount of memory
	// for each 5 articles I need to add 1 add. If len(articles) is not a multiple of 5,
	// I will add 1 memory cell to be sure everything will fit
	itemsLength := len(articles) + len(articles) / 5 + 1
	result.Items = make([]interface{}, itemsLength)

	for ;; {
		var i = 0
		for ; i < len(articles) && i < 5; i++{
			result.Items[iItems] = articles[i]
			iItems++
		}
		articles = articles[i:]

		if len(ads) > 0 {
			result.Items[iItems] = ads[0]
			iItems++
			ads = ads[1:]
		} else {
			result.Items[iItems] = EAd
			iItems++
			ads = []Ad{}
		}

		if len(ads) == 0 && len(articles) == 0 {break}
	}
}