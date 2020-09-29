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
