package modules

import "testing"

func TestFetchArticles(t *testing.T) {
	var (
		err      error
		articles ArticlesResponse
	)

	if err = articles.FetchArticles(); err != nil {
		t.Errorf("Error in articles fetching func: %v", err)
	} else {
		t.Log("FetchArticles works as expected")
	}
}

func TestFetchContentMarketingData(t *testing.T) {
	var (
		err error
		cm  ContentMarketingResponse
	)

	if err = cm.FetchContentMarketingData(); err != nil {
		t.Errorf("Error in content marketing data fetching func: %v", err)
	} else {
		t.Log("FetchContentMarketingData works as expected")
	}
}
