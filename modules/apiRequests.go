package modules

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (ar *ArticlesResponse)FetchArticles () error{
	var (
		err error
	)

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

	if isValid := json.Valid(body); isValid {
		err = json.Unmarshal(body, ar)
		return err
	} else {
		return errors.New("articles JSON is not valid")
	}
}

func (ad *ContentMarketingResponse)FetchContentMarketingData() error {
	var (
		err error
	)

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

	if isValid := json.Valid(body); isValid {
		err = json.Unmarshal(body, ad)
		return err
	} else {
		return errors.New("articles JSON is not valid")
	}
}